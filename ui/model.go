package ui

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kjkondratuk/goblins-and-gold/command"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"golang.org/x/term"
)

// InitMsg is sent to initialize the model with state and app
type InitMsg struct {
	State state.State
	App   command.Command
}

// ShowSelectMsg is sent to show a select control
type ShowSelectMsg struct {
	Label    string
	Items    []string
	ResultCh chan<- SelectResult
}

// ShowPromptMsg is sent to show a prompt
type ShowPromptMsg struct {
	Label        string
	DefaultValue string
	ResultCh     chan<- PromptResult
}

// WindowSizeTickMsg is sent by the window size ticker
type WindowSizeTickMsg time.Time

// Model is the main bubbletea model for the application
type Model struct {
	state       state.State
	app         command.Command
	roomView    RoomView
	commandView CommandView
	messageView MessageView
	selectView  *selectModel
	promptView  *promptModel
	width       int
	height      int
	exit        chan os.Signal
	tooSmall    bool // Flag to track if the window is too small
}

// NewModel creates a new bubbletea model
func NewModel(s state.State, app command.Command, exit chan os.Signal) Model {
	return Model{
		state:       s,
		app:         app,
		roomView:    NewRoomView(s),
		commandView: NewCommandView(s, app),
		messageView: NewMessageView(),
		exit:        exit,
		tooSmall:    false, // Initialize as false, will be updated on first WindowSizeMsg
	}
}

// Init initializes the model
func (m Model) Init() tea.Cmd {
	// Check if we're running on Windows
	if runtime.GOOS == "windows" {
		return m.startWindowSizeTicker()
	}
	return nil
}

// startWindowSizeTicker starts a ticker to check window size on Windows
func (m Model) startWindowSizeTicker() tea.Cmd {
	return tea.Tick(time.Millisecond*250, func(t time.Time) tea.Msg {
		return WindowSizeTickMsg(t)
	})
}

// Minimum window dimensions
const (
	minWidth  = 80
	minHeight = 24
)

// Update handles messages and user input
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case InitMsg:
		// Initialize the model with state and app
		m.state = msg.State
		m.app = msg.App

		// Re-initialize the views with the new state and app
		m.roomView = NewRoomView(m.state)
		m.commandView = NewCommandView(m.state, m.app)

		// Keep the exit channel (it was set during model creation)

		// Add an initial message
		m.messageView.AddMessage("Welcome to Goblins and Gold!", LogInfo)

		// Look at the current room
		return m, func() tea.Msg {
			err := m.app.Run(m.state, "look")
			return CommandExecutedMsg{
				Command: "look",
				Error:   err,
			}
		}
	case ShowSelectMsg:
		// Show a select control
		m.selectView = &selectModel{
			label:    msg.Label,
			items:    msg.Items,
			cursor:   0,
			resultCh: msg.ResultCh,
		}
		return m, nil
	case ShowPromptMsg:
		// Show a prompt
		ti := textinput.New()
		ti.Placeholder = msg.DefaultValue
		ti.Focus()

		m.promptView = &promptModel{
			label:        msg.Label,
			defaultValue: msg.DefaultValue,
			textInput:    ti,
			resultCh:     msg.ResultCh,
		}
		return m, textinput.Blink
	case tea.KeyMsg:
		// If we have an active select view, handle its key events
		if m.selectView != nil {
			// First check for page up/down keys for message view scrolling
			switch msg.String() {
			case "pgup", "pgdown", "home", "end":
				// Update message view for scrolling
				var messageViewCmd tea.Cmd
				m.messageView, messageViewCmd = m.messageView.Update(msg)
				return m, messageViewCmd
			}

			// Then handle select-specific keys
			switch msg.String() {
			case "up", "k":
				if m.selectView.cursor > 0 {
					m.selectView.cursor--
				}
				return m, nil
			case "down", "j":
				if m.selectView.cursor < len(m.selectView.items)-1 {
					m.selectView.cursor++
				}
				return m, nil
			case "enter":
				// Send the result and clear the select view
				m.selectView.resultCh <- SelectResult{
					Index: m.selectView.cursor,
					Value: m.selectView.items[m.selectView.cursor],
				}
				m.selectView = nil
				return m, nil
			case "ctrl+c", "esc":
				// Cancel and clear the select view
				m.selectView.resultCh <- SelectResult{Error: nil}
				m.selectView = nil
				return m, nil
			}
		}
		// If we have an active prompt view, handle its key events
		if m.promptView != nil {
			// First check for page up/down keys for message view scrolling
			switch msg.String() {
			case "pgup", "pgdown", "home", "end":
				// Update message view for scrolling
				var messageViewCmd tea.Cmd
				m.messageView, messageViewCmd = m.messageView.Update(msg)
				return m, messageViewCmd
			}

			switch msg.Type {
			case tea.KeyEnter:
				// Send the result and clear the prompt view
				m.promptView.resultCh <- PromptResult{Value: m.promptView.textInput.Value()}
				m.promptView = nil
				return m, nil
			case tea.KeyCtrlC, tea.KeyEsc:
				// Cancel and clear the prompt view
				m.promptView.resultCh <- PromptResult{Error: nil}
				m.promptView = nil
				return m, nil
			default:
				// Update the text input
				var cmd tea.Cmd
				m.promptView.textInput, cmd = m.promptView.textInput.Update(msg)
				return m, cmd
			}
		}
		// Only handle special keys here, command execution happens in CommandView
		// when Enter is pressed
	case WindowSizeTickMsg:
		// Continue the ticker
		cmds = append(cmds, m.startWindowSizeTicker())

		// Get the current terminal size
		width, height, err := term.GetSize(int(os.Stdout.Fd()))
		if err == nil && (width != m.width || height != m.height) {
			// If the size has changed, send a WindowSizeMsg
			cmds = append(cmds, func() tea.Msg {
				return tea.WindowSizeMsg{Width: width, Height: height}
			})
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		// Check if the window is too small
		wasTooSmall := m.tooSmall
		m.tooSmall = m.width < minWidth || m.height < minHeight

		// If the window was too small but is now big enough, add a message
		if wasTooSmall && !m.tooSmall {
			m.messageView.AddMessage("Window size is now acceptable.", LogSuccess)
		}
	}

	// Update components
	// Only update other components if there's no active select or prompt view
	if m.selectView == nil && m.promptView == nil {
		var roomViewCmd tea.Cmd
		m.roomView, roomViewCmd = m.roomView.Update(msg)
		cmds = append(cmds, roomViewCmd)

		var commandViewCmd tea.Cmd
		m.commandView, commandViewCmd = m.commandView.Update(msg)
		cmds = append(cmds, commandViewCmd)

		var messageViewCmd tea.Cmd
		m.messageView, messageViewCmd = m.messageView.Update(msg)
		cmds = append(cmds, messageViewCmd)
	}

	return m, tea.Batch(cmds...)
}

// View renders the UI
func (m Model) View() string {
	// Check if the window is too small
	if m.tooSmall {

		// Create an error message for small window
		errorStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("196")). // Red color
			Bold(true).
			Padding(1).
			Align(lipgloss.Center)

		message := "Window is too small!\n\n" +
			"Please resize your terminal to at least " +
			fmt.Sprintf("%d×%d", minWidth, minHeight) +
			" characters."

		return errorStyle.Render(message)
	}

	// Calculate layout with adjustments for borders and padding
	// Border and padding take up approximately 4 lines (2 for borders, 2 for padding)
	borderAndPaddingHeight := 4

	// Adjust total available height to account for borders and padding of all three components
	availableHeight := m.height - (borderAndPaddingHeight * 3)

	// Ensure availableHeight is not negative
	if availableHeight < 0 {
		availableHeight = 0
	}

	// Allocate proportions of the available height with minimum values
	roomHeight := availableHeight * 4 / 10 // 40% of available height
	if roomHeight < 3 {
		roomHeight = 3 // minimum 3
	}

	messageHeight := availableHeight * 3 / 10 // 30% of available height
	if messageHeight < 3 {
		messageHeight = 3 // minimum 3
	}

	commandHeight := availableHeight * 3 / 10 // 30% of available height
	if commandHeight < 3 {
		commandHeight = 3 // minimum 3
	}

	// Ensure width is not too small
	width := m.width
	if width < 10 {
		width = 10 // minimum width of 10
	}

	// Render room and message views
	roomView := lipgloss.NewStyle().
		Width(width).
		Height(roomHeight).
		Render(m.roomView.View())

	messageView := lipgloss.NewStyle().
		Width(width).
		Height(messageHeight).
		Render(m.messageView.View())

	// Render command area (either select view, prompt view, or command view)
	var commandAreaContent string
	if m.selectView != nil {
		// Render select view in place of command view
		// Adjust width to account for border (2 characters)
		borderWidth := 2
		adjustedWidth := width - borderWidth
		commandAreaContent = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(1).
			Width(adjustedWidth).
			Height(commandHeight).
			Render(m.selectView.View())
	} else if m.promptView != nil {
		// Render prompt view in place of command view
		// Adjust width to account for border (2 characters)
		borderWidth := 2
		adjustedWidth := width - borderWidth
		commandAreaContent = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(1).
			Width(adjustedWidth).
			Height(commandHeight).
			Render(m.promptView.label + "\n" + m.promptView.textInput.View())
	} else {
		// Render normal command view
		commandAreaContent = lipgloss.NewStyle().
			Width(width).
			Height(commandHeight).
			Render(m.commandView.View())
	}

	// Combine components
	return lipgloss.JoinVertical(lipgloss.Left, roomView, messageView, commandAreaContent)
}

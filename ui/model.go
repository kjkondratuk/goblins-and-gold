package ui

import (
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kjkondratuk/goblins-and-gold/command"
	"github.com/kjkondratuk/goblins-and-gold/state"
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
	}
}

// Init initializes the model
func (m Model) Init() tea.Cmd {
	return nil
}

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
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
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
	// Calculate layout
	roomHeight := m.height / 2
	messageHeight := m.height / 4
	commandHeight := m.height - roomHeight - messageHeight

	// Render room and message views
	roomView := lipgloss.NewStyle().
		Width(m.width).
		Height(roomHeight).
		Render(m.roomView.View())

	messageView := lipgloss.NewStyle().
		Width(m.width).
		Height(messageHeight).
		Render(m.messageView.View())

	// Render command area (either select view, prompt view, or command view)
	var commandAreaContent string
	if m.selectView != nil {
		// Render select view in place of command view
		commandAreaContent = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(1).
			Width(m.width).
			Height(commandHeight).
			Render(m.selectView.View())
	} else if m.promptView != nil {
		// Render prompt view in place of command view
		commandAreaContent = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(1).
			Width(m.width).
			Height(commandHeight).
			Render(m.promptView.label + "\n" + m.promptView.textInput.View())
	} else {
		// Render normal command view
		commandAreaContent = lipgloss.NewStyle().
			Width(m.width).
			Height(commandHeight).
			Render(m.commandView.View())
	}

	// Combine components
	return lipgloss.JoinVertical(lipgloss.Left, roomView, messageView, commandAreaContent)
}

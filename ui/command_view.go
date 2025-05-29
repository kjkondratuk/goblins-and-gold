package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kjkondratuk/goblins-and-gold/command"
	"github.com/kjkondratuk/goblins-and-gold/state"
)

// CommandExecutedMsg is sent when a command is executed
type CommandExecutedMsg struct {
	Command string
	Args    []string
	Error   error
}

// CommandView is a component that provides a command prompt for user input
type CommandView struct {
	state     state.State
	app       command.Command
	textInput textinput.Model
	style     lipgloss.Style
}

// NewCommandView creates a new command view
func NewCommandView(s state.State, app command.Command) CommandView {
	ti := textinput.New()
	ti.Placeholder = "Enter a command..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 50
	ti.Prompt = " >> "
	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return CommandView{
		state:     s,
		app:       app,
		textInput: ti,
		style: lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(1),
	}
}

// Update handles messages and user input
func (c CommandView) Update(msg tea.Msg) (CommandView, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			// Get the command text
			commandText := c.textInput.Value()
			if commandText == "" {
				return c, nil
			}

			// Reset the input
			c.textInput.Reset()

			// Tokenize the command
			args := strings.Split(commandText, " ")

			// Execute the command in a goroutine to avoid blocking the UI
			return c, func() tea.Msg {
				// Check if app or state is nil
				if c.app == nil || c.state == nil {
					return CommandExecutedMsg{
						Command: args[0],
						Args:    args[1:],
						Error:   nil,
					}
				}

				err := c.app.Run(c.state, args...)
				return CommandExecutedMsg{
					Command: args[0],
					Args:    args[1:],
					Error:   err,
				}
			}
		}
	}

	// Update the text input
	c.textInput, cmd = c.textInput.Update(msg)
	return c, cmd
}

// View renders the command prompt
func (c CommandView) View() string {
	return c.style.Render(c.textInput.View())
}

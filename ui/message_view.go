package ui

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// LogMsg is sent when a log message is generated
type LogMsg struct {
	Message string
	Level   LogLevel
}

// LogLevel represents the severity of a log message
type LogLevel int

const (
	LogInfo LogLevel = iota
	LogSuccess
	LogWarning
	LogError
	LogDebug
)

// MessageView is a component that displays system messages and logs
type MessageView struct {
	messages []formattedMessage
	viewport viewport.Model
	style    lipgloss.Style
}

type formattedMessage struct {
	text  string
	style lipgloss.Style
}

// NewMessageView creates a new message view
func NewMessageView() MessageView {
	vp := viewport.New(80, 10)
	vp.SetContent("")

	return MessageView{
		messages: []formattedMessage{},
		viewport: vp,
		style: lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(1),
	}
}

// Update handles messages and user input
func (m MessageView) Update(msg tea.Msg) (MessageView, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case LogMsg:
		// Add the message to the list
		m.messages = append(m.messages, formatLogMessage(msg))

		// Update the viewport content
		m.viewport.SetContent(m.formatMessages())

		// Scroll to the bottom
		m.viewport.GotoBottom()
	case CommandExecutedMsg:
		// If there was an error, log it
		if msg.Error != nil {
			m.messages = append(m.messages, formatLogMessage(LogMsg{
				Message: msg.Error.Error(),
				Level:   LogError,
			}))
			m.viewport.SetContent(m.formatMessages())
			m.viewport.GotoBottom()
		}
	}

	// Update the viewport
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

// View renders the message view
func (m MessageView) View() string {
	return m.style.Render(m.viewport.View())
}

// formatMessages formats all messages for display
func (m MessageView) formatMessages() string {
	var content string
	for _, msg := range m.messages {
		content += msg.style.Render(msg.text) + "\n"
	}
	return content
}

// formatLogMessage formats a log message based on its level
func formatLogMessage(msg LogMsg) formattedMessage {
	var style lipgloss.Style

	switch msg.Level {
	case LogInfo:
		style = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	case LogSuccess:
		style = lipgloss.NewStyle().Foreground(lipgloss.Color("42"))
	case LogWarning:
		style = lipgloss.NewStyle().Foreground(lipgloss.Color("208"))
	case LogError:
		style = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
	case LogDebug:
		style = lipgloss.NewStyle().Foreground(lipgloss.Color("99"))
	}

	return formattedMessage{
		text:  msg.Message,
		style: style,
	}
}

// AddMessage adds a message to the view
func (m *MessageView) AddMessage(message string, level LogLevel) {
	m.messages = append(m.messages, formatLogMessage(LogMsg{
		Message: message,
		Level:   level,
	}))
	m.viewport.SetContent(m.formatMessages())
	m.viewport.GotoBottom()
}

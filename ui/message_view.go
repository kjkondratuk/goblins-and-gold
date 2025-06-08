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
	// Create a viewport with initial dimensions
	// The actual size will be adjusted based on the terminal size
	vp := viewport.New(80, 10)
	vp.SetContent("")

	// Enable mouse wheel for scrolling
	vp.YPosition = 0
	vp.MouseWheelEnabled = true

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
	case tea.KeyMsg:
		// Handle scrolling with page up/down keys
		switch msg.String() {
		case "pgup":
			m.viewport.LineUp(5)
		case "pgdown":
			m.viewport.LineDown(5)
		case "home":
			m.viewport.GotoTop()
		case "end":
			m.viewport.GotoBottom()
		}
	}

	// Update the viewport
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

// View renders the message view
func (m MessageView) View() string {
	// Add help text for scrolling
	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")). // Dark gray color
		Italic(true).
		Align(lipgloss.Right)

	helpText := helpStyle.Render("PgUp/PgDn to scroll")

	// Combine viewport content with help text
	content := m.viewport.View() + "\n" + helpText

	return m.style.Render(content)
}

// formatMessages formats all messages for display
func (m MessageView) formatMessages() string {
	var content string

	// Get the width of the viewport (accounting for padding)
	viewportWidth := m.viewport.Width - 2

	for _, msg := range m.messages {
		// Check if the message needs wrapping
		if len(msg.text) > viewportWidth {
			// Wrap the text
			wrappedText := wrapText(msg.text, viewportWidth)
			content += msg.style.Render(wrappedText) + "\n"
		} else {
			content += msg.style.Render(msg.text) + "\n"
		}
	}
	return content
}

// wrapText wraps text to the specified width and adds a continuation indicator
func wrapText(text string, width int) string {
	if len(text) <= width {
		return text
	}

	var result string
	remaining := text

	for len(remaining) > width {
		// Take the first 'width' characters
		line := remaining[:width-1] + "↪" // Add continuation indicator
		result += line + "\n"

		// Move to the next chunk
		remaining = remaining[width-1:]
	}

	// Add the last part
	if len(remaining) > 0 {
		result += remaining
	}

	return result
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

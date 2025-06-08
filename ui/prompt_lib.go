package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbletea"
	"github.com/kjkondratuk/goblins-and-gold/ux"
)

// PromptResult is sent when a prompt is completed
type PromptResult struct {
	Value string
	Error error
}

// SelectResult is sent when a selection is made
type SelectResult struct {
	Index int
	Value string
	Error error
}

// BubbleTeaPromptLib is an implementation of ux.PromptLib using bubbletea
type BubbleTeaPromptLib struct {
	program *tea.Program
}

// NewBubbleTeaPromptLib creates a new BubbleTeaPromptLib
func NewBubbleTeaPromptLib(program *tea.Program) ux.PromptLib {
	return &BubbleTeaPromptLib{
		program: program,
	}
}

// Prompt displays a prompt and returns the user's input
func (p *BubbleTeaPromptLib) Prompt(label string, defaultValue string) (string, error) {
	// Create a channel to receive the result
	resultCh := make(chan PromptResult)

	// Send a message to the main program to show the prompt
	p.program.Send(ShowPromptMsg{
		Label:        label,
		DefaultValue: defaultValue,
		ResultCh:     resultCh,
	})

	// Wait for the result
	result := <-resultCh
	return result.Value, result.Error
}

// Select displays a selection menu and returns the user's choice
func (p *BubbleTeaPromptLib) Select(label string, items []string) (int, string, error) {
	// Create a channel to receive the result
	resultCh := make(chan SelectResult)

	// Send a message to the main program to show the select control
	p.program.Send(ShowSelectMsg{
		Label:    label,
		Items:    items,
		ResultCh: resultCh,
	})

	// Wait for the result
	result := <-resultCh
	return result.Index, result.Value, result.Error
}

// promptModel is a bubbletea model for prompts
type promptModel struct {
	label        string
	defaultValue string
	textInput    textinput.Model
	resultCh     chan<- PromptResult
}

func (m promptModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m promptModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			// Send the result and quit
			m.resultCh <- PromptResult{Value: m.textInput.Value()}
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			// Cancel and quit
			m.resultCh <- PromptResult{Error: nil}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m promptModel) View() string {
	return m.label + "\n" + m.textInput.View()
}

// selectModel is a bubbletea model for selection menus
type selectModel struct {
	label    string
	items    []string
	cursor   int
	resultCh chan<- SelectResult
}

func (m selectModel) Init() tea.Cmd {
	return nil
}

func (m selectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		case "enter":
			// Send the result and quit
			m.resultCh <- SelectResult{
				Index: m.cursor,
				Value: m.items[m.cursor],
			}
			return m, tea.Quit
		case "ctrl+c", "esc":
			// Cancel and quit
			m.resultCh <- SelectResult{Error: nil}
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m selectModel) View() string {
	s := m.label + "\n\n"

	// Show at most 3 items at a time
	maxVisible := 3

	// Calculate which items to show based on cursor
	startIdx := m.cursor - maxVisible/2
	if startIdx < 0 {
		startIdx = 0
	}

	endIdx := startIdx + maxVisible
	if endIdx > len(m.items) {
		endIdx = len(m.items)
		// Adjust startIdx if we're at the end
		startIdx = endIdx - maxVisible
		if startIdx < 0 {
			startIdx = 0
		}
	}

	// Add scroll indicator if needed
	if startIdx > 0 {
		s += "↑ more options\n"
	}

	// Render visible items
	for i := startIdx; i < endIdx; i++ {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}
		s += cursor + " " + m.items[i] + "\n"
	}

	// Add scroll indicator if needed
	if endIdx < len(m.items) {
		s += "↓ more options\n"
	}

	return s
}

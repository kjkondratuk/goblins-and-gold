package ui

import (
	"io"
	"strings"

	"github.com/charmbracelet/bubbletea"
	"github.com/pterm/pterm"
)

// Logger is a writer that sends log messages to the bubbletea program
type Logger struct {
	program *tea.Program
	level   LogLevel
	buffer  strings.Builder
}

// NewLogger creates a new logger
func NewLogger(program *tea.Program, level LogLevel) *Logger {
	return &Logger{
		program: program,
		level:   level,
	}
}

// Write implements io.Writer
func (l *Logger) Write(p []byte) (n int, err error) {
	// Write to the buffer
	n, err = l.buffer.Write(p)
	if err != nil {
		return n, err
	}

	// Check if we have a complete line
	str := l.buffer.String()
	if strings.Contains(str, "\n") {
		// Split by newlines
		lines := strings.Split(str, "\n")
		
		// Process all complete lines
		for i := 0; i < len(lines)-1; i++ {
			if lines[i] != "" {
				// Send the log message to the program
				l.program.Send(LogMsg{
					Message: lines[i],
					Level:   l.level,
				})
			}
		}
		
		// Keep the last incomplete line in the buffer
		l.buffer.Reset()
		if lines[len(lines)-1] != "" {
			l.buffer.WriteString(lines[len(lines)-1])
		}
	}

	return n, nil
}

// SetupPtermLoggers configures pterm to send logs to the bubbletea program
func SetupPtermLoggers(program *tea.Program) {
	// Create loggers for different levels
	infoLogger := NewLogger(program, LogInfo)
	successLogger := NewLogger(program, LogSuccess)
	warningLogger := NewLogger(program, LogWarning)
	errorLogger := NewLogger(program, LogError)
	debugLogger := NewLogger(program, LogDebug)

	// Set up pterm writers
	pterm.SetDefaultOutput(io.MultiWriter(infoLogger))
	pterm.Info.Writer = infoLogger
	pterm.Success.Writer = successLogger
	pterm.Warning.Writer = warningLogger
	pterm.Error.Writer = errorLogger
	pterm.Debug.Writer = debugLogger

	// Disable pterm output to stdout
	pterm.DisableStyling()
}
package ui

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kjkondratuk/goblins-and-gold/state"
)

// RoomView is a component that displays the current room description
type RoomView struct {
	state state.State
	style lipgloss.Style
}

// NewRoomView creates a new room view
func NewRoomView(s state.State) RoomView {
	return RoomView{
		state: s,
		style: lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(1).
			Bold(true),
	}
}

// Update handles messages and user input
func (r RoomView) Update(msg tea.Msg) (RoomView, tea.Cmd) {
	return r, nil
}

// View renders the room description
func (r RoomView) View() string {
	// Check if state is nil
	if r.state == nil {
		return r.style.Render("Loading...")
	}

	// Get the current room
	roomID := r.state.CurrentRoom()

	// Check if world is nil
	world := r.state.World()
	if world == nil {
		return r.style.Render("World not loaded yet...")
	}

	// Check if rooms map or specific room is nil
	if world.Rooms == nil || world.Rooms[roomID] == nil {
		return r.style.Render("Room not found...")
	}

	room := world.Rooms[roomID]

	// Create the room title
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("205")).
		Bold(true).
		Render(room.Name)

	// Create the room description
	description := room.Description

	// Combine title and description
	content := lipgloss.JoinVertical(lipgloss.Left, title, "", description)

	// Apply the style and return
	return r.style.Render(content)
}

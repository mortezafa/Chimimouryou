package cmd

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)


// Model represents the application's state
type Model struct {
	Count int  // Current count
	Done  bool // Whether the program should quit
}

// TickMsg is a custom message type for when the timer "ticks"
type TickMsg struct{}

// Init initializes the program and starts the first command
func (m Model) Init() tea.Cmd {
	return tick() // Start the timer
}

// Update handles messages and updates the model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case TickMsg: // Timer has ticked
		m.Count++ // Increment the count
		if m.Count >= 5 {
			m.Done = true // Stop after 5 seconds
			return m, nil
		}
		return m, tick() // Start the timer again

	case tea.KeyMsg: // Handle keyboard input
		if msg.String() == "q" {
			m.Done = true
		}
	}

	return m, nil
}

// View renders the UI
func (m Model) View() string {
	if m.Done {
		return fmt.Sprintf("Final Count: %d\nPress any key to exit.", m.Count)
	}
	return fmt.Sprintf("Count: %d\nPress 'q' to quit.", m.Count)
}

// tick returns a command that waits for 1 second and sends a TickMsg
func tick() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(1 * time.Second)
		return TickMsg{} // Send TickMsg to the Update function
	}
}

// main starts the program
func Main() {
	p := tea.NewProgram(Model{})
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}


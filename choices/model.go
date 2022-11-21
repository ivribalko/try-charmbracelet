package choices

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Choices  []string         // items on list
	Cursor   int              // which item our cursor is pointing at
	Selected map[int]struct{} // which items are selected
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

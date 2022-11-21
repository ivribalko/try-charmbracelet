package spinner

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	StopButton string
	quitting   bool
	err        error
	spinner.Model
}

func (m Model) Init() tea.Cmd {
	// TODO if m.Model == nil
	m.Spinner = spinner.Dot
	m.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return m.Tick
}

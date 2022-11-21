package spinner

import (
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg error

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case m.StopButton:
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case errMsg:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.Model, cmd = m.Model.Update(msg)
		return m, cmd
	}
}

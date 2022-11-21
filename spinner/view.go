package spinner

import (
	"fmt"
)

func (m Model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("\n\n   %s Loading forever...press %s to quit\n\n", m.Model.View(), m.StopButton)
	if m.quitting {
		return str + "\n"
	}
	return str
}

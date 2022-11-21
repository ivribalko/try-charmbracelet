package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ivribalko/try-charmbracelet/choices"
	"github.com/ivribalko/try-charmbracelet/spinner"
)

func main() {
	p := tea.NewProgram(spin())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func spin() spinner.Model {
	return spinner.Model{StopButton: "s"}
}

func cart() choices.Model {
	return choices.Model{
		// Our to-do list is a grocery list
		Choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		Selected: make(map[int]struct{}),
	}
}

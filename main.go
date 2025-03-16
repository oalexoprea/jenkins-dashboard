package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Create the Bubbletea program with the initial model
	p := tea.NewProgram(initialModel())

	// Start the program and handle any errors
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}

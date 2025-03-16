package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	if m.errorMsg != "" {
		return fmt.Sprintf("Error: %s", m.errorMsg)
	}

	if m.loading {
		return fmt.Sprintf("%s Loading Jenkins jobs...", m.spinner.View())
	}

	return lipgloss.NewStyle().Padding(1, 2).Render(m.list.View())
}

func (j Job) Title() string {
	return statusStyle(j.Color, j.Name)
}

func (j Job) Description() string {
	return fmt.Sprintf("Status Color: %s", j.Color)
}

func (j Job) FilterValue() string {
	return j.Name
}

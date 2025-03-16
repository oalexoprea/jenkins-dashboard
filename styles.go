package main

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	successStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Bold(true)
	failStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Bold(true)
	runningStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("4")).Bold(true)
	unknownStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
)

func statusStyle(color string, name string) string {
	switch color {
	case "blue":
		return successStyle.Render(name + " (Success)")
	case "red":
		return failStyle.Render(name + " (Failed)")
	case "yellow":
		return runningStyle.Render(name + " (Running)")
	default:
		return unknownStyle.Render(name + " (Unknown)")
	}
}

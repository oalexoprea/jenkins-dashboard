package main

import "github.com/charmbracelet/lipgloss"

// Styles for job and build status indicators
var (
	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("2")). // Green
			Bold(true)

	failStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("1")). // Red
			Bold(true)

	runningStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("3")). // Yellow
			Bold(true)

	unknownStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("7")) // Gray/white
)

// Optional: you can add more styles for the breadcrumbs bar, headers, or modals

// Breadcrumb style (optional)
var breadcrumbStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("241")).
	Padding(0, 1)

// Table header style (optional)
var tableHeaderStyle = lipgloss.NewStyle().
	Bold(true).
	Underline(true).
	Foreground(lipgloss.Color("4")) // Blue

// Modal style (if you reuse it)
var modalStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	Padding(1, 2).
	Width(80).
	Align(lipgloss.Center).
	Foreground(lipgloss.Color("229")).
	Background(lipgloss.Color("0"))

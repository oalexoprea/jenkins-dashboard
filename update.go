package main

import (
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return tea.Batch(fetchJobs(), m.spinner.Tick, tick())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "r":
			m.loading = true
			return m, fetchJobs()
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case jobsMsg:
		m.jobs = msg
		m.loading = false
		items := make([]list.Item, len(msg))
		for i, job := range msg {
			items[i] = job
		}
		m.list.SetItems(items)
		return m, nil

	case errMsg:
		m.errorMsg = string(msg)
		m.loading = false
		return m, nil

	case tickMsg:
		m.loading = true
		return m, tea.Batch(fetchJobs(), tick())
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func tick() tea.Cmd {
	return tea.Tick(time.Second*10, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

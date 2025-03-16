package main

import (
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return tea.Batch(
		fetchJobs(),    // Load initial jobs
		m.spinner.Tick, // Start spinner
		tick(),         // Periodic refresh
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Handle keyboard input
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "r":
			m.loading = true
			return m, fetchJobs()

		case "g":
			m.mode = viewGraphs
			return m, nil

		case "enter":
			selectedItem, ok := m.list.SelectedItem().(Job)
			if !ok {
				return m, nil
			}

			if selectedItem.Class == "com.cloudbees.hudson.plugins.folder.Folder" {
				m.breadcrumbs = append(m.breadcrumbs, selectedItem.URL)
				m.loading = true
				return m, fetchFolderJobs(selectedItem.URL)
			}

			m.loading = true
			return m, fetchJobBuilds(selectedItem.URL)

		case "b":
			if m.mode == viewBuilds {
				m.mode = viewJobs
				return m, nil
			}

			if len(m.breadcrumbs) == 0 {
				m.loading = true
				return m, fetchJobs()
			}

			previousURL := m.breadcrumbs[len(m.breadcrumbs)-1]
			m.breadcrumbs = m.breadcrumbs[:len(m.breadcrumbs)-1]
			m.loading = true
			return m, fetchFolderJobs(previousURL)

		case "n":
			if m.mode == viewBuilds {
				totalPages := (len(m.builds) + m.itemsPerPage - 1) / m.itemsPerPage
				if m.currentPage < totalPages-1 {
					m.currentPage++
				}
				return m, nil
			}

		case "p":
			if m.mode == viewBuilds && m.currentPage > 0 {
				m.currentPage--
				return m, nil
			}
		}

	// Spinner tick
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	// Jobs loaded
	case jobsMsg:
		m.jobs = msg
		m.loading = false

		items := make([]list.Item, len(msg))
		for i, job := range msg {
			items[i] = job
		}

		m.list.SetItems(items)
		m.list.Title = "Jenkins Jobs"
		return m, nil

	// Builds loaded
	case buildsMsg:
		m.builds = msg
		m.mode = viewBuilds
		m.currentPage = 0
		m.loading = false
		return m, nil

	// Error message
	case errMsg:
		m.errorMsg = string(msg)
		m.loading = false
		return m, nil

	// Periodic tick (auto refresh)
	case tickMsg:
		m.loading = true
		return m, tea.Batch(fetchJobs(), tick())
	}

	// Pass through to list update
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func tick() tea.Cmd {
	return tea.Tick(time.Second*10, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	if m.errorMsg != "" {
		return fmt.Sprintf("Error: %s", m.errorMsg)
	}

	if m.loading {
		return fmt.Sprintf("%s Loading...", m.spinner.View())
	}

	breadcrumbBar := lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")).
		Render(breadcrumbsView(m.breadcrumbs))

	var content string
	if m.mode == viewJobs {
		content = m.list.View()
	} else if m.mode == viewBuilds {
		content = m.buildsTableView()
	}

	return fmt.Sprintf("%s\n\n%s", breadcrumbBar, content)
}

func breadcrumbsView(breadcrumbs []string) string {
	if len(breadcrumbs) == 0 {
		return "📂 /"
	}

	view := "📂 /"
	for i, crumb := range breadcrumbs {
		if i != 0 {
			view += " ➜ "
		}
		view += extractFolderName(crumb)
	}
	return view
}

func extractFolderName(url string) string {
	parts := strings.Split(strings.TrimRight(url, "/"), "/")
	if len(parts) == 0 {
		return ""
	}
	return parts[len(parts)-1]
}

func (m model) buildsTableView() string {
	if len(m.builds) == 0 {
		return "No builds found."
	}

	itemsPerPage := m.itemsPerPage
	totalPages := (len(m.builds) + itemsPerPage - 1) / itemsPerPage
	start := m.currentPage * itemsPerPage
	end := start + itemsPerPage

	if start >= len(m.builds) {
		return "No builds on this page."
	}

	if end > len(m.builds) {
		end = len(m.builds)
	}

	visibleBuilds := m.builds[start:end]

	header := fmt.Sprintf("%-10s %-10s %-10s %-20s %s", "Number", "Result", "Duration", "Date", "URL")
	separator := strings.Repeat("-", len(header))

	rows := []string{header, separator}

	for _, build := range visibleBuilds {
		t := time.UnixMilli(build.Timestamp)

		// Colorize result
		resultText := colorizeBuildResult(build.Result)

		rows = append(rows, fmt.Sprintf(
			"%-10d %-10s %-10ds %-20s %s",
			build.Number,
			resultText,
			build.Duration/1000,
			t.Format("2006-01-02 15:04:05"),
			build.URL,
		))
	}

	footer := fmt.Sprintf("\nPage %d/%d [n]ext, [p]rev, [b]ack", m.currentPage+1, totalPages)

	return strings.Join(rows, "\n") + footer
}

func (j Job) Title() string {
	icon := ""
	label := j.Name

	if j.Class == "com.cloudbees.hudson.plugins.folder.Folder" {
		icon = "📂"
	} else {
		switch j.Color {
		case "blue":
			icon = "✅"
		case "red":
			icon = "❌"
		case "yellow":
			icon = "⚠️"
		default:
			icon = "🔧"
		}
	}

	return fmt.Sprintf("%s %s", icon, label)
}

func (j Job) Description() string {
	return "" // Clean, no extra description
}

func (j Job) FilterValue() string {
	return j.Name
}

func colorizeBuildResult(result string) string {
	switch result {
	case "SUCCESS":
		return successStyle.Render(result)
	case "FAILURE":
		return failStyle.Render(result)
	case "UNSTABLE":
		return runningStyle.Render(result)
	default:
		return unknownStyle.Render(result)
	}
}

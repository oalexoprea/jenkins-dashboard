package main

import (
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
)

// View modes to distinguish between job and build views
type viewMode int

const (
	viewJobs viewMode = iota
	viewBuilds
	viewGraphs
)

// Job model for Jenkins jobs (folders or pipelines)
type Job struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Class string `json:"_class"`
	URL   string `json:"url"`
}

// Build model for Jenkins builds
type Build struct {
	Number    int    `json:"number"`
	Result    string `json:"result"`
	Timestamp int64  `json:"timestamp"`
	Duration  int64  `json:"duration"`
	URL       string `json:"url"`
}

// Response model for Jenkins job API
type JenkinsResponse struct {
	Jobs []Job `json:"jobs"`
}

// Main application model
type model struct {
	jobs         []Job
	builds       []Build
	list         list.Model
	spinner      spinner.Model
	mode         viewMode
	loading      bool
	errorMsg     string
	breadcrumbs  []string
	currentPage  int
	itemsPerPage int
}

// Initial model setup
func initialModel() model {
	sp := spinner.New()
	sp.Spinner = spinner.Dot

	delegate := list.NewDefaultDelegate()
	ls := list.New(nil, delegate, 50, 20)
	ls.Title = "Jenkins Jobs"

	return model{
		list:         ls,
		spinner:      sp,
		mode:         viewJobs,
		loading:      true,
		itemsPerPage: 5,
		currentPage:  0,
	}
}

// Messages exchanged between update and commands
type jobsMsg []Job
type buildsMsg []Build
type errMsg string
type tickMsg time.Time

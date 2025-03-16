package main

import (
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
)

type Job struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type JenkinsResponse struct {
	Jobs []Job `json:"jobs"`
}

type model struct {
	jobs     []Job
	list     list.Model
	spinner  spinner.Model
	loading  bool
	errorMsg string
}

func initialModel() model {
	sp := spinner.New()
	sp.Spinner = spinner.Dot

	delegate := list.NewDefaultDelegate()
	ls := list.New(nil, delegate, 50, 20)
	ls.Title = "Jenkins Jobs"

	return model{
		jobs:     []Job{},
		list:     ls,
		spinner:  sp,
		loading:  true,
		errorMsg: "",
	}
}

type jobsMsg []Job
type errMsg string
type tickMsg time.Time

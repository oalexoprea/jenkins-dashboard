package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func fetchJobs() tea.Cmd {
	return fetchFolderJobs(os.Getenv("JENKINS_URL"))
}

func fetchFolderJobs(folderURL string) tea.Cmd {
	return func() tea.Msg {
		fullURL := fmt.Sprintf("%s/api/json", strings.TrimRight(folderURL, "/"))

		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return errMsg(err.Error())
		}

		user := os.Getenv("JENKINS_USER")
		token := os.Getenv("JENKINS_TOKEN")
		req.SetBasicAuth(user, token)

		client := &http.Client{Timeout: 10 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			return errMsg(err.Error())
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errMsg(err.Error())
		}

		var jResp JenkinsResponse
		err = json.Unmarshal(body, &jResp)
		if err != nil {
			return errMsg(err.Error())
		}

		return jobsMsg(jResp.Jobs)
	}
}

func fetchJobBuilds(jobURL string) tea.Cmd {
	return func() tea.Msg {
		fullURL := fmt.Sprintf("%s/api/json?tree=builds[number,result,timestamp,duration,url]", strings.TrimRight(jobURL, "/"))

		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return errMsg(err.Error())
		}

		user := os.Getenv("JENKINS_USER")
		token := os.Getenv("JENKINS_TOKEN")
		req.SetBasicAuth(user, token)

		client := &http.Client{Timeout: 10 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			return errMsg(err.Error())
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errMsg(err.Error())
		}

		var jobData struct {
			Builds []Build `json:"builds"`
		}

		err = json.Unmarshal(body, &jobData)
		if err != nil {
			return errMsg(err.Error())
		}

		return buildsMsg(jobData.Builds)
	}
}

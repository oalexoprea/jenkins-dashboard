package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func fetchJobs() tea.Cmd {
	return func() tea.Msg {
		jenkinsURL := os.Getenv("JENKINS_URL")
		user := os.Getenv("JENKINS_USER")
		token := os.Getenv("JENKINS_TOKEN")

		if jenkinsURL == "" || user == "" || token == "" {
			return errMsg("Missing Jenkins credentials in env vars")
		}

		req, err := http.NewRequest("GET", jenkinsURL+"/api/json", nil)
		if err != nil {
			return errMsg(err.Error())
		}
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

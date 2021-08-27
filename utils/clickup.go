package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var api = "https://api.clickup.com/api/v1"

func FetchTasks(teamID, token string, taskIDs []string) ([]Task, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/team/%s/task", api, teamID), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", token)

	q := req.URL.Query()
	for _, taskID := range taskIDs {
		q.Add("task_ids[]", taskID)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	clickupResp := &ClickupResp{}
	err = json.Unmarshal(body, clickupResp)
	if err != nil {
		return nil, err
	}

	return clickupResp.Tasks, nil
}

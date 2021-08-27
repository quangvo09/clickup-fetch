package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/quangvo09/clickup-fetch/utils"
)

var (
	CLICKUP_TEAM_ID = "CLICKUP_TEAM_ID"
	CLICKUP_TOKEN   = "CLICKUP_TOKEN"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Correct Syntax: ./clickup-fetch task_id1, task_id2, ...")
		panic("Please provide at least one task_id")
	}

	teamID := os.Getenv(CLICKUP_TEAM_ID)
	token := os.Getenv(CLICKUP_TOKEN)

	tasks, err := utils.FetchTasks(teamID, token, args)
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data)
}

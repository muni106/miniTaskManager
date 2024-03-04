package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	taskName string `json:"task_name"`
	isMade   bool   `json:"made"`
}

func writeTask() {

}

func checkFile() {
	jsonFile, err := os.Create("/tasks.json")

	if err != nil {
		fmt.Println(err)
	}

	byteValue, err := io.ReadAll(jsonFile)

	var tasks Tasks

	err = json.Unmarshal(byteValue, &tasks)

	//close file
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {

		}
	}(jsonFile)
}

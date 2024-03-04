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
	Task string `json:"task"`
	Made bool   `json:"made"`
}

const filename = "tasks.json"

func tasksNames() {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened tasks.json")
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)

	var tasks Tasks

	err = json.Unmarshal(byteValue, &tasks)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(tasks.Tasks); i++ {
		fmt.Println("ciao "+tasks.Tasks[i].Task, i)
	}
}

func addATask(taskName string) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened tasks.json")
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)

	var tasks Tasks
	err = json.Unmarshal(byteValue, &tasks)

	task := Task{Task: taskName, Made: false}

	tasks.Tasks = append(tasks.Tasks, task)

	marshall, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Error during Marshaling: ", err)
		return
	}

	err = os.WriteFile(filename, marshall, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

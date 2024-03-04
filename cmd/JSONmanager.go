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

func tasksNames() {
	jsonFile, err := os.Open("tasks.json")
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

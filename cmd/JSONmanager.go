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

func taskNames() {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened tasks.json")
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(filename + " read went wrong")
	}
	var tasks Tasks

	err = json.Unmarshal(byteValue, &tasks)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(tasks.Tasks); i++ {
		fmt.Println(i, ". ", tasks.Tasks[i].Task)
	}
}

func tasksTodo() {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened tasks.json")
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		fmt.Println("Error in file reading: ", err)
		return
	}

	var tasks Tasks

	err = json.Unmarshal(byteValue, &tasks)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(tasks.Tasks); i++ {
		if !tasks.Tasks[i].Made {
			fmt.Println("- ", tasks.Tasks[i].Task)
		}
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
	if err != nil {
		fmt.Println(filename + " read went wrong")
	}

	var tasks Tasks
	err = json.Unmarshal(byteValue, &tasks)
	if err != nil {
		fmt.Println("unmarshall went wrong", err)
	}

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
	fmt.Println("'" + taskName + "' task has been successfully added")
}

func setTaskCompleteByIndex(i int) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened tasks.json")
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(filename + " read went wrong")
	}
	var tasks Tasks

	err = json.Unmarshal(byteValue, &tasks)
	if err != nil {
		fmt.Println(err)
	}

	tasks.Tasks[i].Made = true

	marshall, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("fail on marshaling: ", err)
	}

	err = os.WriteFile(filename, marshall, 0644)
	if err != nil {
		fmt.Println("error on file writing")
	}

	fmt.Println("task : '", tasks.Tasks[i].Task, "' completed")

}

func setTaskCompleteByName(taskName string) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened tasks.json")
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(filename + " read went wrong")
	}
	var tasks Tasks

	err = json.Unmarshal(byteValue, &tasks)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(tasks.Tasks); i++ {
		if tasks.Tasks[i].Task == taskName {
			tasks.Tasks[i].Made = true
			fmt.Println("task : '", tasks.Tasks[i].Task, "' completed")
			marshall, err := json.Marshal(tasks)
			if err != nil {
				fmt.Println("fail on marshaling: ", err)
			}

			err = os.WriteFile(filename, marshall, 0644)
			if err != nil {
				fmt.Println("error on file writing")
			}
			return
		}
	}
	fmt.Println("task not found")
}


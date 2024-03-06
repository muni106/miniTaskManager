package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new task",
	Long:  "add a new task to all your tasks, then 'list' them too see em all",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			fmt.Println("bro we need at least an argument")
			return nil
		} else if len(args) > 2 {
			fmt.Println("bro there are too many arguments")
		}
		taskName := args[0]
		addATask(taskName)
		return nil
	},
}

func Hello() {
	fmt.Println("Hello")
}

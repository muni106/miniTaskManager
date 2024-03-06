package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove permanently an element from the list of elements",
	Long:  "You can remove an element by passing the task name or task index",
	RunE: func(cmd *cobra.Command, args []string) error {
		all, err := cmd.Flags().GetBool("all")
		if all && err == nil {
			removeAllTasks()
			return nil
		}
		item, err := strconv.Atoi(args[0])
		if err != nil {
			removeTaskByName(args[0])
			return nil
		}
		removeTaskByIndex(item)
		return nil
	},
}

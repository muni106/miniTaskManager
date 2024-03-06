package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Set a task complete",
	Long: `using the index of a task it sets it as completed, 
			its the equivalent of a check in a todo list`,
	RunE: func(cmd *cobra.Command, args []string) error {
		item, err := strconv.Atoi(args[0])
		if err != nil {
			setTaskCompleteByName(args[0])
			return nil
		}
		setTaskCompleteByIndex(item)
		return nil
	},
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listTasks = &cobra.Command{
	Use:   "list",
	Short: "list the tasks",
	Long:  "this command is usefull to list the tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		all, err := cmd.Flags().GetBool("all")
		if err != nil {
			fmt.Println("error on getting all flag: ", err)
		}
		if all {
			taskNames()
		} else {
			tasksTodo()
		}
		return nil
	},
}

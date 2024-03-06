package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mitama",
	Short: "mitama is a MIni TAsk MAnager",
	Long: ` Task Manager which can add a task
				complete it, list all tasks and remove tasks`,
}

func init() {
	greetCmd.Flags().StringP("name", "n", "", "Name of the person to greet")
	greetCmd.Flags().StringP("message", "m", "Hello", "Custom message to include in the greeting")
	rootCmd.AddCommand(greetCmd)
	rootCmd.AddCommand(addCmd)
	listCmd.Flags().BoolP("all", "a", false, "see all tasks comleted and uncompleted")
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(completeCmd)
	removeCmd.Flags().BoolP("all", "a", false, "remove all tasks")
	rootCmd.AddCommand(removeCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

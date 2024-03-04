package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
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
	rootCmd.AddCommand(addTask)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Greet someone with a pers message",
	Long:  "yaaaaa nfrkjfekj ewfjknewf",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Access flag values
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		message, err := cmd.Flags().GetString("message")
		if err != nil {
			return err
		}
		fmt.Printf("Hello, %s! %s\n", name, message)
		return nil
	},
}

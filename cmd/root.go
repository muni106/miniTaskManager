package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "basic",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
}

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

func init() {
	greetCmd.Flags().StringP("name", "n", "", "Name of the person to greet")
	greetCmd.Flags().StringP("message", "m", "Hello", "Custom message to include in the greeting")
	rootCmd.AddCommand(greetCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
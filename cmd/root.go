package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "reffactor",
	Short: "Improve your code",
	Long:  "Improve your code",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(0)
	}
}

package cmd

import (
	"os"

	"github.com/Guilherme-Matosoli/Reffactor/internal/handlers"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "reffactor",
	Short: "Improve your code",
	Long:  "Improve your code",
	Run:   handlers.Configurate,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(0)
	}
}

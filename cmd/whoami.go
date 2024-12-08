package cmd

import (
	"os/user"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var whoami = &cobra.Command{
	Use:   "whoami",
	Short: "Display user name",
	Long:  "Display current user name",
	Run: func(cmd *cobra.Command, args []string) {
		usr, err := user.Current()
		if err != nil {
			color.Red("Error geting user information: ", err)
			return
		}

		color.Green("You are logged as %s", usr.Name)
	},
}

func init() {
	rootCmd.AddCommand(whoami)
}

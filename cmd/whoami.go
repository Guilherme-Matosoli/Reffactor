package cmd

import (
	"fmt"
	"os/user"

	"github.com/spf13/cobra"
)

var whoami = &cobra.Command{
	Use:   "whoami",
	Short: "Display user name",
	Long:  "Display current user name",
	Run: func(cmd *cobra.Command, args []string) {
		usr, err := user.Current()
		if err != nil {
			fmt.Println("Error geting user information: ", err)
			return
		}

		fmt.Println("You are logged as ", usr)
	},
}

func init() {
	rootCmd.AddCommand(whoami)
}

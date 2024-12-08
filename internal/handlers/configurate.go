package handlers

import (
	"errors"
	"fmt"

	"github.com/Guilherme-Matosoli/Reffactor/internal/utils"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/spf13/cobra"
)

func Configurate(cmd *cobra.Command, args []string) {
	utils.ShowLogo()
	var apiKey string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewText().
				Title("Set your OpenAI api key: ").
				Value(&apiKey).
				Validate(func(s string) error {
					spinner.
						New().
						Title("Verifying your api key...").
						Run()

					err := utils.VerifyKey(apiKey)
					if err != nil {
						return errors.New("Your key is invalid, please set a valid key and try again!")
					}

					return nil
				}),
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Print(err)
	}

}

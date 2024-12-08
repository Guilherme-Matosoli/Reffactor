package handlers

import (
	"fmt"

	// "github.com/Guilherme-Matosoli/Reffactor/internal/utils"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/spf13/cobra"
)

var asciiArt = `
▗▖ ▗▖▗▞▀▚▖█ ▗▞▀▘ ▄▄▄  ▄▄▄▄  ▗▞▀▚▖       ■   ▄▄▄      ▗▄▄▖ ▗▞▀▚▖▗▞▀▀▘▗▞▀▀▘▗▞▀▜▌▗▞▀▘   ■   ▄▄▄   ▄▄▄ 
▐▌ ▐▌▐▛▀▀▘█ ▝▚▄▖█   █ █ █ █ ▐▛▀▀▘    ▗▄▟▙▄▖█   █     ▐▌ ▐▌▐▛▀▀▘▐▌   ▐▌   ▝▚▄▟▌▝▚▄▖▗▄▟▙▄▖█   █ █    
▐▌ ▐▌▝▚▄▄▖█     ▀▄▄▄▀ █   █ ▝▚▄▄▖      ▐▌  ▀▄▄▄▀     ▐▛▀▚▖▝▚▄▄▖▐▛▀▘ ▐▛▀▘            ▐▌  ▀▄▄▄▀ █    
▐▙█▟▌     █                            ▐▌            ▐▌ ▐▌     ▐▌   ▐▌              ▐▌             
                                       ▐▌                                           ▐▌`

func Configurate(cmd *cobra.Command, args []string) {
	fmt.Print("\033[35m", asciiArt)
	var apiKey string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewText().
				Title("\n\nSet your OpenAI api key: ").
				Value(&apiKey),
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Print(err)
	}

	spinner.New().Title("Verifying your api key...").Run()
}

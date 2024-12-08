package handlers

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var asciiArt = `
▗▖ ▗▖▗▞▀▚▖█ ▗▞▀▘ ▄▄▄  ▄▄▄▄  ▗▞▀▚▖       ■   ▄▄▄      ▗▄▄▖ ▗▞▀▚▖▗▞▀▀▘▗▞▀▀▘▗▞▀▜▌▗▞▀▘   ■   ▄▄▄   ▄▄▄ 
▐▌ ▐▌▐▛▀▀▘█ ▝▚▄▖█   █ █ █ █ ▐▛▀▀▘    ▗▄▟▙▄▖█   █     ▐▌ ▐▌▐▛▀▀▘▐▌   ▐▌   ▝▚▄▟▌▝▚▄▖▗▄▟▙▄▖█   █ █    
▐▌ ▐▌▝▚▄▄▖█     ▀▄▄▄▀ █   █ ▝▚▄▄▖      ▐▌  ▀▄▄▄▀     ▐▛▀▚▖▝▚▄▄▖▐▛▀▘ ▐▛▀▘            ▐▌  ▀▄▄▄▀ █    
▐▙█▟▌     █                            ▐▌            ▐▌ ▐▌     ▐▌   ▐▌              ▐▌             
                                       ▐▌                                           ▐▌`

func Configurate(cmd *cobra.Command, args []string) {
	fmt.Print("\033[32m", asciiArt)
	var option string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose option").
				Options(
					huh.NewOption("a", "a"),
					huh.NewOption("b", "b"),
					huh.NewOption("c", "c"),
					huh.NewOption("d", "d"),
				).
				Value(&option),
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(form)
}

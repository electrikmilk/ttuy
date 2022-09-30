/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

var optionSelected int
var lastOptionSelected = -1

var lastPrompt string

// Prompt prints a prompt for the user to choose yes or no using the arrow keys then returns a boolean based on the option chosen
func Prompt(prompt string) (proceed bool) {
	fmt.Print(eol())
	go readKeys(handlePromptKeys)
	painter(func() (template string) {
		if optionSelected != lastOptionSelected {
			template = Style(prompt, Bold, Green) + eol()
			if optionSelected == 1 {
				template += Style(" Yes ", Inverted, Bold)
				proceed = true
			} else {
				template += Style(" Yes ", Bold)
			}
			template += "\t"
			if optionSelected == 0 {
				template += Style(" No ", Inverted, Bold)
				proceed = false
			} else {
				template += Style(" No ", Bold)
			}
			template += eols(2) + Style(leftArrow+" "+rightArrow+" Use Left & Right Arrow Keys to Choose.", Dim)
			lastOptionSelected = optionSelected
			lastPrompt = template
		} else {
			template = lastPrompt
		}
		return
	})
	fmt.Print(eol())
	return
}

func handlePromptKeys(key any) {
	switch key {
	case keyboard.KeyArrowLeft:
		if optionSelected != 1 {
			optionSelected = 1
		}
	case keyboard.KeyArrowRight:
		if optionSelected != 0 {
			optionSelected = 0
		}
	case keyboard.KeyEnter:
		stopPainting()
		return
	}
}

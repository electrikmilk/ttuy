/*
 * Copyright (c) 2023 Brandon Jordan
 */

package ttuy

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"os"
	"strings"
)

var proceed bool
var lastProceed = true

var lastPrompt string

// Prompt prints a prompt for the user to choose yes or no using the arrow keys then returns a boolean based on the option chosen
func Prompt(prompt string) bool {
	fmt.Print(eol)
	go ReadKeys(handlePromptKeys)
	Painter(func() (template string) {
		if proceed == lastProceed {
			template = lastPrompt
			return
		}

		template = Style(prompt, Bold, GreenText) + eol
		if proceed {
			template += Style(" Yes ", Inverted, Bold)
		} else {
			template += Style(" Yes ", Bold)
		}
		template += "\t"
		if !proceed {
			template += Style(" No ", Inverted, Bold)
		} else {
			template += Style(" No ", Bold)
		}
		template += strings.Repeat(eol, 2) + Style(leftArrow+" "+rightArrow+" Use Left & Right Arrow Keys to Choose.", Dim, Italic)
		lastProceed = proceed
		lastPrompt = template

		return
	})
	fmt.Print(eol)
	return proceed
}

func handlePromptKeys(key keyboard.Key) {
	switch key {
	case keyboard.KeyArrowLeft, keyboard.KeyArrowRight:
		proceed = !proceed
	case keyboard.KeyCtrlC:
		CursorShow()
		os.Exit(0)
	case keyboard.KeyEnter:
		StopPainting()
		return
	}
}

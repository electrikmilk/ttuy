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
var promptMenu = true

func Prompt(prompt string) (proceed bool) {
	cursorHide()
	fmt.Print(eol() + Style(prompt, Bold, Green) + eol())
	go readKeys(handlePromptKeys)
	fmt.Print(eol())
	for i := 0; i < 2; i++ {
		fmt.Print(eol())
	}
	linePrev(2)
	for {
		if !promptMenu {
			break
		}
		if optionSelected != lastOptionSelected {
			cursorStart()
			clearLine()
			if optionSelected == 1 {
				fmt.Print(Style(" Yes ", Inverted, Bold))
			} else {
				fmt.Print(Style(" Yes ", Bold))
			}
			fmt.Print("\t")
			if optionSelected == 0 {
				fmt.Print(Style(" No ", Inverted, Bold))
			} else {
				fmt.Print(Style(" No ", Bold))
			}
			lastOptionSelected = optionSelected
		}
	}
	lineNext(1)
	fmt.Print(eol())
	cursorShow()
	if optionSelected == 1 {
		proceed = true
	} else if optionSelected == 0 {
		proceed = false
	}
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
		promptMenu = false
	}
}

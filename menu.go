/*
 * Copyright (c) 2023 Brandon Jordan
 */

package ttuy

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
)

var menuTitle string
var menuOptions []Option

var optionIndex int
var optionCount int

var lastMenuContent string
var lastOptionIndex = -1

const selected = "[x]"
const notSelected = "[ ]"

type Option struct {
	Label    string
	Callback func()
	Disabled bool
}

// Menu displays a menu of options and triggers the callback corresponding to the option chosen
func Menu(title string, options []Option) {
	menuTitle = title
	menuOptions = options
	optionIndex = 0
	lastOptionIndex = -1
	lastMenuContent = ""
	optionCount = len(menuOptions)
	firstAvailable()
	go ReadKeys(handleMenuKeys)
	Painter(func() (menu string) {
		if lastOptionIndex != optionIndex {
			menu = Style(menuTitle, Bold, GreenText) + eol
			for i, option := range menuOptions {
				switch {
				case option.Disabled:
					menu += Style(fmt.Sprintf(notSelected+" %s"+eol, option.Label), Dim)
				case i == optionIndex:
					menu += Style(selected, MagentaText) + fmt.Sprintf(" %s"+eol, option.Label)
				default:
					menu += fmt.Sprintf(notSelected+" %s"+eol, option.Label)
				}
			}
			lastMenuContent = menu
			lastOptionIndex = optionIndex
		} else {
			menu = lastMenuContent
		}
		return
	})
}

func firstAvailable() {
	if !menuOptions[0].Disabled {
		return
	}
	for _, option := range menuOptions {
		if !option.Disabled {
			continue
		}

		optionIndex++
	}
}

func handleMenuKeys(key keyboard.Key) {
	switch key {
	case keyboard.KeyCtrlC:
		StopPainting()
		CursorShow()
		os.Exit(0)
	case keyboard.KeyArrowUp:
		var nextIndex = optionIndex - 1
		var nextOptionDisabled = true
		for nextOptionDisabled {
			if nextIndex > -1 {
				nextOptionDisabled = menuOptions[nextIndex].Disabled
				if !nextOptionDisabled {
					optionIndex = nextIndex
					break
				}
			} else {
				break
			}
			nextIndex--
		}
	case keyboard.KeyArrowDown:
		var nextIndex = optionIndex + 1
		var nextOptionDisabled = true
		for nextOptionDisabled {
			if nextIndex < optionCount {
				nextOptionDisabled = menuOptions[nextIndex].Disabled
				if !nextOptionDisabled {
					optionIndex = nextIndex
					break
				}
			} else {
				break
			}
			nextIndex++
		}
	case keyboard.KeyEnter:
		selectedOption := menuOptions[optionIndex]
		selectedOption.Callback()
		StopPainting()
		return
	}
}

/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

var menuTitle string
var menuOptions []Option

var optionIndex int
var optionCount int

var lastMenuContent string
var lastOptionIndex = -1

var selected = "[x]"
var notSelected = "[ ]"

type Option struct {
	Label    string
	Callback func()
	Disabled bool
}

func Menu(title string, options []Option) {
	menuTitle = title
	menuOptions = options
	optionIndex = 0
	lastOptionIndex = -1
	lastMenuContent = ""
	optionCount = len(menuOptions)
	go readKeys(handleMenuKeys)
	firstAvailable()
	painter(func() (menu string) {
		if lastOptionIndex != optionIndex {
			menu = Style(menuTitle, Bold, Green) + eol()
			for i, option := range menuOptions {
				switch {
				case option.Disabled:
					menu += Style(fmt.Sprintf(notSelected+" %s"+eol(), option.Label), Dim)
				case i == optionIndex:
					menu += Style(selected, Magenta) + fmt.Sprintf(" %s"+eol(), option.Label)
				default:
					menu += fmt.Sprintf(notSelected+" %s"+eol(), option.Label)
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
	if menuOptions[0].Disabled {
		for _, option := range menuOptions {
			if option.Disabled {
				optionIndex++
			}
		}
	}
}

func handleMenuKeys(key any) {
	switch key {
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
		stopPainting()
		selectedOption := menuOptions[optionIndex]
		selectedOption.Callback()
		return
	}
}

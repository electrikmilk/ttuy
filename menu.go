/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"fmt"
	"strings"

	"github.com/eiannone/keyboard"
)

var content string
var lineCount int
var lines []string

var lastContent string
var lastLineCount int
var lastLines []string
var lastMenuContent string
var lastOptionIndex = -1

var diff int

var menuTitle string
var menuOptions []Option

var optionIndex int
var optionCount int

var selected = "[x]"
var drawing = true

type Option struct {
	Label    string
	Callback func()
	Disabled bool
}

func Menu(title string, options []Option) {
	menuTitle = title
	menuOptions = options
	drawing = true
	setup()
	go readKeys(handleMenuKeys)
	for {
		if !drawing {
			break
		}
		if currentLine != 0 {
			backUp()
		}
		content = template()
		if content != lastContent {
			splitCount()
			paintMenu()
			last()
		}
	}
	fmt.Println(content)
	cursorShow()
}

func setup() {
	cursorHide()
	firstAvailable()
	optionCount = len(menuOptions)
	content = template()
	splitCount()
	makeRoom()
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

func makeRoom() {
	for i := 0; i < lineCount; i++ {
		fmt.Print(eol())
	}
	linePrev(lineCount)
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
		drawing = false
		selectedOption := menuOptions[optionIndex]
		selectedOption.Callback()
	}
}

func template() (menu string) {
	if lastOptionIndex != optionIndex {
		menu = Style(menuTitle, Bold, Green) + eol()
		for i, option := range menuOptions {
			switch {
			case option.Disabled:
				menu += Style(fmt.Sprintf("[ ] %s"+eol(), option.Label), Dim)
			case i == optionIndex:
				menu += Style(selected, Magenta) + fmt.Sprintf(" %s"+eol(), option.Label)
			default:
				menu += fmt.Sprintf("[ ] %s"+eol(), option.Label)
			}
		}
		lastMenuContent = menu
		lastOptionIndex = optionIndex
	} else {
		menu = lastMenuContent
	}
	return
}

func paintMenu() {
	for i, line := range lines {
		if len(lastLines) > i {
			if lastLines[i] != lines[i] {
				replaceLine(line)
			}
		} else {
			replaceLine(line)
		}
		lineNext(1)
	}
	if lastLineCount != 0 {
		if lastLineCount > lineCount {
			diff = lastLineCount - lineCount
			for i := 0; i < diff; i++ {
				clearLine()
				lineNext(1)
			}
		}
	}
}

func replaceLine(line string) {
	clearLine()
	fmt.Printf("%s\r", line)
}

func splitCount() {
	lines = strings.Split(content, eol())
	lineCount = len(lines)
}

func last() {
	lastContent = content
	lastLineCount = lineCount
	lastLines = lines
}

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

var diff int

var menuTitle string
var menuOptions []Option

var optionIndex int
var optionCount int

var selected = "[x]"
var drawing bool = true

type Option struct {
	Label    string
	Callback func()
	Disabled bool
}

func Menu(title string, options []Option) {
	menuTitle = title
	menuOptions = options
	drawing = true
	cursorHide()
	content = template()
	splitCount()
	makeRoom()
	firstAvailable()
	go readKeys(func(key any) {
		switch key {
		case keyboard.KeyArrowUp:
			var nextIndex = optionIndex - 1
			var nextOptionDisabled bool = true
			for nextOptionDisabled == true {
				if nextIndex > -1 {
					nextOptionDisabled = menuOptions[nextIndex].Disabled
					if nextOptionDisabled == false {
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
			var nextOptionDisabled bool = true
			for nextOptionDisabled == true {
				if nextIndex < optionCount {
					nextOptionDisabled = menuOptions[nextIndex].Disabled
					if nextOptionDisabled == false {
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
			drawing = false
		}
	})
	for {
		if drawing == false {
			break
		}
		if currentLine != 0 {
			backUp()
		}
		content = template()
		splitCount()
		if content != lastContent {
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
		last()
	}
	fmt.Println(content)
	cursorShow()
}

func template() (menu string) {
	optionCount = len(menuOptions)
	menu = Style(menuTitle, BOLD, GREEN) + "\n"
	for i, option := range menuOptions {
		if option.Disabled == true {
			menu += Style(fmt.Sprintf("[ ] %s\n", option.Label), DIM)
		} else if i == optionIndex {
			menu += Style(selected, MAGENTA) + fmt.Sprintf(" %s\n", option.Label)
		} else {
			menu += fmt.Sprintf("[ ] %s\n", option.Label)
		}
	}
	return
}

func firstAvailable() {
	if menuOptions[0].Disabled == true {
		for _, option := range menuOptions {
			if option.Disabled == true {
				optionIndex++
			}
		}
	}
}

func makeRoom() {
	for i := 0; i < lineCount; i++ {
		fmt.Printf("\n")
	}
	linePrev(lineCount)
}

func replaceLine(line string) {
	clearLine()
	fmt.Printf("%s\r", line)
}

func splitCount() {
	lines = strings.Split(content, "\n")
	lineCount = len(lines)
}

func last() {
	lastContent = content
	lastLineCount = lineCount
	lastLines = lines
}

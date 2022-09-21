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
}

func Menu(title string, options []Option) {
	cursorHide()
	menuTitle = title
	menuOptions = options
	drawing = true
	go readKeys(func(key any) {
		switch key {
		case keyboard.KeyArrowUp:
			var newIndex = optionIndex - 1
			if newIndex <= 0 {
				optionIndex = 0
			} else {
				optionIndex = newIndex
			}
		case keyboard.KeyArrowDown:
			var newIndex = optionIndex + 1
			if newIndex < optionCount {
				optionIndex = newIndex
			}
		case keyboard.KeyEnter:
			selectedOption := menuOptions[optionIndex]
			selectedOption.Callback()
			cursorShow()
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
		split_count()
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
		if i == optionIndex {
			menu += Style(selected, MAGENTA) + fmt.Sprintf(" %s\n", option.Label)
		} else {
			menu += fmt.Sprintf("[ ] %s\n", option.Label)
		}
	}
	return
}

func replaceLine(line string) {
	clearLine()
	fmt.Printf("%s\r", line)
}

func split_count() {
	lines = strings.Split(content, "\n")
	lineCount = len(lines) - 1
}

func last() {
	lastContent = content
	lastLineCount = lineCount
	lastLines = lines
}

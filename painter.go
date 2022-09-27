/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"fmt"
	"strings"
)

var content string
var lineCount int
var lines []string

var lastContent string
var lastLineCount int
var lastLines []string

var diff int

type template func() string

var stopPaint = make(chan bool)

func painter(callback template) {
	content = callback()
	setup()
	for {
		select {
		case <-stopPaint:
			fmt.Println(content)
			cursorShow()
			return
		default:
			if currentLine != 0 {
				backUp()
			}
			content = callback()
			if content != lastContent {
				splitCount()
				paint()
				last()
			}
		}
	}
}

func setup() {
	lastContent = ""
	lastLineCount = 0
	lastLines = []string{}
	cursorHide()
	splitCount()
	makeRoom()
}

func makeRoom() {
	for i := 0; i < lineCount; i++ {
		fmt.Print(eol())
	}
	linePrev(lineCount)
}

func paint() {
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

func stopPainting() {
	stopPaint <- true
}

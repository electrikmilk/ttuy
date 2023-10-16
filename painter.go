/*
 * Copyright (c) 2023 Brandon Jordan
 */

package ttuy

import (
	"fmt"
	"strings"
	"time"
)

var content string
var lineCount int
var lines []string

var lastContent string
var lastLineCount int

var diff int

type Template func() string

var stopPaint = make(chan bool)

// Painter prints and updates content received from Template
// Learn more: https://github.com/electrikmilk/ttuy/wiki/Painter()
func Painter(callback Template) {
	lastContent = ""
	lastLineCount = 0
	currentLine = 0
	content = callback()
	CursorHide()
	splitCount()
	for {
		select {
		case <-stopPaint:
			fmt.Println(content)
			CursorShow()
			return
		default:
			if currentLine != 0 {
				BackUp()
			}
			content = callback()
			if content != lastContent {
				splitCount()
				paint()
				storeLast()
			}
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func makeRoom() {
	fmt.Print(strings.Repeat(eol, lineCount))
	LinePrev(lineCount)
}

func paint() {
	if lastLineCount < lineCount {
		makeRoom()
	}
	for _, line := range lines {
		replaceLine(line)
		LineNext(1)
	}
	if lastLineCount == 0 || lastLineCount < lineCount {
		return
	}
	diff = lastLineCount - lineCount
	for i := 0; i < diff; i++ {
		ClearLine()
		LineNext(1)
	}
}

func replaceLine(line string) {
	ClearLine()
	fmt.Print(line)
}

func splitCount() {
	lines = strings.Split(content, eol)
	lineCount = len(lines)
}

func storeLast() {
	lastContent = content
	lastLineCount = lineCount
}

// StopPainting triggers current usage of Painter to stop
func StopPainting() {
	stopPaint <- true
}

/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"fmt"
	"strings"

	"github.com/eiannone/keyboard"
	"github.com/muesli/reflow/wordwrap"
)

var contents string
var contentsLines []string
var contentsLinesCount int

var lineIdx int
var reading = true

var lastLineIdx = -1

func Viewport(content string) {
	cursorHide()
	terminalRows()
	terminalCols()
	reading = true
	lineIdx = 0
	lastLineIdx = -1
	contents = wordwrap.String(content, cols-2)
	contentsLines = strings.Split(contents, eol())
	contentsLinesCount = len(contentsLines)
	drawOutline()
	lineFeed()
	linePrev(2)
	go readKeys(handleViewportKeys)
	for {
		if !reading {
			break
		}
		if lineIdx != lastLineIdx {
			drawBox()
		}
	}
	lineFeed()
	lineFeed()
	cursorShow()
}

func drawOutline() {
	fmt.Print("\u250C")
	for i := 0; i < cols; i++ {
		fmt.Print("\u2500")
	}
	fmt.Print("\u2510")
	for i := 0; i < rows; i++ {
		fmt.Print(eol())
	}
	fmt.Print("\u2514")
	for i := 0; i < cols; i++ {
		fmt.Print("\u2500")
	}
	fmt.Print("\u2518")
}

func drawBox() {
	linePrev(rows - 2)
	var matchingRows = rows + lineIdx - 1
	for i := lineIdx; i < matchingRows; i++ {
		if i < contentsLinesCount {
			clearLine()
			fmt.Print("\u2502 ")
			fmt.Print(contentsLines[i])
			for c := 0; c < (cols - len(contentsLines[i]) - 1); c++ {
				fmt.Print(" ")
			}
			fmt.Print("\u2502 ")
			lineNext(1)
		}
	}
	if matchingRows < rows {
		for i := 0; i < ((rows - 1) - contentsLinesCount); i++ {
			fmt.Print("\u2502")
			for c := 0; c < cols; c++ {
				fmt.Print(" ")
			}
			fmt.Print("\u2502")
			lineNext(1)
		}
	}
	lineNext(2)
	clearLine()
	fmt.Print("^C - Quit\t" + upArrow + " - Up\t" + downArrow + " - Down")
	linePrev(1)
	if lineIdx < contentsLinesCount {
		linePrev(1)
	}
	lastLineIdx = lineIdx
}

func handleViewportKeys(key any) {
	switch key {
	case keyboard.KeyCtrlC:
		reading = false
	case keyboard.KeyArrowUp:
		if (lineIdx - 1) >= 0 {
			lineIdx--
		}
	case keyboard.KeyArrowDown:
		if (rows + lineIdx - 1) < contentsLinesCount {
			lineIdx++
		}
	}
}

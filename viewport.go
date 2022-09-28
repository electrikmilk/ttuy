/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"strings"

	"github.com/eiannone/keyboard"
	"github.com/muesli/reflow/wordwrap"
)

var lineIdx int
var lastLineIdx = -1

var stopViewport = make(chan bool)

var topBox string
var bottomBox string

var contents string
var contentsLines []string
var contentsLinesCount int

var lastViewport string

func Viewport(content string) {
	lineIdx = 0
	lastLineIdx = -1
	contents = wordwrap.String(content, cols-2)
	contentsLines = strings.Split(contents, eol())
	contentsLinesCount = len(contentsLines)
	terminalRows()
	terminalCols()
	cursorHide()
	drawBox()
	go readKeys(handleViewportKeys)
	painter(func() (template string) {
		if lineIdx != lastLineIdx {
			template += topBox + "\n"
			var matchingRows = rows + lineIdx - 1
			for i := lineIdx; i < matchingRows; i++ {
				if i < contentsLinesCount {
					template += "\u2502 "
					template += contentsLines[i]
					for c := 0; c < (cols - len(contentsLines[i]) - 1); c++ {
						template += " "
					}
					template += "\u2502\n"
				}
			}
			if matchingRows < rows {
				for i := 0; i < ((rows - 1) - contentsLinesCount); i++ {
					template += "\u2502"
					for c := 0; c < cols; c++ {
						template += " "
					}
					template += "\u2502\n"
				}
			}
			template += bottomBox + "\n^C Exit - " + upArrow + " Up - " + downArrow + " Down"
			lastViewport = template
		} else {
			template = lastViewport
		}
		return
	})
}

func drawBox() {
	topBox = "\u250C"
	for i := 0; i < cols; i++ {
		topBox += "\u2500"
	}
	topBox += "\u2510"
	bottomBox = "\u2514"
	for i := 0; i < cols; i++ {
		bottomBox += "\u2500"
	}
	bottomBox += "\u2518"
}

func handleViewportKeys(key any) {
	select {
	case <-stopViewport:
		return
	default:
		switch key {
		case keyboard.KeyCtrlC:
			stopPainting()
			stopViewport <- true
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
}

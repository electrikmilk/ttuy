/*
 * Copyright (c) 2023 Brandon Jordan
 */

package ttuy

import (
	"strings"

	"github.com/eiannone/keyboard"
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

// Viewport displays content in a scrollable widget
func Viewport(content string) {
	terminalRows()
	terminalCols()
	lineIdx = 0
	lastLineIdx = -1
	contents = wrapString(&content, cols)
	contentsLines = strings.Split(contents, eol)
	contentsLinesCount = len(contentsLines)
	go readKeys(handleViewportKeys)
	Painter(func() (template string) {
		if lineIdx != lastLineIdx {
			var matchingRows = rows + lineIdx - 1
			for i := lineIdx; i < matchingRows; i++ {
				if i < contentsLinesCount {
					template += contentsLines[i]
					for c := 0; c < (cols - len(contentsLines[i]) - 1); c++ {
						template += " "
					}
					template += eol
				}
			}
			if matchingRows < rows {
				for i := 0; i < ((rows - 1) - contentsLinesCount); i++ {
					for c := 0; c < cols; c++ {
						template += " "
					}
					template += eol
				}
			}
			template += eol + Style("^C Exit \t "+upArrow+" "+downArrow+" Scroll", Dim)
			lastViewport = template
		} else {
			template = lastViewport
		}
		return
	})
}

func wrapString(str *string, limit int) (wrapped string) {
	chars := strings.Split(*str, "")
	i := 0
	for _, char := range chars {
		if char == eol || i == limit {
			if i == limit {
				wrapped += eol + char
				i = 1
			} else {
				wrapped += eol
				i = 0
			}
			continue
		}
		wrapped += char
		i++
	}
	return
}

func moveUp(n int) {
	if (lineIdx - 1) >= 0 {
		lineIdx -= n
	} else {
		Bell()
	}
}

func moveDown(n int) {
	if (rows + lineIdx - 1) < contentsLinesCount {
		lineIdx += n
	} else {
		Bell()
	}
}

func handleViewportKeys(key any) {
	select {
	case <-stopViewport:
		return
	default:
		switch key {
		case keyboard.KeyCtrlC:
			StopPainting()
			stopViewport <- true
		case keyboard.KeyPgup:
			moveUp(5)
		case keyboard.KeyPgdn:
			moveDown(5)
		case keyboard.KeyArrowUp:
			moveUp(1)
		case keyboard.KeyArrowDown:
			moveDown(1)
		}
	}
}

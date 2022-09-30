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
	go readKeys(handleViewportKeys)
	painter(func() (template string) {
		if lineIdx != lastLineIdx {
			var matchingRows = rows + lineIdx - 1
			for i := lineIdx; i < matchingRows; i++ {
				if i < contentsLinesCount {
					template += contentsLines[i]
					for c := 0; c < (cols - len(contentsLines[i]) - 1); c++ {
						template += " "
					}
					template += eol()
				}
			}
			if matchingRows < rows {
				for i := 0; i < ((rows - 1) - contentsLinesCount); i++ {
					for c := 0; c < cols; c++ {
						template += " "
					}
					template += eol()
				}
			}
			template += eol() + Style("^C Exit \t "+upArrow+" "+downArrow+" Scroll", Dim)
			lastViewport = template
		} else {
			template = lastViewport
		}
		return
	})
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

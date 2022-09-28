/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"runtime"

	"github.com/olekukonko/ts"
)

var rows int
var cols int

var upArrow = "\u2191"
var downArrow = "\u2193"

// var leftArrow = "\u2190"
// var rightArrow = "\u2192"
var bullet = "\u2022"

func eol() (eol string) {
	eol = "\n"
	if runtime.GOOS == "windows" {
		eol = "\r\n"
	}
	return
}

func terminalCols() {
	size, _ := ts.GetSize()
	cols = size.Col() - 5
}

func terminalRows() {
	size, _ := ts.GetSize()
	rows = size.Row() - 3
}

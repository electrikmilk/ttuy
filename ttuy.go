/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"github.com/olekukonko/ts"
)

var rows int
var cols int

var upArrow = "\u2191"
var downArrow = "\u2193"
var leftArrow = "\u2190"
var rightArrow = "\u2192"
var bullet = "\u2022"

var dimensions map[int]int

type Row struct {
	Columns []string
}

func eols(times int) (eols string) {
	for i := 0; i < times; i++ {
		eols += eol
	}
	return
}

func terminalCols() {
	size, _ := ts.GetSize()
	cols = size.Col()
}

func terminalRows() {
	size, _ := ts.GetSize()
	rows = size.Row()
}

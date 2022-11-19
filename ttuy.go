/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"github.com/olekukonko/ts"
)

const CSI = "\033["

var rows int
var cols int

const (
	upArrow    = "\u2191"
	downArrow  = "\u2193"
	leftArrow  = "\u2190"
	rightArrow = "\u2192"
	bullet     = "\u2022"
	upper      = "\u203E"
	lower      = "_"
	line       = "|"
)

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

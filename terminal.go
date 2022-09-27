/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"github.com/olekukonko/ts"
)

var rows int
var cols int

func terminalCols() {
	size, _ := ts.GetSize()
	cols = size.Col() - 5
}

func terminalRows() {
	size, _ := ts.GetSize()
	rows = size.Row() - 3
}

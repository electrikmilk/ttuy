/*
 * Copyright (c) 2023 Brandon Jordan
 */

package ttuy

import (
	"fmt"
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

func terminalCols() {
	size, _ := ts.GetSize()
	cols = size.Col()
}

func terminalRows() {
	size, _ := ts.GetSize()
	rows = size.Row()
}

// Bell triggers the system bell which makes the machine make an alarm noise.
func Bell() {
	fmt.Print("\a")
}

func checkEven(rows *[]Row) {
	var evenColumns = len(combined[0].Columns)
	for i, row := range *rows {
		var rowColumns = len(row.Columns)
		if rowColumns != evenColumns {
			panic(fmt.Errorf("uneven columns! Row %d has %d columns when the rest have %d", i, rowColumns, evenColumns))
		}
	}
}

func calcDimensions(rows *[]Row) {
	dimensions = make(map[int]int)
	for _, row := range *rows {
		for c, column := range row.Columns {
			if length, ok := dimensions[c]; ok {
				if len(column) > length {
					dimensions[c] = len(column)
				}
			} else {
				dimensions[c] = len(column)
			}
		}
	}
}

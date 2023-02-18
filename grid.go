/*
 * Copyright (c) 2023 Brandon Jordan
 */

package ttuy

import (
	"fmt"
	"strings"
)

// Grid arranges cells in rows in a grid
func Grid(rows []Row) (grid string) {
	terminalCols()
	calcDimensions(&rows)
	checkEven(&rows)
	var colLength = cols / len(rows[0].Columns)
	for _, row := range rows {
		for _, col := range row.Columns {
			var colChars = strings.Split(col, "")
			if len(col) > colLength {
				for i := 0; i < (colLength - 3); i++ {
					grid += colChars[i]
				}
				grid += "..."
			} else {
				fmt.Print(col)
				for i := 0; i < (colLength - len(colChars)); i++ {
					grid += " "
				}
			}
		}
		grid += eol
	}
	return
}

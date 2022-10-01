/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"fmt"
	"strings"
)

// Grid evenly cells in rows in a grid
func Grid(rows []Row) {
	terminalCols()
	calcDimensions(&rows)
	checkEven(&rows)
	var colLength = cols / len(rows[0].columns)
	for _, row := range rows {
		for _, col := range row.columns {
			var colChars = strings.Split(col, "")
			if len(col) > colLength {
				for i := 0; i < (colLength - 3); i++ {
					fmt.Print(colChars[i])
				}
				fmt.Print("...")
			} else {
				fmt.Print(col)
				for i := 0; i < (colLength - len(colChars)); i++ {
					fmt.Print(" ")
				}
			}
		}
		fmt.Print(eol())
	}
	fmt.Print(eol())
}

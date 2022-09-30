/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"fmt"
)

var combined []Row

// Table prints rows with headers in a text table
func Table(headers []string, rows []Row) {
	combineHeadersRows(&headers, &rows)
	checkEven(&combined)
	calcDimensions(&combined)
	var line = makeLine()
	fmt.Println(line)
	fmt.Print(Style("|", Dim))
	for h, header := range headers {
		fmt.Printf(" %s", header)
		if len(header) < dimensions[h] {
			var spaceDiff = dimensions[h] - len(header)
			for i := 0; i < spaceDiff; i++ {
				fmt.Print(" ")
			}
		}
		fmt.Print(Style(" |", Dim))
	}
	fmt.Print(eol())
	fmt.Println(line)
	for _, row := range rows {
		fmt.Print(Style("|", Dim))
		for c, column := range row.columns {
			fmt.Printf(" %s", column)
			if len(column) < dimensions[c] {
				var spaceDiff = dimensions[c] - len(column)
				for i := 0; i < spaceDiff; i++ {
					fmt.Print(" ")
				}
			}
			fmt.Print(Style(" |", Dim))
		}
		fmt.Print(eol())
	}
	fmt.Println(line)
}

func combineHeadersRows(headers *[]string, rows *[]Row) {
	var headersRow = Row{columns: *headers}
	combined = append(combined, headersRow)
	combined = append(combined, *rows...)
}

func makeLine() (line string) {
	line = "+"
	for c := range combined[0].columns {
		var columnLength = dimensions[c] + 2
		for idx := 0; idx < columnLength; idx++ {
			line += "-"
		}
		line += "+"
	}
	line = Style(line, Dim)
	return
}

func checkEven(rows *[]Row) {
	for _, row := range *rows {
		for _, otherRow := range combined {
			if len(row.columns) != len(otherRow.columns) {
				panic("Uneven table! Number of columns inconsistent!")
			}
		}
	}
}

func calcDimensions(rows *[]Row) {
	dimensions = make(map[int]int)
	for _, row := range *rows {
		for c, column := range row.columns {
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

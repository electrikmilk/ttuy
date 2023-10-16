/*
 * Copyright (c) 2023 Brandon Jordan
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
	fmt.Print(eol)
	fmt.Println(line)
	for _, row := range rows {
		fmt.Print(Style("|", Dim))
		for c, column := range row.Columns {
			fmt.Printf(" %s", column)
			if len(column) < dimensions[c] {
				var spaceDiff = dimensions[c] - len(column)
				for i := 0; i < spaceDiff; i++ {
					fmt.Print(" ")
				}
			}
			fmt.Print(Style(" |", Dim))
		}
		fmt.Print(eol)
	}
	fmt.Println(line)
}

func combineHeadersRows(headers *[]string, rows *[]Row) {
	var headersRow = Row{Columns: *headers}
	combined = append(combined, headersRow)
	combined = append(combined, *rows...)
}

func makeLine() (line string) {
	line = "+"
	for c := range combined[0].Columns {
		var columnLength = dimensions[c] + 2
		for idx := 0; idx < columnLength; idx++ {
			line += "-"
		}
		line += "+"
	}
	line = Style(line, Dim)
	return
}

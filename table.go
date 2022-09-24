/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"fmt"
)

var dimensions map[int]int

type row struct {
	columns []string
}

// Table prints rows with headers in a text table
func Table(headers []string, rows []row) {
	checkEven(&headers, &rows)
	calcDimensions(&headers, &rows)
	var line string = makeLine(&headers)
	fmt.Println(line)
	fmt.Printf("|")
	for h, header := range headers {
		fmt.Printf(" %s", header)
		if len(header) < dimensions[h] {
			var spaceDiff = dimensions[h] - len(header)
			for i := 0; i < spaceDiff; i++ {
				fmt.Printf(" ")
			}
		}
		fmt.Printf(" |")
	}
	fmt.Printf("\n")
	fmt.Println(line)
	for _, row := range rows {
		fmt.Printf("|")
		for c, column := range row.columns {
			fmt.Printf(" %s", column)
			if len(column) < dimensions[c] {
				var spaceDiff = dimensions[c] - len(column)
				for i := 0; i < spaceDiff; i++ {
					fmt.Printf(" ")
				}
			}
			fmt.Printf(" |")
		}
		fmt.Printf("\n")
	}
	fmt.Println(line)
}

func makeLine(headers *[]string) (line string) {
	line = "+"
	for h := range *headers {
		var columnLength int = dimensions[h] + 2
		for idx := 0; idx < columnLength; idx++ {
			line += "-"
		}
		line += "+"
	}
	return
}

func checkEven(headers *[]string, rows *[]row) {
	var headersCount int = len(*headers)
	for _, row := range *rows {
		var rowCount int = len(row.columns)
		if rowCount < headersCount || rowCount > headersCount {
			fmt.Println(headers, rows)
			panic("Uneven table! Number of columns in rows and header columns do not match!")
		}
	}
}

func calcDimensions(headers *[]string, rows *[]row) {
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
	for h, header := range *headers {
		if length, ok := dimensions[h]; ok {
			if len(header) > length {
				dimensions[h] = len(header)
			}
		} else {
			dimensions[h] = len(header)
		}
	}
}
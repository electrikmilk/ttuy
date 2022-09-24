/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"testing"
)

func TestTable(t *testing.T) {
	var headers = []string{"Header 1", "Header 2", "Header 3", "Header 4"}
	var rows []row
	for i := 0; i < 5; i++ {
		var rowRows = []string{"Cell 1", "Cell 2", "Cell 3", "Cell 4"}
		rows = append(rows, row{columns: rowRows})
	}
	Table(headers, rows)
}

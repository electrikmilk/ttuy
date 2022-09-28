/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"fmt"
)

// ProgressBar prints a progress bar
func ProgressBar(status int) {
	cursorHide()
	bar(status)
	if status == 100 {
		lineFeed()
		lineFeed()
		cursorShow()
	}
}

func bar(status int) (bar string) {
	var progress = float64(status) / float64(100)
	terminalCols()
	cols -= 3
	progressBlocks := int(progress * float64(cols))
	remain := cols - progressBlocks
	if status == 0 {
		lineFeed()
		linePrev(1)
		fmt.Print("\u250C")
		for i := 0; i < (progressBlocks + remain); i++ {
			fmt.Print("\u2500")
		}
		fmt.Print("\u2510")
	}
	lineNext(1)
	// middle
	clearLine()
	fmt.Print("\u2502")
	for i := 0; i < progressBlocks; i++ {
		fmt.Print(Style("\u2590", Cyan))
	}
	for i := 0; i < remain; i++ {
		fmt.Print(" ")
	}
	fmt.Print("\u2502 " + Style(fmt.Sprintf("%d%%", status), Bold, Green))
	// bottom
	if status == 0 {
		fmt.Print(eol())
		fmt.Print("\u2514")
		for i := 0; i < (progressBlocks + remain); i++ {
			fmt.Print("\u2500")
		}
		fmt.Print("\u2518")
		linePrev(1)
	}
	linePrev(1)
	return
}

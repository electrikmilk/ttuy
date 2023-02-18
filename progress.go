/*
 * Copyright (c) 2023 Brandon Jordan
 */

package ttuy

import (
	"fmt"
)

// ProgressBar prints a progress bar
func ProgressBar(status int) {
	CursorHide()
	bar(status)
	if status == 100 {
		LineFeed()
		LineFeed()
		CursorShow()
	}
}

func bar(status int) (bar string) {
	var progress = float64(status) / float64(100)
	terminalCols()
	cols -= 8
	progressBlocks := int(progress * float64(cols))
	remain := cols - progressBlocks
	if status == 0 {
		LineFeed()
		LinePrev(1)
		fmt.Print("\u250C")
		for i := 0; i < (progressBlocks + remain); i++ {
			fmt.Print("\u2500")
		}
		fmt.Print("\u2510")
	}
	LineNext(1)
	// middle
	ClearLine()
	fmt.Print("\u2502")
	for i := 0; i < progressBlocks; i++ {
		fmt.Print(Style("\u2590", CyanText))
	}
	for i := 0; i < remain; i++ {
		fmt.Print(" ")
	}
	fmt.Print("\u2502 " + Style(fmt.Sprintf("%d%%", status), Bold, GreenText))
	// bottom
	if status == 0 {
		fmt.Print(eol)
		fmt.Print("\u2514")
		for i := 0; i < (progressBlocks + remain); i++ {
			fmt.Print("\u2500")
		}
		fmt.Print("\u2518")
		LinePrev(1)
	}
	LinePrev(1)
	return
}

/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// ProgressBar prints a progress bar
func ProgressBar(status int) {
	cursorHide()
	bar(status)
	if status == 100 {
		lineNext(2)
		fmt.Print(eol())
		cursorShow()
	}
}

func bar(status int) (bar string) {
	var progress = float64(status) / float64(100)
	out, err := exec.Command("tput", "cols").Output()
	if err != nil {
		panic(err)
	}
	cols, _ := strconv.Atoi(strings.TrimSuffix(string(out), eol()))
	progressBlocks := int(progress * float64(cols))
	remain := cols - progressBlocks
	if status == 0 {
		for i := 0; i < 3; i++ {
			lineFeed()
		}
		linePrev(3)
		fmt.Print("\u250C")
		for i := 0; i < (progressBlocks + remain); i++ {
			fmt.Print("\u2500")
		}
		fmt.Print("\u2510")
	}
	lineNext(1)
	// middle
	fmt.Print("\u2502")
	for i := 0; i < progressBlocks; i++ {
		fmt.Print(Style("█", Cyan))
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

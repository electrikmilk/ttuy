/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import "fmt"

const CSI = "\033["

/*
A = UP
B = DOWN
C = FORWARD
D = BACKWARD
E = NEXT
F = PREV
G = BEGIN
r;c H = CUSTOM
K = ERASE IN LINE
J = ERASE IN DISPLAY
S = SCROLL UP
T = SCROLL DOWN
?25l = HIDE
?25h = SHOW
ES = LF
*/

var currentLine int

func backUp() {
	linePrev(currentLine)
	currentLine = 0
}

func lineUp(lines int) {
	fmt.Printf(CSI+"%dA", lines)
}

func lineDown(lines int) {
	fmt.Printf(CSI+"%dB", lines)
}

func lineNext(lines int) {
	fmt.Printf(CSI+"%dE", lines)
	currentLine++
}

func linePrev(lines int) {
	fmt.Printf(CSI+"%dF", lines)
}

func lineFeed() {
	lineNext(1)
	scrollUp(1)
}

func cursor(row int, col int) {
	fmt.Printf(CSI+"%d;%dH", row, col)
}

func cursorFwd(columns int) {
	fmt.Printf(CSI+"%dC", columns)
}

func cursorBack(columns int) {
	fmt.Printf(CSI+"%dD", columns)
}

func cursorStart() {
	fmt.Print(CSI + "1G")
}

func cursorHide() {
	fmt.Printf(CSI + "?25l")
}

func cursorShow() {
	fmt.Printf(CSI + "?25h")
}

func clearLine() {
	fmt.Printf(CSI + "0K")
}

func clearDisplay() {
	fmt.Printf(CSI + "0J")
}

func scrollUp(lines int) {
	fmt.Printf(CSI+"%dS", lines)
}

func scrollDown(lines int) {
	fmt.Printf(CSI+"%dT", lines)
}

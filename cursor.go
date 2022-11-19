/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import "fmt"

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

func BackUp() {
	LinePrev(currentLine)
	currentLine = 0
}

func LineUp(lines int) {
	fmt.Printf(CSI+"%dA", lines)
}

func LineDown(lines int) {
	fmt.Printf(CSI+"%dB", lines)
}

func LineNext(lines int) {
	fmt.Printf(CSI+"%dE", lines)
	currentLine++
}

func LinePrev(lines int) {
	fmt.Printf(CSI+"%dF", lines)
}

func lineFeed() {
	LineNext(1)
	ScrollUp(1)
}

func Cursor(row int, col int) {
	fmt.Printf(CSI+"%d;%dH", row, col)
}

func CursorFwd(columns int) {
	fmt.Printf(CSI+"%dC", columns)
}

func CursorBack(columns int) {
	fmt.Printf(CSI+"%dD", columns)
}

func CursorStart() {
	fmt.Print(CSI + "1G")
}

func CursorHide() {
	fmt.Printf(CSI + "?25l")
}

func CursorShow() {
	fmt.Printf(CSI + "?25h")
}

func ClearLine() {
	fmt.Printf(CSI + "0K")
}

func ClearDisplay() {
	fmt.Printf(CSI + "0J")
}

func ScrollUp(lines int) {
	fmt.Printf(CSI+"%dS", lines)
}

func ScrollDown(lines int) {
	fmt.Printf(CSI+"%dT", lines)
}

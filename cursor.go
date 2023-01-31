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

// BackUp passes `currentLine` to LinePrev to go to the previous line `currentLine` number of times.
// It then resets `currentLine` to 0.
func BackUp() {
	LinePrev(currentLine)
	currentLine = 0
}

// LineUp moves the cursor up `lines` number of lines
func LineUp(lines int) {
	fmt.Printf(CSI+"%dA", lines)
}

// LineDown moves the cursor down `lines` number of lines
func LineDown(lines int) {
	fmt.Printf(CSI+"%dB", lines)
}

// LineNext moves the cursor to the next line, `lines` number of times
func LineNext(lines int) {
	fmt.Printf(CSI+"%dE", lines)
	currentLine++
}

// LinePrev moves the cursor to the previous line, `lines` number of times
func LinePrev(lines int) {
	fmt.Printf(CSI+"%dF", lines)
}

// LineFeed simulates a line feed by going to the next line and scrolling up once
func LineFeed() {
	LineNext(1)
	ScrollUp(1)
}

// Cursor moves the cursor to a specific position on the screen.
// `row` sets the cursors Y position
// `col` sets the cursors X position
func Cursor(row int, col int) {
	fmt.Printf(CSI+"%d;%dH", row, col)
}

// CursorFwd moves the cursor forward `columns` number of times
func CursorFwd(columns int) {
	fmt.Printf(CSI+"%dC", columns)
}

// CursorBack moves the cursor back `columns` number of times
func CursorBack(columns int) {
	fmt.Printf(CSI+"%dD", columns)
}

// CursorStart moves the cursor to the start of the current line
func CursorStart() {
	fmt.Print(CSI + "1G")
}

// CursorHide makes the cursor hidden
func CursorHide() {
	fmt.Printf(CSI + "?25l")
}

// CursorShow makes the cursor visible
func CursorShow() {
	fmt.Printf(CSI + "?25h")
}

// ClearLine clears the current line
func ClearLine() {
	fmt.Printf(CSI + "0K")
}

// ClearDisplay clears part of the screen
func ClearDisplay() {
	fmt.Printf(CSI + "2J")
}

// ScrollUp scrolls the whole page up by `lines` lines
func ScrollUp(lines int) {
	fmt.Printf(CSI+"%dS", lines)
}

// ScrollDown scrolls the whole page down by `lines` lines
func ScrollDown(lines int) {
	fmt.Printf(CSI+"%dT", lines)
}

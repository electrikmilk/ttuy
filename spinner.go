/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"fmt"
	"time"
)

type SpinnerStyle string

const Ticker SpinnerStyle = "ticker"
const DotDotDot SpinnerStyle = "dotdotdot"
const Throbber SpinnerStyle = "throbber"
const Blinker SpinnerStyle = "blinker"

var ticks = []string{"|", "/", "—", "\\"}
var dots = []string{"", ".", "..", "..."}
var throbs = []string{
	Style(bullet, CyanText) + Style(bullet, Dim) + Style(bullet, Dim),
	Style(bullet, Dim) + Style(bullet, CyanText) + Style(bullet, Dim),
	Style(bullet, Dim) + Style(bullet, Dim) + Style(bullet, CyanText),
}
var blinks = []string{
	Style(bullet, Dim),
	Style(bullet, CyanText, Blink),
}

var stop = make(chan bool)

// Spinner prints a progress indicator in style until StopSpinner() is called
// You must use a goroutine when running this function (e.g. go Spinner(...))
func Spinner(status string, style SpinnerStyle) {
	CursorHide()
	fmt.Print(eol)
	LinePrev(1)
	for {
		select {
		case <-stop:
			return
		default:
			switch style {
			case Ticker:
				for i := 0; i < len(ticks); i++ {
					ClearLine()
					fmt.Print(Style(ticks[i], CyanText))
					fmt.Print(" " + Style(status, Bold))
					fmt.Print("\r")
					time.Sleep(100 * time.Millisecond)
				}
			case DotDotDot:
				for i := 0; i < len(dots); i++ {
					ClearLine()
					fmt.Print(Style(status, Bold))
					fmt.Print(Style(dots[i], Dim))
					fmt.Print("\r")
					time.Sleep(300 * time.Millisecond)
				}
			case Throbber:
				for i := 0; i < len(throbs); i++ {
					ClearLine()
					fmt.Print(throbs[i])
					fmt.Print(" " + Style(status, Bold))
					fmt.Print("\r")
					time.Sleep(200 * time.Millisecond)
				}
			case Blinker:
				for i := 0; i < len(blinks); i++ {
					ClearLine()
					fmt.Print(blinks[i])
					fmt.Print(" " + Style(status, Bold))
					fmt.Print("\r")
					time.Sleep(500 * time.Millisecond)
				}
			}
		}
	}
}

// StopSpinner stops the current spinner
func StopSpinner() {
	stop <- true
	ClearLine()
	fmt.Print(eol)
	CursorShow()
}

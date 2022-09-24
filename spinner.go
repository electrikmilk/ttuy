/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"fmt"
	"time"
)

type SpinnerStyle string

const (
	Ticker    SpinnerStyle = "ticker"
	DotDotDot              = "dotdotdot"
	Throbber               = "throbber"
	Blinker                = "blinker"
)

var ticks = []string{"|", "/", "—", "\\"}
var dots = []string{"", ".", "..", "..."}
var throbs = []string{
	Style("•", Cyan) + Style("•", Dim) + Style("•", Dim),
	Style("•", Dim) + Style("•", Cyan) + Style("•", Dim),
	Style("•", Dim) + Style("•", Dim) + Style("•", Cyan),
}
var blinks = []string{
	Style("•", Dim),
	Style("•", Cyan, BlinkFast),
}

var stop = make(chan bool)

func Spinner(status string, style SpinnerStyle) {
	cursorHide()
	fmt.Print(eol())
	linePrev(1)
	for {
		select {
		case <-stop:
			return
		default:
			switch style {
			case Ticker:
				for i := 0; i < len(ticks); i++ {
					clearLine()
					fmt.Print(Style(ticks[i], Cyan))
					fmt.Print(Style(fmt.Sprintf(" %s", status), Bold))
					fmt.Print("\r")
					time.Sleep(100 * time.Millisecond)
				}
			case DotDotDot:
				for i := 0; i < len(dots); i++ {
					clearLine()
					fmt.Print(Style(fmt.Sprintf("%s", status), Bold))
					fmt.Print(Style(dots[i], Dim))
					fmt.Print("\r")
					time.Sleep(300 * time.Millisecond)
				}
			case Throbber:
				for i := 0; i < len(throbs); i++ {
					clearLine()
					fmt.Print(throbs[i])
					fmt.Print(Style(fmt.Sprintf(" %s", status), Bold))
					fmt.Print("\r")
					time.Sleep(200 * time.Millisecond)
				}
			case Blinker:
				for i := 0; i < len(blinks); i++ {
					clearLine()
					fmt.Print(blinks[i])
					fmt.Print(Style(fmt.Sprintf(" %s", status), Bold))
					fmt.Print("\r")
					time.Sleep(500 * time.Millisecond)
				}
			}
		}
	}
}

func StopSpinner() {
	stop <- true
	clearLine()
	fmt.Print(eol())
	cursorShow()
}

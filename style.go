/*
 * Copyright (c) 2023 Brandon Jordan
 */

package ttuy

import "fmt"

// Style formats str using styles to prefix str with SGR sequences
func Style(str string, styles ...SGR) string {
	return style(str, true, styles...)
}

// StylePersist outputs SGR sequences with no reset
// Anything after will be formatted using styles until the Reset constant is used
func StylePersist(styles ...SGR) string {
	return style("", false, styles...)
}

func style(str string, reset bool, styles ...SGR) (styled string) {
	for _, s := range styles {
		styled += fmt.Sprintf(CSI+"%sm", s)
	}
	styled += str
	if reset {
		styled += CSI + "0m"
	}
	return
}

// Foreground style str using code
func Foreground(str string, code int) string {
	return customStyle(str, fg, code)
}

// Background style str using code
func Background(str string, code int) string {
	return customStyle(str, bg, code)
}

func customStyle(str string, plane SGR, code int) string {
	return fmt.Sprintf(CSI+"%s%dm%s", plane, code, str) + CSI + "0m"
}

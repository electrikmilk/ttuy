/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import "fmt"

// Style formats str using styles to prefix str with SGR sequences
func Style(str string, styles ...SGR) (styled string) {
	for _, s := range styles {
		styled += fmt.Sprintf(CSI+"%sm", s)
	}
	styled += str + CSI + "0m"
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

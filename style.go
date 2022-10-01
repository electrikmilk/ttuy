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

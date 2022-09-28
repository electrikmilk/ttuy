/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import "runtime"

func eol() (eol string) {
	eol = "\n"
	if runtime.GOOS == "windows" {
		eol = "\r\n"
	}
	return
}

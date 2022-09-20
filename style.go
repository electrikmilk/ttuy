/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import "fmt"

type SGR int

// Decorations
const (
	BOLD      SGR = 1
	DIM           = 2
	ITALIC        = 3
	UNDERLINE     = 4
	SLOWBLINK     = 5
	FASTBLINK     = 6
	INVERT        = 7
	HIDE          = 8
	CROSS         = 9
	STOPBLINK     = 25
	SHOW          = 28
	STOPCROSS     = 29
)

// Foreground Colors
const (
	BLACK          SGR = 30
	RED                = 31
	GREEN              = 32
	YELLOW             = 33
	BLUE               = 34
	MAGENTA            = 35
	CYAN               = 36
	WHITE              = 37
	BRIGHT_BLACK       = 90
	BRIGHT_RED         = 91
	BRIGHT_GREEN       = 92
	BRIGHT_YELLOW      = 93
	BRIGHT_BLUE        = 94
	BRIGHT_MAGENTA     = 95
	BRIGHT_CYAN        = 96
	BRIGHT_WHITE       = 97
)

// Background Colors
const (
	BG_BLACK          SGR = 40
	BG_RED                = 41
	BG_GREEN              = 42
	BG_YELLOW             = 43
	BG_BLUE               = 44
	BG_MAGENTA            = 45
	BG_CYAN               = 46
	BG_WHITE              = 47
	BG_BRIGHT_BLACK       = 100
	BG_BRIGHT_RED         = 101
	BG_BRIGHT_GREEN       = 102
	BG_BRIGHT_YELLOW      = 103
	BG_BRIGHT_BLUE        = 104
	BG_BRIGHT_MAGENTA     = 105
	BG_BRIGHT_CYAN        = 106
	BG_BRIGHT_WHITE       = 107
)

func Style(str string, styles ...SGR) (styled string) {
	for _, s := range styles {
		styled += fmt.Sprintf("\033[%dm", s)
	}
	styled += fmt.Sprintf("%s\033[0m", str)
	return
}

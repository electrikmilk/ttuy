/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import "fmt"

type SGR int

// Decorations

const Bold SGR = 1
const Dim SGR = 2
const Italic SGR = 3
const Underlined SGR = 4
const Blink SGR = 5
const BlinkFast SGR = 6
const Inverted SGR = 7
const Hidden SGR = 8
const Crossed SGR = 9

// Foreground Colors

const Black SGR = 30
const Red SGR = 31
const Green SGR = 32
const Yellow SGR = 33
const Blue SGR = 34
const Magenta SGR = 35
const Cyan SGR = 36
const White SGR = 37
const LightBlack SGR = 90
const LightRed SGR = 91
const LightGreen SGR = 92
const LightYellow SGR = 93
const LightBlue SGR = 94
const LightMagenta SGR = 95
const LightCyan SGR = 96
const LightWhite SGR = 97

// Background Colors

const BlackBg SGR = 40
const RedBg SGR = 41
const GreenBg SGR = 42
const YellowBg SGR = 43
const BlueBg SGR = 44
const MagentaBg SGR = 45
const CyanBg SGR = 46
const WhiteBg SGR = 47
const LightBlackBg SGR = 100
const LightRedBg SGR = 101
const LightGreenBg SGR = 102
const LightYellowBg SGR = 103
const LightBlueBg SGR = 104
const LightMagentaBg SGR = 105
const LightCyanBg SGR = 106
const LightWhiteBg SGR = 107

// Style formats str using styles to prefix str with SGR sequences
func Style(str string, styles ...SGR) (styled string) {
	for _, s := range styles {
		styled += fmt.Sprintf(CSI+"%dm", s)
	}
	styled += str + CSI + "0m"
	return
}

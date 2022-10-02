/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

type SGR string

var FG SGR = "38;5;"
var BG SGR = "48;5;"

// Decorations

const Bold SGR = "1"
const Dim SGR = "2"
const Italic SGR = "3"
const Underlined SGR = "4"
const Blink SGR = "5"
const BlinkFast SGR = "6"
const Inverted SGR = "7"
const Hidden SGR = "8"
const Crossed SGR = "9"

// Basic

const black SGR = "0"
const red SGR = "1"
const green SGR = "2"
const yellow SGR = "3"
const blue SGR = "4"
const magenta SGR = "5"
const cyan SGR = "6"
const gray SGR = "8"

const brightRed SGR = "9"
const brightGreen SGR = "10"
const brightYellow SGR = "11"
const brightBlue SGR = "12"
const brightMagenta SGR = "13"
const brightCyan SGR = "14"
const white SGR = "15"

var BlackText = FG + black
var RedText = FG + red
var GreenText = FG + green
var YellowText = FG + yellow
var BlueText = FG + blue
var MagentaText = FG + magenta
var CyanText = FG + cyan
var WhiteText = FG + white

var GrayText = FG + gray
var BrightRedText = FG + brightRed
var BrightGreenText = FG + brightGreen
var BrightYellowText = FG + brightYellow
var BrightBlueText = FG + brightBlue
var BrightMagentaText = FG + brightMagenta
var BrightCyanText = FG + brightCyan

var BlackBg = BG + black
var RedBg = BG + red
var GreenBg = BG + green
var YellowBg = BG + yellow
var BlueBg = BG + blue
var MagentaBg = BG + magenta
var CyanBg = BG + cyan
var WhiteBg = BG + white

var GrayBg = BG + gray
var BrightRedBg = BG + brightRed
var BrightGreenBg = BG + brightGreen
var BrightYellowBg = BG + brightYellow
var BrightBlueBg = BG + brightBlue
var BrightMagentaBg = BG + brightMagenta
var BrightCyanBg = BG + brightCyan

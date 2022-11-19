/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

type SGR string

var fg SGR = "38;5;"
var bg SGR = "48;5;"

const Reset = CSI + "0m"

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

var BlackText = fg + black
var RedText = fg + red
var GreenText = fg + green
var YellowText = fg + yellow
var BlueText = fg + blue
var MagentaText = fg + magenta
var CyanText = fg + cyan
var WhiteText = fg + white

var GrayText = fg + gray
var BrightRedText = fg + brightRed
var BrightGreenText = fg + brightGreen
var BrightYellowText = fg + brightYellow
var BrightBlueText = fg + brightBlue
var BrightMagentaText = fg + brightMagenta
var BrightCyanText = fg + brightCyan

var BlackBg = bg + black
var RedBg = bg + red
var GreenBg = bg + green
var YellowBg = bg + yellow
var BlueBg = bg + blue
var MagentaBg = bg + magenta
var CyanBg = bg + cyan
var WhiteBg = bg + white

var GrayBg = bg + gray
var BrightRedBg = bg + brightRed
var BrightGreenBg = bg + brightGreen
var BrightYellowBg = bg + brightYellow
var BrightBlueBg = bg + brightBlue
var BrightMagentaBg = bg + brightMagenta
var BrightCyanBg = bg + brightCyan

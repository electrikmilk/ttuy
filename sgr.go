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

var Black = FG + black
var Red = FG + red
var Green = FG + green
var Yellow = FG + yellow
var Blue = FG + blue
var Magenta = FG + magenta
var Cyan = FG + cyan

var BlackBg = BG + black
var RedBg = BG + red
var GreenBg = BG + green
var YellowBg = BG + yellow
var BlueBg = BG + blue
var MagentaBg = BG + magenta
var CyanBg = BG + cyan

var Gray = FG + gray
var BrightRed = FG + brightRed
var BrightGreen = FG + brightGreen
var BrightYellow = FG + brightYellow
var BrightBlue = FG + brightBlue
var BrightMagenta = FG + brightMagenta
var BrightCyan = FG + brightCyan

var GrayBg = BG + gray
var BrightRedBg = BG + brightRed
var BrightGreenBg = BG + brightGreen
var BrightYellowBg = BG + brightYellow
var BrightBlueBg = BG + brightBlue
var BrightMagentaBg = BG + brightMagenta
var BrightCyanBg = BG + brightCyan

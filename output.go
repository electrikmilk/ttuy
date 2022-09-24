/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"fmt"
	"os"
)

// Fail prints message styled then exits
func Fail(message string) {
	var label = Style(" \u2613 Error ", LightRedBg, Black, Bold)
	message = fmt.Sprintf(" %s", message)
	message = Style(message, LightRed, Bold)
	fmt.Println(label + message)
	os.Exit(1)
}

// Failf prints message formatted and styled then exits
func Failf(message string, v ...any) {
	message = fmt.Sprintf(message, v...)
	Fail(message)
}

// FailErr handles err and prints label styled then exits
func FailErr(label string, err error) {
	if err != nil {
		label = Style(fmt.Sprintf(" %s ", label), RedBg, Black, Bold)
		fmt.Println(label + fmt.Sprintf(" %s", err))
		os.Exit(1)
	}
}

// Warn displays a styled warning message
func Warn(message string) {
	var label = Style(" \u2691 Warning ", YellowBg, Black, Bold)
	message = fmt.Sprintf(" %s", message)
	message = Style(message, Yellow, Bold)
	fmt.Println(label + message)
}

// Warnf prints a formatted styled warning message
func Warnf(message string, v ...any) {
	message = fmt.Sprintf(message, v...)
	Warn(message)
}

// Info prints a styled info message
func Info(message string) {
	message = Style(message, Blue, Bold)
	fmt.Println(message)
}

// Infof prints a formatted styled info message
func Infof(message string, v ...any) {
	message = fmt.Sprintf(message, v...)
	Info(message)
}

// Success prints a styled success message
func Success(message string) {
	var label = Style(" \u2714 ", GreenBg, Black, Bold)
	message = fmt.Sprintf(" %s", message)
	message = Style(message, Green, Bold)
	fmt.Println(label + message)
}

// Successf prints a formatted styled success message
func Successf(message string, v ...any) {
	message = fmt.Sprintf(message, v...)
	Success(message)
}

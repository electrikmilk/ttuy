/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"fmt"
	"os"
)

func Fail(message string) {
	var label = Style(" \u2613 Error ", LightRedBg, Black, Bold)
	message = fmt.Sprintf(" %s", message)
	message = Style(message, LightRed, Bold)
	fmt.Println(label + message)
	os.Exit(1)
}

func Failf(message string, v ...any) {
	message = fmt.Sprintf(message, v...)
	Fail(message)
}

func FailErr(label string, err error) {
	if err != nil {
		label = Style(fmt.Sprintf(" %s ", label), RedBg, Black, Bold)
		fmt.Println(label + fmt.Sprintf(" %s", err))
		os.Exit(1)
	}
}

func Warn(message string) {
	var label = Style(" \u2691 Warning ", YellowBg, Black, Bold)
	message = fmt.Sprintf(" %s", message)
	message = Style(message, Yellow, Bold)
	fmt.Println(label + message)
}

func Warnf(message string, v ...any) {
	message = fmt.Sprintf(message, v...)
	Warn(message)
}

func Info(message string) {
	message = Style(message, Blue, Bold)
	fmt.Println(message)
}

func Infof(message string, v ...any) {
	message = fmt.Sprintf(message, v...)
	Info(message)
}

func Success(message string) {
	var label = Style(" \u2714 ", GreenBg, Black, Bold)
	message = fmt.Sprintf(" %s", message)
	message = Style(message, Green, Bold)
	fmt.Println(label + message)
}

func Successf(message string, v ...any) {
	message = fmt.Sprintf(message, v...)
	Success(message)
}

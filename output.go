/*
 * Copyright (c) 2023 Brandon Jordan
 */

package ttuy

import (
	"fmt"
	"os"
)

func printError(label string, message string) {
	label = Style(fmt.Sprintf(" \u2613 %s ", label), BrightRedBg, BlackText, Bold)
	message = fmt.Sprintf(" %s", message)
	message = Style(message, BrightRedText, Bold)
	fmt.Println(label + message)
	os.Exit(1)
}

// Fail prints message styled as an error then exits
func Fail(message string) {
	printError("Error", message)
}

// Failf prints message formatted and styled as an error then exits
func Failf(message string, v ...any) {
	message = fmt.Sprintf(message, v...)
	printError("Error", message)
}

// FailErr handles err and prints label with err as a string styled then exits
func FailErr(label string, err error) {
	if err != nil {
		var message = fmt.Sprintf("%s", err)
		printError(label, message)
	}
}

// Warn prints message styled as a warning
func Warn(message string) {
	var label = Style(" \u2691 Warning ", YellowBg, BlackText, Bold)
	message = fmt.Sprintf(" %s", message)
	message = Style(message, YellowText, Bold)
	fmt.Println(label + message)
}

// Warnf prints message formatted and styled as a warning
func Warnf(message string, v ...any) {
	message = fmt.Sprintf(message, v...)
	Warn(message)
}

// Info prints message highlighted
func Info(message string) {
	message = Style(message, BlueText, Bold)
	fmt.Println(message)
}

// Infof prints message formatted and highlighted
func Infof(message string, v ...any) {
	message = fmt.Sprintf(message, v...)
	Info(message)
}

// Success prints message styled as a success
func Success(message string) {
	var label = Style(" \u2714 ", GreenBg, BlackText, Bold)
	message = fmt.Sprintf(" %s", message)
	message = Style(message, GreenText, Bold)
	fmt.Println(label + message)
}

// Successf prints message styled and formatted as a success
func Successf(message string, v ...any) {
	message = fmt.Sprintf(message, v...)
	Success(message)
}

// BigBanner returns a box around `message` made of # signs.
// ###########
// # EXAMPLE #
// ###########
func BigBanner(message string) (banner string) {
	terminalRows()
	terminalCols()
	banner = "\n"
	var spaces = (cols - (len(message) + 2)) / 2
	for i := 0; i < cols; i++ {
		banner += "#"
	}

	banner += "\n#"
	for i := 0; i < cols-2; i++ {
		banner += " "
	}
	banner += "#"

	banner += "\n#"
	for i := 0; i < spaces; i++ {
		banner += " "
	}
	banner += message
	for i := 0; i < spaces; i++ {
		banner += " "
	}
	banner += "#\n"

	banner += "#"
	for i := 0; i < cols-2; i++ {
		banner += " "
	}
	banner += "#\n"

	for i := 0; i < cols; i++ {
		banner += "#"
	}
	banner += "\n\n"

	return
}

// Banner returns a banner around `message` made of # signs.
// ######## EXAMPLE ########
func Banner(message string) (banner string) {
	terminalRows()
	terminalCols()
	banner = "\n"
	var spaces = (cols - (len(message) + 2)) / 2
	for i := 0; i < spaces; i++ {
		banner += "#"
	}
	banner += " " + message + " "
	for i := 0; i < spaces; i++ {
		banner += "#"
	}
	banner += "#\n\n"
	return
}

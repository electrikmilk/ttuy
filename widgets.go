/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

func eol() (eol string) {
	eol = "\n"
	if runtime.GOOS == "windows" {
		eol = "\r\n"
	}
	return
}

// File print the content of the file at path
func File(path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(bytes) + eol())
}

// Ask prints your prompt and feeds input from os.Stdin into store
func Ask(prompt string, store *string) {
	prompt = Style(fmt.Sprintf(" %s: ", prompt), Bold, Inverted)
	fmt.Printf("%s ", prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		*store = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Print(eol())
}

// AskSecret prints your prompt and feeds input from os.Stdin into store but hides the input
func AskSecret(prompt string, store *string) {
	prompt = Style(fmt.Sprintf("\U0001F512 %s: ", prompt), Bold, RedBg, Black)
	fmt.Printf("%s "+CSI+"8m", prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		*store = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Print(CSI + "0m" + eol())
}

// YesNo prints a prompt that waits for "y" or "n" from os.Stdin and returns a boolean based on the input
func YesNo(prompt string) (proceed bool) {
	var input string
	for {
		Ask(prompt+" (y/n)", &input)
		input = strings.ToLower(input)
		if input == "y" || input == "n" {
			if input == "y" {
				proceed = true
			} else if input == "n" {
				proceed = false
			}
			break
		}
		input = ""
	}
	return
}

// Typewriter prints characters in message out one at a time
func Typewriter(message string) {
	TypewriterTimed(message, 100)
}

// TypewriterTimed prints characters in message out one at a time at duration
func TypewriterTimed(message string, duration time.Duration) {
	cursorHide()
	var strChars = strings.Split(message, "")
	var chars int
	for i := 0; i <= len(strChars); i++ {
		for _, char := range strChars[:i] {
			fmt.Print(char)
		}
		fmt.Printf("\r")
		chars++
		time.Sleep(duration * time.Millisecond)
	}
	fmt.Print(eol())
	cursorShow()
}

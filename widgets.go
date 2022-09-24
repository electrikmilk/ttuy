/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
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

func File(name string) {
	bytes, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(bytes) + eol())
}

func ProgressBar(status int) {
	cursorHide()
	bar(status)
	if status == 100 {
		lineNext(2)
		fmt.Print(eol())
		cursorShow()
	}
}

func bar(status int) (bar string) {
	var prog = float64(status) / float64(100)
	out, err := exec.Command("tput", "cols").Output()
	if err != nil {
		panic(err)
	}
	cols, _ := strconv.Atoi(strings.TrimSuffix(string(out), eol()))
	progBlocks := int(prog * float64(cols))
	remain := cols - progBlocks
	if status == 0 {
		for i := 0; i < 3; i++ {
			lineFeed()
		}
		linePrev(3)
		fmt.Print("\u250C")
		for i := 0; i < (progBlocks + remain); i++ {
			fmt.Print("\u2500")
		}
		fmt.Print("\u2510")
	}
	lineNext(1)
	// middle
	fmt.Print("\u2502")
	for i := 0; i < progBlocks; i++ {
		fmt.Print(Style("█", Cyan))
	}
	for i := 0; i < remain; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("\u2502 %d%%", status)
	// bottom
	if status == 0 {
		fmt.Print(eol())
		fmt.Print("\u2514")
		for i := 0; i < (progBlocks + remain); i++ {
			fmt.Print("\u2500")
		}
		fmt.Print("\u2518")
		linePrev(1)
	}
	linePrev(1)
	return
}

func wait() {
	time.Sleep(30 * time.Millisecond)
}

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

func Typewriter(message string) {
	TypewriterTimed(message, 100)
}

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

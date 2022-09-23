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

func Progress(status int) {
	cursorHide()
	fmt.Printf(bar(status)+" %d%%\r", status)
	if status == 100 {
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
	bar = "["
	for i := 0; i < progBlocks; i++ {
		bar += "="
	}
	for i := 0; i < remain; i++ {
		bar += " "
	}
	bar += "]"
	return
}

func Ask(prompt string, store *string) {
	prompt = Style(fmt.Sprintf(" %s: ", prompt), Bold, Inverted)
	fmt.Printf("%s ", prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		*store = scanner.Text()
	}
	fmt.Print(eol())
}

func YesNo(prompt string) (proceed bool) {
	var input string
	for {
		Ask(prompt+" (y/n)?", &input)
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

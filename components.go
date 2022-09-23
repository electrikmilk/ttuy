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

func EOL() (EOL string) {
	EOL = "\n"
	if runtime.GOOS == "windows" {
		EOL = "\r\n"
	}
	return
}

func File(name string) {
	bytes, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(bytes[:]) + EOL())
}

func Progress(status int) {
	cursorHide()
	fmt.Printf(bar(status)+" %d%%\r", status)
	if status == 100 {
		fmt.Printf(EOL())
		cursorShow()
	}
}

func bar(status int) (bar string) {
	var prog float64 = float64(status) / float64(100)
	out, err := exec.Command("tput", "cols").Output()
	if err != nil {
		panic(err)
	}
	cols, _ := strconv.Atoi(strings.TrimSuffix(string(out), EOL()))
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
	prompt = Style(fmt.Sprintf(" %s: ", prompt), BOLD, INVERT)
	fmt.Printf("%s ", prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		*store = scanner.Text()
	}
	fmt.Printf(EOL())
}

func Typewriter(message string) {
	TypewriterTimed(message, 100)
}

func TypewriterTimed(message string, duration time.Duration) {
	cursorHide()
	var strChars []string = strings.Split(message, "")
	var chars int = 0
	for i := 0; i <= len(strChars); i++ {
		for _, char := range strChars[:i] {
			fmt.Print(char)
		}
		fmt.Printf("\r")
		chars++
		time.Sleep(duration * time.Millisecond)
	}
	fmt.Printf(EOL())
	cursorShow()
}

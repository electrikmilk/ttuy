package ttuy

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func File(name string) {
	bytes, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(bytes[:]))
}

func Progress(status int) {
	cursorHide()
	fmt.Printf(bar(status)+" %d%%\r", status)
	if status == 100 {
		fmt.Printf("\n")
		cursorShow()
	}
}

func bar(status int) (bar string) {
	var prog float64 = float64(status) / float64(100)
	out, err := exec.Command("tput", "cols").Output()
	if err != nil {
		panic(err)
	}
	cols, _ := strconv.Atoi(strings.TrimSuffix(string(out), "\n"))
	prog_blocks := int(prog * float64(cols))
	remain := cols - prog_blocks
	bar = "["
	for i := 0; i < prog_blocks; i++ {
		bar += "="
	}
	for i := 0; i < remain; i++ {
		bar += " "
	}
	bar += "]"
	return
}

func Ask(prompt string, input *string) {
	prompt = fmt.Sprintf(" %s ", prompt)
	fmt.Printf("\n%s ", Style(&prompt, INVERT))
	fmt.Scanln(&*input)
	fmt.Printf("\n")
}

func Typewriter(str string) {
	cursorHide()
	var strChars []string = strings.Split(str, "")
	var chars int = 0
	for i := 0; i <= len(strChars); i++ {
		for _, char := range strChars[:i] {
			fmt.Print(char)
		}
		fmt.Printf("\r")
		chars++
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("\n")
	cursorShow()
}

/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"fmt"
	"testing"
	"time"
)

func TestStyle(t *testing.T) {
	fmt.Println(Style("Text", Bold, Red))
	fmt.Println(Style("Text", Bold, Yellow))
	fmt.Println(Style("Text", Bold, Green))
	fmt.Println(Style("Text", Bold, Cyan))
	fmt.Println(Style("Text", Bold, Blue))
	fmt.Println(Style("Text", Bold, Magenta))

	fmt.Println(Style("Text", Black, Bold, RedBg))
	fmt.Println(Style("Text", Black, Bold, YellowBg))
	fmt.Println(Style("Text", Black, Bold, GreenBg))
	fmt.Println(Style("Text", Black, Bold, CyanBg))
	fmt.Println(Style("Text", Black, Bold, BlueBg))
	fmt.Println(Style("Text", Black, Bold, MagentaBg))
}

func TestProgress(t *testing.T) {
	for i := 0; i <= 100; i++ {
		Progress(i)
		time.Sleep(50 * time.Millisecond)
	}
}

func TestAsk(t *testing.T) {
	var name string
	Ask("Enter your name", &name)
	fmt.Println(name)
}

func TestTypewriter(t *testing.T) {
	Typewriter("Typed out one character at a time")
}

func TestMenu(*testing.T) {
	var chosen string
	Menu("Title", []Option{
		{
			Label: "Option 1",
			Callback: func() {
				chosen = "You chose option 1."
			},
		},
		{
			Label: "Option 2",
			Callback: func() {
				chosen = "You chose option 2."
			},
			Disabled: true,
		},
		{
			Label: "Option 3",
			Callback: func() {
				chosen = "You chose option 3."
			},
		},
	})
	fmt.Println(chosen)
}

func TestPrompt(*testing.T) {
	if Prompt("Install Program? This will take up 586kb.") {
		fmt.Println("Installing...")
	} else {
		fmt.Println("Installation Canceled.")
	}
}

func TestTable(t *testing.T) {
	var headers = []string{"Header 1", "Header 2", "Header 3", "Header 4"}
	var rows []row
	for i := 0; i < 5; i++ {
		var rowRows = []string{"Cell 1", "Cell 2", "Cell 3", "Cell 4"}
		rows = append(rows, row{columns: rowRows})
	}
	Table(headers, rows)
}

func TestFile(t *testing.T) {
	File("LICENSE")
}

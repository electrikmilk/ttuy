/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestProgressBar(t *testing.T) {
	for i := 0; i <= 100; i++ {
		ProgressBar(i)
		time.Sleep(50 * time.Millisecond)
	}
}

func TestAsk(t *testing.T) {
	input := []byte("Brandon")
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	_, err = w.Write(input)
	if err != nil {
		t.Error(err)
	}
	err = w.Close()
	if err != nil {
		panic(err)
	}

	// Restore stdin right after the test.
	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	os.Stdin = r

	var received string
	Ask("Enter your name", &received)
	fmt.Println("You entered:", received)
}

func TestAskSecret(t *testing.T) {
	input := []byte("Password123")
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	_, err = w.Write(input)
	if err != nil {
		t.Error(err)
	}
	err = w.Close()
	if err != nil {
		panic(err)
	}

	// Restore stdin right after the test.
	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	os.Stdin = r

	var received string
	AskSecret("Password", &received)
	fmt.Println("You entered:", received)
}

func TestTypewriter(t *testing.T) {
	Typewriter("Typed out one character at a time")
}

func TestTypewriterTimed(t *testing.T) {
	TypewriterTimed(":^)", 1000)
}

func TestFile(t *testing.T) {
	File("LICENSE")
}

func TestViewport(t *testing.T) {
	var bytes, err = os.ReadFile("LICENSE")
	if err != nil {
		panic(err)
	}
	Viewport(string(bytes))
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

func TestStyle(t *testing.T) {
	fmt.Println(Style("Text", Bold))
	fmt.Println(Style("Text", Italic))
	fmt.Println(Style("Text", Underlined))
	fmt.Println(Style("Text", Dim))
	fmt.Println(Style("Text", Blink))
	fmt.Println(Style("Text", Inverted))
	fmt.Println(Style("Text", Hidden))
	fmt.Println(Style("Text", Crossed))

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

func TestSpinner(t *testing.T) {
	fmt.Println("Wait a tick")
	go Spinner("Indeterminate Progress", Ticker)
	for i := 0; i < 50; i++ {
		time.Sleep(30 * time.Millisecond)
	}
	StopSpinner()
	fmt.Println("Wait a tick")
	go Spinner("Loading", DotDotDot)
	for i := 0; i < 50; i++ {
		time.Sleep(30 * time.Millisecond)
	}
	StopSpinner()
	fmt.Println("Wait a tick")
	go Spinner("Getting Ready", Throbber)
	for i := 0; i < 50; i++ {
		time.Sleep(30 * time.Millisecond)
	}
	StopSpinner()
	fmt.Println("Wait a tick")
	go Spinner("Installing", Blinker)
	for i := 0; i < 50; i++ {
		time.Sleep(30 * time.Millisecond)
	}
	StopSpinner()
	fmt.Println("Wait a tick")
}

func TestPrompt(*testing.T) {
	if Prompt("Install Program? This will take up 586kb.") {
		fmt.Println("Installing...")
	} else {
		fmt.Println("Installation Canceled.")
	}
}

func TestOutput(t *testing.T) {
	// Success
	Success("Done!")
	var success = "Done!"
	Successf("%s", success)
	// Info
	Info("FYI!")
	var info = "FYI!"
	Infof("%s", info)
	// Warning
	Warn("Unable to do something!")
	var warning = "Unable to do something!"
	Warnf("%s", warning)
	// Error
	// Fail("Unable to do something!")
	// var _, err = os.ReadDir("doesntExist")
	// FailErr("Label", err)
	// var message = "Unable to do something!"
	// Failf("%s", message)
}

func TestMenu(*testing.T) {
	var chosen string
	Menu("Title", []Option{
		{
			Label: "Option 1",
			Callback: func() {
				chosen = "You have chosen option 1."
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
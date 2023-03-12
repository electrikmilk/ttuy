/*
 * Copyright (c) 2023 Brandon Jordan
 */

package ttuy

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/eiannone/keyboard"
)

func TestProgressBar(t *testing.T) {
	for i := 0; i <= 100; i++ {
		ProgressBar(i)
		time.Sleep(100 * time.Millisecond)
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

func TestYesNo(t *testing.T) {
	input := []byte("y")
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
	YesNo("Are you sure?")
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
	var rows []Row
	for i := 0; i < 5; i++ {
		var rowRows = []string{"Cell 1", "Cell 2", "Cell 3", "Cell 4"}
		rows = append(rows, Row{Columns: rowRows})
	}
	Table(headers, rows)
}

func TestGrid(t *testing.T) {
	var alphabet = strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	var rows []Row
	for i := 0; i < 10; i++ {
		var rowRows []string
		for c := 1; c < 5; c++ {
			var cellContent string
			for w := 0; w < 26; w++ {
				cellContent += alphabet[w]
			}
			rowRows = append(rowRows, cellContent)
		}
		rows = append(rows, Row{Columns: rowRows})
	}
	fmt.Println(Grid(rows))
}

func TestStyle(t *testing.T) {
	fmt.Println(Foreground("Text", 176))
	fmt.Println(Background("Text", 127))

	fmt.Println(Style("Text", Bold))
	fmt.Println(Style("Text", Italic))
	fmt.Println(Style("Text", Underlined))
	fmt.Println(Style("Text", Dim))
	fmt.Println(Style("Text", Blink))
	fmt.Println(Style("Text", Inverted))
	fmt.Println(Style("Text", Hidden))
	fmt.Println(Style("Text", Crossed))

	fmt.Println(Style("Text", Bold, RedText))
	fmt.Println(Style("Text", Bold, YellowText))
	fmt.Println(Style("Text", Bold, GreenText))
	fmt.Println(Style("Text", Bold, CyanText))
	fmt.Println(Style("Text", Bold, BlueText))
	fmt.Println(Style("Text", Bold, MagentaText))

	fmt.Println(Style("Text", BlackText, Bold, RedBg))
	fmt.Println(Style("Text", BlackText, Bold, YellowBg))
	fmt.Println(Style("Text", BlackText, Bold, GreenBg))
	fmt.Println(Style("Text", BlackText, Bold, CyanBg))
	fmt.Println(Style("Text", BlackText, Bold, BlueBg))
	fmt.Println(Style("Text", BlackText, Bold, MagentaBg))
}

func TestStylePersist(t *testing.T) {
	fmt.Print(StylePersist(YellowText, GrayBg))
	fmt.Println("This should be in the set style." + Reset)
	fmt.Println("This should be reset.")
}

func TestSpinner(t *testing.T) {
	go Spinner("Indeterminate Progress", Ticker)
	for i := 0; i < 50; i++ {
		time.Sleep(30 * time.Millisecond)
	}
	StopSpinner()
	go Spinner("Loading", DotDotDot)
	for i := 0; i < 50; i++ {
		time.Sleep(30 * time.Millisecond)
	}
	StopSpinner()
	go Spinner("Getting Ready", Throbber)
	for i := 0; i < 50; i++ {
		time.Sleep(30 * time.Millisecond)
	}
	StopSpinner()
	go Spinner("Installing", Blinker)
	for i := 0; i < 50; i++ {
		time.Sleep(30 * time.Millisecond)
	}
	StopSpinner()
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

func TestPainter(*testing.T) {
	var lastText string
	go readKeys(func(key any) {
		if key == keyboard.KeyCtrlC {
			StopPainting()
		}
	})
	Painter(func() (template string) {
		var text string
		switch rand.Intn(4) {
		case 1:
			text = "Painter Test - Press ^C to exit"
		case 2:
			text = "Painter Test\nPainter Test\nPress ^C to exit"
		case 3:
			text = "Painter Test\nPress ^C to exit"
		}
		if text != lastText {
			template = text
			lastText = template
		} else {
			template = lastText
		}
		return
	})
}

func TestSegmented(t *testing.T) {
	fmt.Print(Segmented(fmt.Sprintf("%d:%d:%d", rand.Intn(100), rand.Intn(100), rand.Intn(100))))
}

func TestPrintBanner(t *testing.T) {
	fmt.Print(Style(BigBanner("TESTING"), RedText))
	fmt.Print(Style(Banner("TESTING"), RedText))
}

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

func TestFile(t *testing.T) {
	File("LICENSE")
}

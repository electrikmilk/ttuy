/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"fmt"
	"testing"
	"time"
)

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

func TestFile(t *testing.T) {
	File("LICENSE")
}

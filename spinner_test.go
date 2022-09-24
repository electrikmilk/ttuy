/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"fmt"
	"testing"
	"time"
)

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

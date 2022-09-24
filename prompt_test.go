/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"fmt"
	"testing"
)

func TestPrompt(*testing.T) {
	if Prompt("Install Program? This will take up 586kb.") {
		fmt.Println("Installing...")
	} else {
		fmt.Println("Installation Canceled.")
	}
}

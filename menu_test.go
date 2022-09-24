/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"fmt"
	"testing"
)

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

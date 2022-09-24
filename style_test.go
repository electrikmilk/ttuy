/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"fmt"
	"testing"
)

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

/*
 * Copyright (c) 2023 Brandon Jordan
 */

package ttuy

import "strings"

type segments struct {
	a string
	b string
	c string
	d string
	e string
	f string
	g string
	h string
}

var numbers map[string]segments

// Segmented transforms n into a series of characters resembling a segmented display.
// It accepts numbers 0-9 and colons. Example: "20:00"
func Segmented(n string) string {
	if len(numbers) == 0 {
		numbers = make(map[string]segments)
	}
	makeNumbers()
	var nums = strings.Split(n, "")
	var line1 string
	var line2 string
	var line3 string
	var line4 string
	for _, number := range nums {
		line1 += " " + numbers[number].a + "  "
	}
	for _, number := range nums {
		line2 += numbers[number].f + numbers[number].g + numbers[number].b + " "
	}
	for _, number := range nums {
		line3 += numbers[number].e + numbers[number].h + numbers[number].c + " "
	}
	for _, number := range nums {
		line4 += " " + numbers[number].d + "  "
	}

	return strings.Join([]string{line1, line2, line3, line4}, eol)
}

func makeNumbers() {
	numbers["0"] = segments{
		a: lower,
		b: line,
		c: line,
		d: upper,
		e: line,
		f: line,
		g: " ",
		h: " ",
	}
	numbers["1"] = segments{
		a: " ",
		b: line,
		c: line,
		d: " ",
		e: " ",
		f: " ",
		g: " ",
		h: " ",
	}
	numbers["2"] = segments{
		a: lower,
		b: line,
		c: " ",
		d: upper,
		e: line,
		f: " ",
		g: lower,
		h: " ",
	}
	numbers["3"] = segments{
		a: lower,
		b: line,
		c: line,
		d: upper,
		e: " ",
		f: " ",
		g: lower,
		h: " ",
	}
	numbers["4"] = segments{
		a: " ",
		b: line,
		c: line,
		d: " ",
		e: " ",
		f: line,
		g: lower,
		h: " ",
	}
	numbers["5"] = segments{
		a: lower,
		b: " ",
		c: line,
		d: upper,
		e: " ",
		f: line,
		g: lower,
		h: " ",
	}
	numbers["6"] = segments{
		a: lower,
		b: " ",
		c: line,
		d: upper,
		e: line,
		f: line,
		g: lower,
		h: " ",
	}
	numbers["7"] = segments{
		a: lower,
		b: line,
		c: line,
		d: " ",
		e: " ",
		f: " ",
		g: " ",
		h: " ",
	}
	numbers["8"] = segments{
		a: lower,
		b: line,
		c: line,
		d: upper,
		e: line,
		f: line,
		g: lower,
		h: " ",
	}
	numbers["9"] = segments{
		a: lower,
		b: line,
		c: line,
		d: upper,
		e: " ",
		f: line,
		g: lower,
		h: " ",
	}
	numbers[":"] = segments{
		a: " ",
		b: " ",
		c: " ",
		d: " ",
		e: " ",
		f: " ",
		g: bullet,
		h: bullet,
	}
}

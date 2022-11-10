/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

const (
	upper = "\u203E"
	lower = "_"
	line  = "|"
)

type segments struct {
	a string
	b string
	c string
	d string
	e string
	f string
	g string
}

var numbers map[int]segments

func Segmented(nums ...int) string {
	if len(numbers) == 0 {
		numbers = make(map[int]segments)
	}
	makeNumbers()
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
		line3 += numbers[number].e + " " + numbers[number].c + " "
	}
	for _, number := range nums {
		line4 += " " + numbers[number].d + "  "
	}
	var segments string
	segments += line1 + eol
	segments += line2 + eol
	segments += line3 + eol
	segments += line4 + eol
	return segments
	// Painter(func() string {
	// 	return ""
	// })
}

func makeNumbers() {
	numbers[0] = segments{
		a: lower,
		b: line,
		c: line,
		d: upper,
		e: line,
		f: line,
		g: " ",
	}
	numbers[1] = segments{
		a: " ",
		b: line,
		c: line,
		d: " ",
		e: " ",
		f: " ",
		g: " ",
	}
	numbers[2] = segments{
		a: lower,
		b: line,
		c: " ",
		d: upper,
		e: line,
		f: " ",
		g: lower,
	}
	numbers[3] = segments{
		a: lower,
		b: line,
		c: line,
		d: upper,
		e: " ",
		f: " ",
		g: lower,
	}
	numbers[4] = segments{
		a: " ",
		b: line,
		c: line,
		d: " ",
		e: " ",
		f: line,
		g: lower,
	}
	numbers[5] = segments{
		a: lower,
		b: " ",
		c: line,
		d: upper,
		e: " ",
		f: line,
		g: lower,
	}
	numbers[6] = segments{
		a: lower,
		b: " ",
		c: line,
		d: upper,
		e: line,
		f: line,
		g: lower,
	}
	numbers[7] = segments{
		a: lower,
		b: line,
		c: line,
		d: " ",
		e: " ",
		f: " ",
		g: " ",
	}
	numbers[8] = segments{
		a: lower,
		b: line,
		c: line,
		d: upper,
		e: line,
		f: line,
		g: lower,
	}
	numbers[9] = segments{
		a: lower,
		b: line,
		c: line,
		d: upper,
		e: " ",
		f: line,
		g: lower,
	}
}

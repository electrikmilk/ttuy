/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"os"
	"testing"
)

func TestViewport(t *testing.T) {
	var bytes, err = os.ReadFile("LICENSE")
	if err != nil {
		panic(err)
	}
	Viewport(string(bytes))
}

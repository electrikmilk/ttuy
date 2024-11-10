/*
 * Copyright (c) 2023 Brandon Jordan
 */

package ttuy

import (
	"github.com/eiannone/keyboard"
)

type inputCallback func(key any)

func ReadKeys(callback inputCallback) {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		callback(key)
		if key == keyboard.KeyEnter {
			_ = keyboard.Close()
			break
		}
	}
}

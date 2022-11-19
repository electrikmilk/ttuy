/*
 * Copyright (c) 2022 Brandon Jordan
 */

package ttuy

import (
	"github.com/eiannone/keyboard"
)

type InputCallback func(key any)

func ReadKeys(callback InputCallback) {
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

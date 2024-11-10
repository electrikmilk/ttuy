/*
 * Copyright (c) 2024 Brandon Jordan
 */

package ttuy

import (
	"github.com/eiannone/keyboard"
	"slices"
	"strings"
)

var keyMap map[keyboard.Key]KeyBinding
var listeningKeys = false
var labels []string

// KeyBinding describes a key binding to an action.
type KeyBinding struct {
	// Description of what the handler will do.
	descriptor string

	// Representation of key (e.g. ^C).
	keyString string

	// Called when the key is pressed.
	handler func()
}

// ReadKeys abstracts the keyboard package.
func ReadKeys(callback func(key keyboard.Key)) {
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

// BindKey binds key to binding.
func BindKey(key keyboard.Key, binding KeyBinding) {
	if !listeningKeys {
		keyMap = make(map[keyboard.Key]KeyBinding)
		initKeyboard()
		listeningKeys = true
	}

	keyMap[key] = binding
}

func initKeyboard() {
	go ReadKeys(func(key keyboard.Key) {
		if binding, ok := keyMap[key]; ok {
			binding.handler()
		}
	})
}

// PrintKeyBindings prints the registered key bindings between separator.
func PrintKeyBindings(separator string) string {
	labels = []string{}
	for _, binding := range keyMap {
		labels = append(labels, keyLabel(binding))
	}

	slices.Sort(labels)

	return strings.Join(labels, Style(separator, Dim))
}

func keyLabel(kb KeyBinding) string {
	var label strings.Builder
	label.WriteString(Style(kb.keyString, Dim))
	label.WriteRune(' ')
	label.WriteString(Style(kb.descriptor, GrayText))

	return label.String()
}

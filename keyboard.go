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

type KeyBinding struct {
	descriptor string
	keyString  string
	handler    func()
}

func BindKey(key keyboard.Key, binding KeyBinding) {
	if !listeningKeys {
		keyMap = make(map[keyboard.Key]KeyBinding)
		initKeyboard()
		listeningKeys = true
	}

	keyMap[key] = binding
}

func Hotkeys(separator string) string {
	labels = []string{}
	for _, keybind := range keyMap {
		labels = append(labels, keyLabel(keybind))
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

func initKeyboard() {
	go ReadKeys(func(key keyboard.Key) {
		if binding, ok := keyMap[key]; ok {
			binding.handler()
		}
	})
}

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

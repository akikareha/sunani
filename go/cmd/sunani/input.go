package main

import (
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tetratelabs/wazero/api"
)

type Input struct {
	window *glfw.Window

	key api.Function
}

func NewInput() *Input {
	return &Input{}
}

func (in *Input) Preinit() {
	in.key = mod.ExportedFunction("sunani_input_key")
}

func (in *Input) IsEnabled() bool {
	return in.key != nil
}

func (in *Input) Init(window *glfw.Window) {
	if !in.IsEnabled() {
		return
	}

	in.window = window

	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		k := mapGLFWKey(key)
		a := mapGLFWAction(action)

		_, err := in.key.Call(ctx, uint64(k), uint64(a))
		if err != nil {
			log.Fatalln("input key call failed:", err)
		}
	})
}

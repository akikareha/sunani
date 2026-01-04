package main

import (
	"log"
	"math"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tetratelabs/wazero/api"
)

type Input struct {
	window *glfw.Window

	key   api.Function
	mouseMotion api.Function
	mouseButton api.Function
}

func NewInput() *Input {
	return &Input{}
}

func (in *Input) Preinit() {
	in.key = mod.ExportedFunction("sunani_input_key")
	in.mouseMotion = mod.ExportedFunction("sunani_input_mouse_motion")
	in.mouseButton = mod.ExportedFunction("sunani_input_mouse_button")
}

func (in *Input) IsKeyEnabled() bool {
	return in.key != nil
}

func (in *Input) IsMouseMotionEnabled() bool {
	return in.mouseMotion != nil
}

func (in *Input) IsMouseButtonEnabled() bool {
	return in.mouseButton != nil
}

func (in *Input) IsMouseEnabled() bool {
	return in.IsMouseMotionEnabled() || in.IsMouseButtonEnabled()
}

func (in *Input) IsEnabled() bool {
	return in.IsKeyEnabled() || in.IsMouseEnabled()
}

func (in *Input) Init(window *glfw.Window) {
	if !in.IsEnabled() {
		return
	}

	in.window = window

	if in.IsKeyEnabled() {
		window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
			k := mapGLFWKey(key)
			a := mapGLFWAction(action)

			_, err := in.key.Call(ctx, uint64(k), uint64(a))
			if err != nil {
				log.Fatalln("input key call failed:", err)
			}
		})
	}

	if in.IsMouseMotionEnabled() {
		window.SetCursorPosCallback(func(w *glfw.Window, x, y float64) {
			_, err := in.mouseMotion.Call(
				ctx,
				uint64(math.Float32bits(float32(x))),
				uint64(math.Float32bits(float32(y))),
			)
			if err != nil {
				log.Fatalln("input mouse call failed:", err)
			}
		})
	}

	if in.IsMouseButtonEnabled() {
		window.SetMouseButtonCallback(func(
			w *glfw.Window,
			button glfw.MouseButton,
			action glfw.Action,
			mods glfw.ModifierKey,
		) {
			b := mapGLFWMouseButton(button)
			a := mapGLFWAction(action)

			_, err := in.mouseButton.Call(
				ctx,
				uint64(b),
				uint64(a),
			)
			if err != nil {
				log.Fatalln("input mouse call failed:", err)
			}
		})
	}
}

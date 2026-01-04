package main

import (
	"log"
	"math"

	"github.com/go-gl/glfw/v3.3/glfw"
	wapi "github.com/tetratelabs/wazero/api"

	"github.com/akikareha/sunani/go/api"
)

type Input struct {
	window *glfw.Window

	key   wapi.Function
	mouse wapi.Function

	hasPrev bool
	prevX   float64
	prevY   float64
}

func NewInput() *Input {
	return &Input{}
}

func (in *Input) Preinit() {
	in.key = mod.ExportedFunction("sunani_input_key")
	in.mouse = mod.ExportedFunction("sunani_input_mouse")
}

func (in *Input) IsKeyEnabled() bool {
	return in.key != nil
}

func (in *Input) IsMouseEnabled() bool {
	return in.mouse != nil
}

func (in *Input) IsEnabled() bool {
	return in.key != nil || in.mouse != nil
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

	if in.IsMouseEnabled() {
		window.SetCursorPosCallback(func(w *glfw.Window, x, y float64) {
			if !in.hasPrev {
				in.prevX = x
				in.prevY = y

				in.hasPrev = true
			}

			dx := x - in.prevX
			dy := y - in.prevY

			wheelX := 0
			wheelY := 0

			_, err := in.mouse.Call(
				ctx,
				uint64(api.MouseMove),
				uint64(math.Float32bits(float32(x))),
				uint64(math.Float32bits(float32(y))),
				uint64(math.Float32bits(float32(dx))),
				uint64(math.Float32bits(float32(dy))),
				uint64(math.Float32bits(float32(wheelX))),
				uint64(math.Float32bits(float32(wheelY))),
				uint64(api.MouseButtonUnknown),
			)
			if err != nil {
				log.Fatalln("input mouse call failed:", err)
			}

			in.prevX = x
			in.prevY = y
		})

		window.SetMouseButtonCallback(func(
			w *glfw.Window,
			button glfw.MouseButton,
			action glfw.Action,
			mods glfw.ModifierKey,
		) {
			a := mapGLFWMouseAction(action)

			dx := 0
			dy := 0

			wheelX := 0
			wheelY := 0

			b := mapGLFWMouseButton(button)

			_, err := in.mouse.Call(
				ctx,
				uint64(a),
				uint64(math.Float32bits(float32(in.prevX))),
				uint64(math.Float32bits(float32(in.prevY))),
				uint64(math.Float32bits(float32(dx))),
				uint64(math.Float32bits(float32(dy))),
				uint64(math.Float32bits(float32(wheelX))),
				uint64(math.Float32bits(float32(wheelY))),
				uint64(b),
			)
			if err != nil {
				log.Fatalln("input mouse call failed:", err)
			}
		})
	}
}

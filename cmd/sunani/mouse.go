package main

import (
	"log"
	"math"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tetratelabs/wazero/api"
)

type Mouse struct {
	window *glfw.Window

	init api.Function
	motion api.Function
	button api.Function
	wheel api.Function
}

func NewMouse() *Mouse {
	return &Mouse{}
}

func (m *Mouse) Preinit() {
	m.init = mod.ExportedFunction("sunani_mouse_init")
	m.motion = mod.ExportedFunction("sunani_mouse_motion")
	m.button = mod.ExportedFunction("sunani_mouse_button")
	m.wheel = mod.ExportedFunction("sunani_mouse_wheel")
}

func (m *Mouse) IsEnabled() bool {
	return m.init != nil
}

func (m *Mouse) Init(window *glfw.Window) {
	if !m.IsEnabled() {
		return
	}

	m.window = window

	window.SetCursorPosCallback(func(
		w *glfw.Window,
		x, y float64,
	) {
		_, err := m.motion.Call(
			ctx,
			uint64(math.Float32bits(float32(x))),
			uint64(math.Float32bits(float32(y))),
		)
		if err != nil {
			log.Fatalln("mouse motion call failed:", err)
		}
	})

	window.SetMouseButtonCallback(func(
		w *glfw.Window,
		button glfw.MouseButton,
		action glfw.Action,
		mods glfw.ModifierKey,
	) {
		b := mapGLFWMouseButton(button)
		a := mapGLFWAction(action)

		_, err := m.button.Call(
			ctx,
			uint64(b),
			uint64(a),
		)
		if err != nil {
			log.Fatalln("mouse button call failed:", err)
		}
	})

	window.SetScrollCallback(func(
		w *glfw.Window,
		xoff float64,
		yoff float64,
	) {
		_, err := m.wheel.Call(
			ctx,
			uint64(math.Float32bits(float32(xoff))),
			uint64(math.Float32bits(float32(yoff))),
		)
		if err != nil {
			log.Fatalln("mouse wheel call failed:", err)
		}
	})
}

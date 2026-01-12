package main

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tetratelabs/wazero/api"
)

type Mouse struct {
	init   api.Function
	motion api.Function
	button api.Function
	wheel  api.Function

	window *glfw.Window
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
	if window == nil {
		panic("window is nil")
	}
	m.window = window

	if m.init != nil {
		_, err := m.init.Call(ctx)
		if err != nil {
			die("sunani_mouse_init call failed:", err)
		}
	}

	if m.motion != nil {
		window.SetCursorPosCallback(func(
			w *glfw.Window,
			x, y float64,
		) {
			_, err := m.motion.Call(
				ctx,
				uint64(x),
				uint64(y),
			)
			if err != nil {
				die("sunani_mouse_motion call failed:", err)
			}
		})
	}

	if m.button != nil {
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
				die("sunani_mouse_button call failed:", err)
			}
		})
	}

	if m.wheel != nil {
		window.SetScrollCallback(func(
			w *glfw.Window,
			xoff, yoff float64,
		) {
			_, err := m.wheel.Call(
				ctx,
				uint64(xoff),
				uint64(yoff),
			)
			if err != nil {
				die("sunani_mouse_wheel call failed:", err)
			}
		})
	}
}

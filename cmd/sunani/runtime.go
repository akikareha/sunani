package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tetratelabs/wazero/api"
)

const appTitle = "Sunani"

type Runtime struct {
	init   api.Function
	resize api.Function
	frame  api.Function

	window *glfw.Window
}

func NewRuntime() *Runtime {
	return &Runtime{}
}

func (rt *Runtime) Preinit() {
	rt.init = mod.ExportedFunction("sunani_runtime_init")
	rt.resize = mod.ExportedFunction("sunani_runtime_resize")
	rt.frame = mod.ExportedFunction("sunani_runtime_frame")
}

func (rt *Runtime) IsEnabled() bool {
	return rt.init != nil
}

func (rt *Runtime) doResize(width, height uint32) {
	_, err := rt.resize.Call(
		ctx,
		uint64(width),
		uint64(height),
	)
	if err != nil {
		die("sunani_runtime_resize call failed:", err)
	}
}

func (rt *Runtime) Init(window *glfw.Window) {
	if !rt.IsEnabled() {
		return
	}
	if window == nil {
		panic("window is nil")
	}
	rt.window = window

	if rt.init != nil {
		_, err := rt.init.Call(ctx)
		if err != nil {
			die("sunani_runtime_init call failed:", err)
		}
	}

	if rt.resize != nil {
		window.SetFramebufferSizeCallback(func(
			w *glfw.Window,
			width, height int,
		) {
			rt.doResize(uint32(width), uint32(height))
		})

		fbw, fbh := window.GetFramebufferSize()
		rt.doResize(uint32(fbw), uint32(fbh))
	}
}

func (rt *Runtime) Halt() {
	if !rt.IsEnabled() {
		errlog("sunani runtime.halt was called, but Runtime API is not enabled.\nExport snunani_runtime_init to enable this API.")
		return
	}
	if rt.window == nil {
		panic("window is nil")
	}

	rt.window.SetShouldClose(true)
}

func (rt *Runtime) Title(ptr uint32, length uint32) {
	if !rt.IsEnabled() {
		errlog("sunani runtime.halt was called, but Runtime API is not enabled.\nExport snunani_runtime_init to enable this API.")
		return
	}
	if rt.window == nil {
		panic("window is nil")
	}

	mem := mod.Memory()
	buf, ok := mem.Read(ptr, length)
	if !ok {
		errlog("mem.Read failed")
		return
	}
	title := string(buf)
	if title == "" {
		title = appTitle
	} else {
		title += " - " + appTitle
	}
	rt.window.SetTitle(title)
}

func (rt *Runtime) Cursor(enabled uint32) {
	if !rt.IsEnabled() {
		errlog("sunani runtime.halt was called, but Runtime API is not enabled.\nExport snunani_runtime_init to enable this API.")
		return
	}
	if rt.window == nil {
		panic("window is nil")
	}

	if enabled == 0 {
		rt.window.SetInputMode(glfw.CursorMode, glfw.CursorHidden)
	} else {
		rt.window.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
	}
}

func (rt *Runtime) Frame() {
	if !rt.IsEnabled() {
		return
	}
	if rt.window == nil {
		panic("window is nil")
	}

	if rt.frame != nil {
		_, err := rt.frame.Call(ctx)
		if err != nil {
			die("sunani_runtime_frame call failed:", err)
		}
	}
}

func (rt *Runtime) Clear(r, g, b, a uint32) {
	if !rt.IsEnabled() {
		errlog("sunani runtime.clear was called, but Runtime API is not enabled.\nExport snunani_runtime_init to enable this API.")
		return
	}
	if rt.window == nil {
		panic("window is nil")
	}

	gl.ClearColor(
		float32(r)/255,
		float32(g)/255,
		float32(b)/255,
		float32(a)/255,
	)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

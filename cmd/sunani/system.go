package main

import (
	"math"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tetratelabs/wazero/api"
)

const appTitle = "Sunani"

type System struct {
	init   api.Function
	resize api.Function
	frame  api.Function

	window *glfw.Window
}

func NewSystem() *System {
	return &System{}
}

func (sys *System) Preinit() {
	sys.init = mod.ExportedFunction("sunani_system_init")
	sys.resize = mod.ExportedFunction("sunani_system_resize")
	sys.frame = mod.ExportedFunction("sunani_system_frame")
}

func (sys *System) IsEnabled() bool {
	return sys.init != nil
}

func (sys *System) doResize(width, height float32) {
	_, err := sys.resize.Call(
		ctx,
		uint64(math.Float32bits(width)),
		uint64(math.Float32bits(height)),
	)
	if err != nil {
		die("sunani_system_resize call failed:", err)
	}
}

func (sys *System) Init(window *glfw.Window) {
	if !sys.IsEnabled() {
		return
	}
	if window == nil {
		panic("window is nil")
	}
	sys.window = window

	if sys.init != nil {
		_, err := sys.init.Call(ctx)
		if err != nil {
			die("sunani_system_init call failed:", err)
		}
	}

	if sys.resize != nil {
		window.SetFramebufferSizeCallback(func(
			w *glfw.Window,
			width, height int,
		) {
			sys.doResize(float32(width), float32(height))
		})

		fbw, fbh := window.GetFramebufferSize()
		sys.doResize(float32(fbw), float32(fbh))
	}
}

func (sys *System) Halt() {
	if !sys.IsEnabled() {
		errlog("sunani system.halt was called, but System API is not enabled.\nExport snunani_system_init to enable this API.")
		return
	}
	if sys.window == nil {
		panic("window is nil")
	}

	sys.window.SetShouldClose(true)
}

func (sys *System) Title(ptr uint32, length uint32) {
	if !sys.IsEnabled() {
		errlog("sunani system.halt was called, but System API is not enabled.\nExport snunani_system_init to enable this API.")
		return
	}
	if sys.window == nil {
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
	sys.window.SetTitle(title)
}

func (sys *System) Cursor(enabled uint32) {
	if !sys.IsEnabled() {
		errlog("sunani system.halt was called, but System API is not enabled.\nExport snunani_system_init to enable this API.")
		return
	}
	if sys.window == nil {
		panic("window is nil")
	}

	if enabled == 0 {
		sys.window.SetInputMode(glfw.CursorMode, glfw.CursorHidden)
	} else {
		sys.window.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
	}
}

func (sys *System) Frame() {
	if !sys.IsEnabled() {
		return
	}
	if sys.window == nil {
		panic("window is nil")
	}

	if sys.frame != nil {
		_, err := sys.frame.Call(ctx)
		if err != nil {
			die("sunani_system_frame call failed:", err)
		}
	}
}

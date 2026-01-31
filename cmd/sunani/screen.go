package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tetratelabs/wazero/api"
)

type Screen struct {
	init   api.Function
	resize api.Function
	frame  api.Function

	window *glfw.Window
}

func NewScreen() *Screen {
	return &Screen{}
}

func (scr *Screen) Preinit() {
	scr.init = mod.ExportedFunction("sunani_screen_init")
	scr.resize = mod.ExportedFunction("sunani_screen_resize")
	scr.frame = mod.ExportedFunction("sunani_screen_frame")
}

func (scr *Screen) IsEnabled() bool {
	return scr.init != nil
}

func (scr *Screen) DoResize(width, height int) {
	if scr.resize == nil {
		return
	}
	_, err := scr.resize.Call(
		ctx,
		uint64(width),
		uint64(height),
	)
	if err != nil {
		die("sunani_screen_resize call failed:", err)
	}
}

func (scr *Screen) Init(window *glfw.Window) {
	if !scr.IsEnabled() {
		return
	}
	if window == nil {
		panic("window is nil")
	}
	scr.window = window

	if scr.init != nil {
		_, err := scr.init.Call(ctx)
		if err != nil {
			die("sunani_screen_init call failed:", err)
		}
	}

	fbw, fbh := window.GetFramebufferSize()
	scr.DoResize(fbw, fbh)
}

func (scr *Screen) Halt() {
	if !scr.IsEnabled() {
		errlog("sunani screen.halt was called, but Screen API is not enabled.\nExport snunani_screen_init to enable this API.")
		return
	}
	if scr.window == nil {
		panic("window is nil")
	}

	scr.window.SetShouldClose(true)
}

func (scr *Screen) Title(ptr uint32, length uint32) {
	if !scr.IsEnabled() {
		errlog("sunani screen.halt was called, but Screen API is not enabled.\nExport snunani_screen_init to enable this API.")
		return
	}
	if scr.window == nil {
		panic("window is nil")
	}

	if length < 1 {
		scr.window.SetTitle("")
		return
	}
	mem := mod.Memory()
	buf, ok := mem.Read(ptr, length)
	if !ok {
		errlog("mem.Read failed")
		return
	}
	title := string(buf)
	scr.window.SetTitle(title)
}

func (scr *Screen) Cursor(visible bool) {
	if !scr.IsEnabled() {
		errlog("sunani screen.halt was called, but Screen API is not enabled.\nExport snunani_screen_init to enable this API.")
		return
	}
	if scr.window == nil {
		panic("window is nil")
	}

	if visible {
		scr.window.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
	} else {
		scr.window.SetInputMode(glfw.CursorMode, glfw.CursorHidden)
	}
}

func (scr *Screen) Frame() {
	if !scr.IsEnabled() {
		return
	}
	if scr.window == nil {
		panic("window is nil")
	}

	if scr.frame != nil {
		_, err := scr.frame.Call(ctx)
		if err != nil {
			die("sunani_screen_frame call failed:", err)
		}
	}
}

func (scr *Screen) Clear(r, g, b, a int) {
	if !scr.IsEnabled() {
		errlog("sunani screen.clear was called, but Screen API is not enabled.\nExport snunani_screen_init to enable this API.")
		return
	}
	if scr.window == nil {
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

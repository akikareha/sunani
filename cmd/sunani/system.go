package main

import (
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tetratelabs/wazero/api"
)

type System struct {
	window *glfw.Window

	init  api.Function
	frame api.Function
}

func NewSystem() *System {
	return &System{}
}

func (sys *System) Preinit() {
	sys.init = mod.ExportedFunction("sunani_system_init")
	sys.frame = mod.ExportedFunction("sunani_system_frame")
}

func (sys *System) IsEnabled() bool {
	return sys.init != nil
}

func (sys *System) Init(window *glfw.Window) {
	if !sys.IsEnabled() {
		return
	}

	sys.window = window

	_, err := sys.init.Call(ctx)
	if err != nil {
		log.Fatalln("system init call failed:", err)
	}
}

func (sys *System) Halt() {
	if !sys.IsEnabled() {
		return
	}

	sys.window.SetShouldClose(true)
}

func (sys *System) Title(ptr uint32, length uint32) {
	if !sys.IsEnabled() {
		return
	}
	if sys.window == nil {
		return
	}

	mem := mod.Memory()
	buf, ok := mem.Read(ptr, length)
	if !ok {
		return
	}
	title := string(buf)
	if title == "" {
		title = "Sunani"
	} else {
		title += " - Sunani"
	}
	sys.window.SetTitle(title)
}

func (sys *System) Cursor(enabled uint32) {
	if !sys.IsEnabled() {
		return
	}
	if sys.window == nil {
		return
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

	_, err := sys.frame.Call(ctx)
	if err != nil {
		log.Fatalln("system frame call failed:", err)
	}
}

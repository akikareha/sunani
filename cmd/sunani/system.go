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
}

func (sys *System) Quit() {
	if !sys.IsEnabled() {
		return
	}

	sys.window.SetShouldClose(true)
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

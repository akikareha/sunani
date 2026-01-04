package main

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

type System struct {
	window *glfw.Window
}

func NewSystem() *System {
	return &System{}
}

func (sys *System) Init(window *glfw.Window) {
	sys.window = window
}

func (sys *System) Quit() {
	if sys.window == nil {
		return
	}

	sys.window.SetShouldClose(true)
}

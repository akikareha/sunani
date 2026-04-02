package main

import (
	"github.com/go-gl/glfw/v3.3/glfw"

	"tea.kareha.org/loom/sunani/lib"
)

func mapGLFWAction(a glfw.Action) lib.Action {
	switch a {
	case glfw.Press:
		return lib.ActionPress
	case glfw.Release:
		return lib.ActionRelease

	default:
		return lib.ActionUnknown
	}
}

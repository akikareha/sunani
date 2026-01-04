package main

import (
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/akikareha/sunani/input"
)

func mapGLFWAction(a glfw.Action) input.Action {
	switch a {
	case glfw.Press:
		return input.ActionPress
	case glfw.Release:
		return input.ActionRelease

	default:
		return input.ActionUnknown
	}
}

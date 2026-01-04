package main

import (
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/akikareha/sunani/api"
)

func mapGLFWAction(a glfw.Action) api.Action {
	switch a {
	case glfw.Press:
		return api.ActionPress
	case glfw.Release:
		return api.ActionRelease

	default:
		return api.ActionUnknown
	}
}

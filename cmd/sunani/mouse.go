package main

import (
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/akikareha/sunani/api"
)

func mapGLFWMouseButton(b glfw.MouseButton) api.Mouse {
	switch b {
	case glfw.MouseButtonLeft:
		return api.MouseLeft
	case glfw.MouseButtonMiddle:
		return api.MouseMiddle
	case glfw.MouseButtonRight:
		return api.MouseRight

	default:
		return api.MouseUnknown
	}
}

package main

import (
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/akikareha/sunani/input"
)

func mapGLFWMouseButton(b glfw.MouseButton) input.Mouse {
	switch b {
	case glfw.MouseButtonLeft:
		return input.MouseLeft
	case glfw.MouseButtonMiddle:
		return input.MouseMiddle
	case glfw.MouseButtonRight:
		return input.MouseRight

	default:
		return input.MouseUnknown
	}
}

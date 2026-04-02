package main

import (
	"github.com/go-gl/glfw/v3.3/glfw"

	"tea.kareha.org/loom/sunani/lib"
)

func mapGLFWMouseButton(b glfw.MouseButton) lib.Mouse {
	switch b {
	case glfw.MouseButtonLeft:
		return lib.MouseLeft
	case glfw.MouseButtonMiddle:
		return lib.MouseMiddle
	case glfw.MouseButtonRight:
		return lib.MouseRight

	default:
		return lib.MouseUnknown
	}
}

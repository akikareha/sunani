package main

import (
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/akikareha/sunani/input"
)

func mapGLFWKey(k glfw.Key) input.Key {
	switch k {

	// Letters
	case glfw.KeyA:
		return input.KeyA
	case glfw.KeyB:
		return input.KeyB
	case glfw.KeyC:
		return input.KeyC
	case glfw.KeyD:
		return input.KeyD
	case glfw.KeyE:
		return input.KeyE
	case glfw.KeyF:
		return input.KeyF
	case glfw.KeyG:
		return input.KeyG
	case glfw.KeyH:
		return input.KeyH
	case glfw.KeyI:
		return input.KeyI
	case glfw.KeyJ:
		return input.KeyJ
	case glfw.KeyK:
		return input.KeyK
	case glfw.KeyL:
		return input.KeyL
	case glfw.KeyM:
		return input.KeyM
	case glfw.KeyN:
		return input.KeyN
	case glfw.KeyO:
		return input.KeyO
	case glfw.KeyP:
		return input.KeyP
	case glfw.KeyQ:
		return input.KeyQ
	case glfw.KeyR:
		return input.KeyR
	case glfw.KeyS:
		return input.KeyS
	case glfw.KeyT:
		return input.KeyT
	case glfw.KeyU:
		return input.KeyU
	case glfw.KeyV:
		return input.KeyV
	case glfw.KeyW:
		return input.KeyW
	case glfw.KeyX:
		return input.KeyX
	case glfw.KeyY:
		return input.KeyY
	case glfw.KeyZ:
		return input.KeyZ

	// Digits
	case glfw.Key0:
		return input.Key0
	case glfw.Key1:
		return input.Key1
	case glfw.Key2:
		return input.Key2
	case glfw.Key3:
		return input.Key3
	case glfw.Key4:
		return input.Key4
	case glfw.Key5:
		return input.Key5
	case glfw.Key6:
		return input.Key6
	case glfw.Key7:
		return input.Key7
	case glfw.Key8:
		return input.Key8
	case glfw.Key9:
		return input.Key9

	// Control
	case glfw.KeyEscape:
		return input.KeyEscape
	case glfw.KeyEnter:
		return input.KeyEnter
	case glfw.KeySpace:
		return input.KeySpace
	case glfw.KeyTab:
		return input.KeyTab
	case glfw.KeyBackspace:
		return input.KeyBackspace

	// Arrows
	case glfw.KeyUp:
		return input.KeyUp
	case glfw.KeyDown:
		return input.KeyDown
	case glfw.KeyLeft:
		return input.KeyLeft
	case glfw.KeyRight:
		return input.KeyRight

	// Others
	default:
		return input.KeyUnknown
	}
}

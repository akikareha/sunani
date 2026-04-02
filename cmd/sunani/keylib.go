package main

import (
	"github.com/go-gl/glfw/v3.3/glfw"

	"tea.kareha.org/loom/sunani/lib"
)

func mapGLFWKey(k glfw.Key) lib.Key {
	switch k {

	// Letters
	case glfw.KeyA:
		return lib.KeyA
	case glfw.KeyB:
		return lib.KeyB
	case glfw.KeyC:
		return lib.KeyC
	case glfw.KeyD:
		return lib.KeyD
	case glfw.KeyE:
		return lib.KeyE
	case glfw.KeyF:
		return lib.KeyF
	case glfw.KeyG:
		return lib.KeyG
	case glfw.KeyH:
		return lib.KeyH
	case glfw.KeyI:
		return lib.KeyI
	case glfw.KeyJ:
		return lib.KeyJ
	case glfw.KeyK:
		return lib.KeyK
	case glfw.KeyL:
		return lib.KeyL
	case glfw.KeyM:
		return lib.KeyM
	case glfw.KeyN:
		return lib.KeyN
	case glfw.KeyO:
		return lib.KeyO
	case glfw.KeyP:
		return lib.KeyP
	case glfw.KeyQ:
		return lib.KeyQ
	case glfw.KeyR:
		return lib.KeyR
	case glfw.KeyS:
		return lib.KeyS
	case glfw.KeyT:
		return lib.KeyT
	case glfw.KeyU:
		return lib.KeyU
	case glfw.KeyV:
		return lib.KeyV
	case glfw.KeyW:
		return lib.KeyW
	case glfw.KeyX:
		return lib.KeyX
	case glfw.KeyY:
		return lib.KeyY
	case glfw.KeyZ:
		return lib.KeyZ

	// Digits
	case glfw.Key0:
		return lib.Key0
	case glfw.Key1:
		return lib.Key1
	case glfw.Key2:
		return lib.Key2
	case glfw.Key3:
		return lib.Key3
	case glfw.Key4:
		return lib.Key4
	case glfw.Key5:
		return lib.Key5
	case glfw.Key6:
		return lib.Key6
	case glfw.Key7:
		return lib.Key7
	case glfw.Key8:
		return lib.Key8
	case glfw.Key9:
		return lib.Key9

	// Control
	case glfw.KeyEscape:
		return lib.KeyEscape
	case glfw.KeyEnter:
		return lib.KeyEnter
	case glfw.KeySpace:
		return lib.KeySpace
	case glfw.KeyTab:
		return lib.KeyTab
	case glfw.KeyBackspace:
		return lib.KeyBackspace

	// Arrows
	case glfw.KeyUp:
		return lib.KeyUp
	case glfw.KeyDown:
		return lib.KeyDown
	case glfw.KeyLeft:
		return lib.KeyLeft
	case glfw.KeyRight:
		return lib.KeyRight

	// Others
	default:
		return lib.KeyUnknown
	}
}

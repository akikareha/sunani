package main

import (
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/akikareha/sunani/go/api"
)

func mapGLFWKey(k glfw.Key) api.Key {
	switch k {

	// Letters
	case glfw.KeyA:
		return api.KeyA
	case glfw.KeyB:
		return api.KeyB
	case glfw.KeyC:
		return api.KeyC
	case glfw.KeyD:
		return api.KeyD
	case glfw.KeyE:
		return api.KeyE
	case glfw.KeyF:
		return api.KeyF
	case glfw.KeyG:
		return api.KeyG
	case glfw.KeyH:
		return api.KeyH
	case glfw.KeyI:
		return api.KeyI
	case glfw.KeyJ:
		return api.KeyJ
	case glfw.KeyK:
		return api.KeyK
	case glfw.KeyL:
		return api.KeyL
	case glfw.KeyM:
		return api.KeyM
	case glfw.KeyN:
		return api.KeyN
	case glfw.KeyO:
		return api.KeyO
	case glfw.KeyP:
		return api.KeyP
	case glfw.KeyQ:
		return api.KeyQ
	case glfw.KeyR:
		return api.KeyR
	case glfw.KeyS:
		return api.KeyS
	case glfw.KeyT:
		return api.KeyT
	case glfw.KeyU:
		return api.KeyU
	case glfw.KeyV:
		return api.KeyV
	case glfw.KeyW:
		return api.KeyW
	case glfw.KeyX:
		return api.KeyX
	case glfw.KeyY:
		return api.KeyY
	case glfw.KeyZ:
		return api.KeyZ

	// Digits
	case glfw.Key0:
		return api.Key0
	case glfw.Key1:
		return api.Key1
	case glfw.Key2:
		return api.Key2
	case glfw.Key3:
		return api.Key3
	case glfw.Key4:
		return api.Key4
	case glfw.Key5:
		return api.Key5
	case glfw.Key6:
		return api.Key6
	case glfw.Key7:
		return api.Key7
	case glfw.Key8:
		return api.Key8
	case glfw.Key9:
		return api.Key9

	// Control
	case glfw.KeyEscape:
		return api.KeyEscape
	case glfw.KeyEnter:
		return api.KeyEnter
	case glfw.KeySpace:
		return api.KeySpace
	case glfw.KeyTab:
		return api.KeyTab
	case glfw.KeyBackspace:
		return api.KeyBackspace

	// Arrows
	case glfw.KeyUp:
		return api.KeyUp
	case glfw.KeyDown:
		return api.KeyDown
	case glfw.KeyLeft:
		return api.KeyLeft
	case glfw.KeyRight:
		return api.KeyRight

	// Others
	default:
		return api.KeyUnknown
	}
}

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

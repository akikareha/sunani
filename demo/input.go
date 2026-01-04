package main

import (
	"github.com/akikareha/sunani/api"
	"github.com/akikareha/sunani/input"
)

//export sunani_input_key
func InputKey(key uint32, action uint32) {
	k := input.Key(key)
	a := input.Action(action)

	if k == input.KeyQ && a == input.ActionPress {
		api.SystemQuit()
	}
}

//export sunani_input_mouse_motion
func InputMouseMotion(
	x float32,
	y float32,
) {
	mouseX = x
	mouseY = y
}

//export sunani_input_mouse_button
func InputMouseButton(
	button uint32,
	action uint32,
) {
	b := input.Mouse(button)
	a := input.Action(action)

	switch b {
	case input.MouseLeft:
		if a == input.ActionPress {
			anchorEnabled = true
			anchorX = mouseX
			anchorY = mouseY
		}
	case input.MouseRight:
		if a == input.ActionPress {
			anchorEnabled = false
		}
	}
}

//export sunani_input_mouse_wheel
func InputMouseWheel(
	xoff float32,
	yoff float32,
) {
	mouseSize += yoff
	if mouseSize < 1 {
		mouseSize = 1
	} else if mouseSize > 128 {
		mouseSize = 128
	}
}

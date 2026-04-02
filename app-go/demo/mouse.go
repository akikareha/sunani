package main

import (
	"tea.kareha.org/loom/sunani/lib"
)

//export sunani_mouse_motion
func mouseMotion(x, y int32) {
	mouseX = int(x)
	mouseY = int(y)
}

//export sunani_mouse_button
func mouseButton(button uint32, action uint32) {
	b := lib.Mouse(button)
	a := lib.Action(action)

	switch b {
	case lib.MouseLeft:
		if a == lib.ActionPress {
			anchorEnabled = true
			anchorX = mouseX
			anchorY = mouseY
		}
	case lib.MouseRight:
		if a == lib.ActionPress {
			anchorEnabled = false
		}
	}
}

//export sunani_mouse_wheel
func mouseWheel(xoff, yoff int32) {
	mouseSize += int(yoff)
	if mouseSize < 1 {
		mouseSize = 1
	} else if mouseSize > 128 {
		mouseSize = 128
	}
}

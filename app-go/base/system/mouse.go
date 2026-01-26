package system

import (
	"github.com/akikareha/sunani/lib"
)

//export sunani_mouse_init
func mouseInit() {}

var mouseX int32
var mouseY int32

func GetMouse() (int, int) {
	return int(mouseX), int(mouseY)
}

var mouseMotionCallback func(int, int)

func SetMouseMotion(callback func(int, int)) {
	mouseMotionCallback = callback
}

//export sunani_mouse_motion
func mouseMotion(x, y int32) {
	mouseX = x
	mouseY = y

	if mouseMotionCallback != nil {
		mouseMotionCallback(int(x), int(y))
	}
}

var mouseLeft bool
var mouseRight bool
var mouseMiddle bool

func GetButtons() (bool, bool, bool) {
	return mouseLeft, mouseMiddle, mouseRight
}

var mouseButtonCallback func(int, int)

func SetMouseButton(callback func(int, int)) {
	mouseButtonCallback = callback
}

//export sunani_mouse_button
func mouseButton(button uint32, action uint32) {
	b := lib.Mouse(button)
	a := lib.Action(action)

	switch b {
	case lib.MouseLeft:
		if a == lib.ActionPress {
			mouseLeft = true
		} else if a == lib.ActionRelease {
			mouseLeft = false
		}
	case lib.MouseRight:
		if a == lib.ActionPress {
			mouseRight = true
		} else if a == lib.ActionRelease {
			mouseRight = false
		}
	case lib.MouseMiddle:
		if a == lib.ActionPress {
			mouseMiddle = true
		} else if a == lib.ActionRelease {
			mouseMiddle = false
		}
	}

	if mouseButtonCallback != nil {
		mouseButtonCallback(int(button), int(action))
	}
}

var wheelX, wheelY int32

func GetWheel() (int, int) {
	return int(wheelX), int(wheelY)
}

//export sunani_mouse_wheel
func mouseWheel(xoff, yoff int32) {
	wheelX += xoff
	wheelY += yoff
}

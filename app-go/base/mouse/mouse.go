package mouse

import (
	"github.com/akikareha/sunani/lib"
)

//export sunani_mouse_init
func mouseInit() {}

var px, py int

func GetPosition() (int, int) {
	return px, py
}

type MotionHandler func(int, int)

var motionHandlers []MotionHandler = make([]MotionHandler, 0)

func AddMotionHandler(handler MotionHandler) {
	motionHandlers = append(motionHandlers, handler)
}

//export sunani_mouse_motion
func mouseMotion(x, y int32) {
	px = int(x)
	py = int(y)

	for _, handler := range motionHandlers {
		handler(px, py)
	}
}

var left bool
var right bool
var middle bool

func GetButtons() (bool, bool, bool) {
	return left, right, middle
}

type ButtonHandler func(lib.Mouse, lib.Action)

var buttonHandlers []ButtonHandler = make([]ButtonHandler, 0)

func AddButtonHandler(handler ButtonHandler) {
	buttonHandlers = append(buttonHandlers, handler)
}

//export sunani_mouse_button
func mouseButton(button uint32, action uint32) {
	b := lib.Mouse(button)
	a := lib.Action(action)

	switch b {
	case lib.MouseLeft:
		if a == lib.ActionPress {
			left = true
		} else if a == lib.ActionRelease {
			left = false
		}
	case lib.MouseRight:
		if a == lib.ActionPress {
			right = true
		} else if a == lib.ActionRelease {
			right = false
		}
	case lib.MouseMiddle:
		if a == lib.ActionPress {
			middle = true
		} else if a == lib.ActionRelease {
			middle = false
		}
	}

	for _, handler := range buttonHandlers {
		handler(b, a)
	}
}

var wx, wy int

func GetWheel() (int, int) {
	return wx, wy
}

type WheelHandler func(int, int)

var wheelHandlers []WheelHandler = make([]WheelHandler, 0)

//export sunani_mouse_wheel
func mouseWheel(xoff, yoff int32) {
	dx := int(xoff)
	dy := int(yoff)

	wx += dx
	wy += dy

	for _, handler := range wheelHandlers {
		handler(dx, dy)
	}
}

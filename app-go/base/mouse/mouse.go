package mouse

import (
	"github.com/akikareha/sunani/lib"
)

type MotionHandler func(int, int)
type ButtonHandler func(lib.Mouse, lib.Action)
type WheelHandler func(int, int)

var px, py int
var motionHandlers = make([]MotionHandler, 0)

var left bool
var right bool
var middle bool
var buttonHandlers = make([]ButtonHandler, 0)

var wx, wy int
var wheelHandlers = make([]WheelHandler, 0)

//export sunani_mouse_init
func mouseInit() {}

//export sunani_mouse_motion
func mouseMotion(x, y int32) {
	px = int(x)
	py = int(y)

	for _, handler := range motionHandlers {
		handler(px, py)
	}
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

func AddMotionHandler(handler MotionHandler) {
	motionHandlers = append(motionHandlers, handler)
}

func AddButtonHandler(handler ButtonHandler) {
	buttonHandlers = append(buttonHandlers, handler)
}

func AddWheelHandler(handler WheelHandler) {
	wheelHandlers = append(wheelHandlers, handler)
}

func GetPosition() (int, int) {
	return px, py
}

func GetButtons() (bool, bool, bool) {
	return left, right, middle
}

func GetWheel() (int, int) {
	return wx, wy
}

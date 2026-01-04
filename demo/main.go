package main

import (
	"fmt"

	"github.com/akikareha/sunani/api/canvas"
	"github.com/akikareha/sunani/api/fb"
)

var clock uint64

var mouseX float32
var mouseY float32
var mouseBlink int
var mouseSize float32 = 16

var anchorEnabled bool
var anchorX float32
var anchorY float32

//export sunani_system_frame
func frame() {
	clock++

	canvasDraw()
	fbDraw()
}

func canvasDraw() {
	canvas.Begin()
	canvas.Clear(0.10, 0.10, 0.15, 1.0)

	canvas.Color(1, 1, 1, 1)
	canvas.Line(50, 50, 300, 200)

	canvas.Color(0.2, 0.8, 0.4, 1)
	canvas.FillRect(100, 300, 200, 120)

	canvas.Color(1, 0.3, 0.3, 1)
	canvas.Rect(400, 100, 180, 180)

	if anchorEnabled {
		canvas.Color(0, 1, 0, 1)
		canvas.FillRect(anchorX-8, anchorY-8, 16, 16)
	}

	if anchorEnabled {
		canvas.Color(0, 1, 1, 1)
		canvas.Line(anchorX, anchorY, mouseX, mouseY)
	}

	if mouseBlink&0x10 == 0 {
		canvas.Color(1, 0, 0, 1)
	} else {
		canvas.Color(1, 1, 0, 1)
	}
	canvas.FillRect(mouseX-mouseSize/2, mouseY-mouseSize/2, mouseSize, mouseSize)

	mouseBlink++
}

func fbDraw() {
	Clear(0, 0, 0, 0)
	hello := fmt.Sprintf("Hello, Sunani! %d", clock)
	DrawText(16, 32, hello)

	fb.Paint()
}

package main

import (
	"github.com/akikareha/sunani/api"
)

var mouseX float32
var mouseY float32
var mouseBlink int

var anchorEnabled bool
var anchorX float32
var anchorY float32

//export sunani_canvas_draw
func draw() {
	api.CanvasClear(0.10, 0.10, 0.15, 1.0)

	api.CanvasColor(1, 1, 1, 1)
	api.CanvasLine(50, 50, 300, 200)

	api.CanvasColor(0.2, 0.8, 0.4, 1)
	api.CanvasRect(100, 300, 200, 120, 1)

	api.CanvasColor(1, 0.3, 0.3, 1)
	api.CanvasRect(400, 100, 180, 180, 0)

	if anchorEnabled {
		api.CanvasColor(0, 1, 0, 1)
		api.CanvasRect(anchorX-8, anchorY-8, 16, 16, 1)
	}

	if anchorEnabled {
		api.CanvasColor(0, 1, 1, 1)
		api.CanvasLine(anchorX, anchorY, mouseX, mouseY)
	}

	if mouseBlink&0x10 == 0 {
		api.CanvasColor(1, 0, 0, 1)
	} else {
		api.CanvasColor(1, 1, 0, 1)
	}
	api.CanvasRect(mouseX-8, mouseY-8, 16, 16, 1)

	mouseBlink++
}

//export sunani_fb_draw
func FBDraw() {
	hello := "Hello, Sunani!"

	Clear(0, 0, 0, 0)
	for i, b := range []byte(hello) {
		textBuf[i] = b
	}
	DrawText(16, 32, TextBufPtr(), uint32(len(hello)))
}

package main

import (
	"fmt"
	"unsafe"

	"github.com/akikareha/sunani/api/canvas"
	"github.com/akikareha/sunani/api/console"
	"github.com/akikareha/sunani/api/fb"
)

var width, height float32

var clock uint64

var mouseX float32
var mouseY float32
var mouseBlink int
var mouseSize float32 = 16

var anchorEnabled bool
var anchorX float32
var anchorY float32

//export sunani_system_resize
func systemResize(w, h float32) {
	width = w
	height = h
}

//export sunani_system_frame
func systemFrame() {
	clock++

	canvasDraw()
	fbDraw()
}

//export sunani_console_get
func consoleGet(ptr uint32, length uint32) {
	b := unsafe.Slice((*byte)(unsafe.Pointer(uintptr(ptr))), length)
	line := string(b)

	reply := fmt.Sprintf("input: %s\n", line)

	b2 := []byte(reply)
	console.Put(
		uint32(uintptr(unsafe.Pointer(&b2[0]))),
		uint32(len(reply)),
	)
}

func canvasDraw() {
	canvas.Begin()
	canvas.Clear(0.10, 0.10, 0.15, 1.0)

	canvas.Color(0.5, 0.5, 0.5, 1)
	canvas.Rect(8, 8, width-16, height-16)

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

	canvas.Path(mouseX, mouseY)
	canvas.Vertex(mouseX+48, mouseY+24)
	canvas.Vertex(mouseX+24, mouseY+24)
	canvas.Vertex(mouseX+24, mouseY+48)
	canvas.FillPolygon()

	mouseBlink++
}

func fbDraw() {
	fbClear(0, 0, 0, 0)
	hello := fmt.Sprintf("Hello, Sunani! %d", clock)
	drawText(16, 32, hello)

	fb.Paint()
}

package main

import (
	"fmt"
	"unsafe"

	"github.com/akikareha/sunani/app-go/api/canvas"
	"github.com/akikareha/sunani/app-go/api/console"
	"github.com/akikareha/sunani/app-go/api/fb"
	"github.com/akikareha/sunani/app-go/api/runtime"
	"github.com/akikareha/sunani/app-go/base/font"
)

var width, height int32

var clock uint64

var mouseX int32
var mouseY int32
var mouseBlink int
var mouseSize int32 = 16

var anchorEnabled bool
var anchorX int32
var anchorY int32

//export sunani_runtime_resize
func runtimeResize(w, h int32) {
	width = w
	height = h
}

//export sunani_runtime_frame
func runtimeFrame() {
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
	runtime.Clear(16, 16, 24, 255)

	canvas.Begin()

	canvas.Color(127, 127, 127, 255)
	canvas.Rect(8, 8, width-16, height-16)

	canvas.Color(255, 255, 255, 255)
	canvas.Line(50, 50, 300, 200)

	canvas.Color(32, 239, 96, 255)
	canvas.FillRect(100, 300, 200, 120)

	canvas.Color(255, 64, 64, 255)
	canvas.Rect(400, 100, 180, 180)

	if anchorEnabled {
		canvas.Color(0, 255, 0, 255)
		canvas.FillRect(anchorX-8, anchorY-8, 16, 16)
	}

	if anchorEnabled {
		canvas.Color(0, 255, 255, 255)
		canvas.Line(anchorX, anchorY, mouseX, mouseY)
	}

	if mouseBlink&0x10 == 0 {
		canvas.Color(255, 0, 0, 255)
	} else {
		canvas.Color(255, 255, 0, 255)
	}
	canvas.FillRect(mouseX-mouseSize/2, mouseY-mouseSize/2, mouseSize, mouseSize)

	canvas.Path()
	canvas.Vertex(mouseX, mouseY)
	canvas.Vertex(mouseX+48, mouseY+24)
	canvas.Vertex(mouseX+24, mouseY+24)
	canvas.Vertex(mouseX+24, mouseY+48)
	canvas.FillPolygon()

	mouseBlink++

	canvas.Color(255, 255, 255, 255)
	for j := 2; j <= 7; j++ {
		for i := 0; i <= 15; i++ {
			font.Default.Draw(
				16+i*int(mouseSize),
				128+j*int(mouseSize)*2,
				int(mouseSize),
				int(mouseSize)*2,
				rune(j*16+i),
			)
		}
	}
}

func fbDraw() {
	fbClear(0, 0, 0, 0)
	hello := fmt.Sprintf("Hello, Sunani! %d", clock)
	drawText(16, 32, hello)
	size := fmt.Sprintf("Size = %d", mouseSize)
	drawText(16, 48, size)

	fb.Paint()
}

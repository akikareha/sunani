package main

import (
	"fmt"
	"unsafe"

	"github.com/akikareha/sunani/app-go/api/console"
	"github.com/akikareha/sunani/app-go/api/fb"
	"github.com/akikareha/sunani/app-go/api/screen"
	"github.com/akikareha/sunani/app-go/base/canvas"
	"github.com/akikareha/sunani/app-go/base/color"
	"github.com/akikareha/sunani/app-go/base/font"
)

var width, height int

var clock uint64

var mouseX int
var mouseY int
var mouseBlink int
var mouseSize = 32

var anchorEnabled bool
var anchorX int
var anchorY int

var cursor = canvas.Polygon{
	0, 0,
	48, 24,
	24, 24,
	24, 48,
}

//export sunani_screen_resize
func screenResize(w, h int32) {
	width = int(w)
	height = int(h)
}

//export sunani_screen_frame
func screenFrame() {
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
	screen.Clear(16, 16, 24, 255)

	canvas.SetColor(color.New(127, 127, 127, 255))
	canvas.DrawRect(8, 8, width-16, height-16)

	canvas.SetColor(color.New(255, 255, 255, 255))
	canvas.DrawLine(50, 50, 300, 200)

	canvas.SetColor(color.New(32, 239, 96, 255))
	canvas.FillRect(100, 300, 200, 120)

	canvas.SetColor(color.New(255, 64, 64, 255))
	canvas.DrawRect(400, 100, 180, 180)

	if anchorEnabled {
		canvas.SetColor(color.New(0, 255, 0, 255))
		canvas.FillRect(anchorX-8, anchorY-8, 16, 16)
	}

	if anchorEnabled {
		canvas.SetColor(color.New(0, 255, 255, 255))
		canvas.DrawLine(anchorX, anchorY, mouseX, mouseY)
	}

	if mouseBlink&0x10 == 0 {
		canvas.SetColor(color.New(255, 0, 0, 255))
	} else {
		canvas.SetColor(color.New(255, 255, 0, 255))
	}
	canvas.FillRect(
		mouseX-mouseSize/2,
		mouseY-mouseSize/2,
		mouseSize,
		mouseSize,
	)

	denomW, denomH := canvas.GetDenoms()
	cursor.Fill(mouseX, mouseY, denomW, denomH)

	mouseBlink++

	canvas.SetColor(color.New(255, 255, 255, 255))
	for j := 2; j <= 7; j++ {
		for i := 0; i <= 15; i++ {
			font.ASCII.Draw(
				16+i*int(mouseSize),
				128+j*int(mouseSize)*2,
				int(mouseSize),
				int(mouseSize),
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

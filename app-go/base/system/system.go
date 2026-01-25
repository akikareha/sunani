package system

import (
	"fmt"
	"unsafe"

	"github.com/akikareha/sunani/app-go/api/canvas"
	"github.com/akikareha/sunani/app-go/api/console"
	"github.com/akikareha/sunani/app-go/api/fb"
	"github.com/akikareha/sunani/app-go/api/runtime"
	"github.com/akikareha/sunani/app-go/api/std"
	"github.com/akikareha/sunani/app-go/base/font"
	"github.com/akikareha/sunani/app-go/base/key"
	"github.com/akikareha/sunani/lib"
)

func Run() {
	// do nothing
}

//
// Runtime
//

const title = "System"
const hello = "Sunani System v0.0.0\n"

//export sunani_runtime_init
func runtimeInit() {
	b := []byte(title)
	runtime.Title(
		uint32(uintptr(unsafe.Pointer(&b[0]))),
		uint32(len(title)),
	)

	runtime.Cursor(0)
}

var width, height int32

//export sunani_runtime_resize
func runtimeResize(w, h int32) {
	width = w
	height = h
}

var clock uint64
var now int64
var prev int64
var fps float64

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

//export sunani_runtime_frame
func runtimeFrame() {
	clock++
	now = std.Now()
	if clock%16 == 0 {
		elapsed := now - prev
		fps = 16 * 1000 / float64(elapsed)
		prev = now
	}

	runtime.Clear(0, 0, 0, 255)
	canvas.Begin()

	pitch := min(width/15, height/5)
	ox := (width - pitch*15) / 2
	oy := height - pitch*5
	key.ShowKeyboard(ox, oy, pitch)

	showMouse()
	showInfo()
}

//
// Console
//

const consoleBufferLength = 4096

var consoleBuffer = make([]byte, consoleBufferLength)

//export sunani_console_init
func consoleInit() {
	console.Params(
		uint32(uintptr(unsafe.Pointer(&consoleBuffer[0]))),
		uint32(consoleBufferLength),
	)

	b := []byte(hello)
	console.Put(
		uint32(uintptr(unsafe.Pointer(&b[0]))),
		uint32(len(hello)),
	)
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

//
// Canvas
//

//export sunani_canvas_init
func canvasInit() {}

//
// Framebuffer
//

const fbWidth int32 = 256
const fbHeight int32 = 256

var framebuffer = make([]byte, fbWidth*fbHeight*4)

//export sunani_fb_init
func fbInit() {
	fb.Params(
		uint32(uintptr(unsafe.Pointer(&framebuffer[0]))),
		fbWidth,
		fbHeight,
	)
}

//
// Keyboard
//

//export sunani_key_init
func keyInit() {}

//
// Mouse
//

//export sunani_mouse_init
func mouseInit() {}

var mouseX int32
var mouseY int32

//export sunani_mouse_motion
func mouseMotion(x, y int32) {
	mouseX = x
	mouseY = y
}

var mouseLeft bool
var mouseRight bool
var mouseMiddle bool

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
}

var wheelX, wheelY int32

//export sunani_mouse_wheel
func mouseWheel(xoff, yoff int32) {
	wheelX += xoff
	wheelY += yoff
}

//
// System
//

var mouseSignX int32 = 1
var mouseSignY int32 = 1

func showMouse() {
	if mouseX < 64 {
		mouseSignX = 1
	} else if mouseX >= width - 64 {
		mouseSignX = -1
	}
	if mouseY < 64 {
		mouseSignY = 1
	} else if mouseY >= height - 64 {
		mouseSignY = -1
	}

	canvas.Color(0, 0, 255, 128)
	canvas.Path()
	canvas.Vertex(mouseX, mouseY)
	canvas.Vertex(mouseX+48*mouseSignX, mouseY+24*mouseSignY)
	canvas.Vertex(mouseX+24*mouseSignX, mouseY+24*mouseSignY)
	canvas.Vertex(mouseX+24*mouseSignX, mouseY+48*mouseSignY)
	canvas.Polygon()

	canvas.Color(255, 255, 0, 128)
	canvas.Path()
	canvas.Vertex(mouseX, mouseY)
	canvas.Vertex(mouseX+48*mouseSignX, mouseY+24*mouseSignY)
	canvas.Vertex(mouseX+24*mouseSignX, mouseY+24*mouseSignY)
	canvas.Vertex(mouseX+24*mouseSignX, mouseY+48*mouseSignY)
	canvas.FillPolygon()
}

func Print(x, y, sx, sy int32, s string) {
	i := 0
	for _, r := range s {
		font.DrawRune(x+int32(i)*sx, y, sx, sy, r)
		i++
	}
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func showInfo() {
	canvas.Color(255, 255, 255, 128)
	Print(0, 0, 8, 16, fmt.Sprintf("Size=%dx%d Mouse=%d,%d Button=%d,%d,%d Wheel=%d,%d", width, height, mouseX, mouseY, btoi(mouseLeft), btoi(mouseMiddle), btoi(mouseRight), wheelX, wheelY))
	Print(0, 16, 8, 16, fmt.Sprintf("Now=%d Clock=%d FPS=%3.2f", now, clock, fps))
}

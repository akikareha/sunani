package system

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/akikareha/sunani/app-go/api/canvas"
	"github.com/akikareha/sunani/app-go/api/fb"
	"github.com/akikareha/sunani/app-go/base/font"
)

func Run() {
	// do nothing
}

const dummyString = "Dummy"

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
// Start
//

var startCallback func()

func SetStart(callback func()) {
	startCallback = callback
}

//export sunani_start
func start() {
	if startCallback != nil {
		startCallback()
	}
}

//
// System
//

var mouseSignX int32 = 1
var mouseSignY int32 = 1

func showMouse() {
	if mouseX < 64 {
		mouseSignX = 1
	} else if mouseX >= width-64 {
		mouseSignX = -1
	}
	if mouseY < 64 {
		mouseSignY = 1
	} else if mouseY >= height-64 {
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

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func showInfo() {
	canvas.Color(255, 255, 255, 128)
	font.ASCII.DrawString(0, 0, 16, 16, fmt.Sprintf("Size=%dx%d Mouse=%d,%d Button=%d,%d,%d Wheel=%d,%d", width, height, mouseX, mouseY, btoi(mouseLeft), btoi(mouseMiddle), btoi(mouseRight), wheelX, wheelY))
	font.ASCII.DrawString(0, 16, 16, 16, fmt.Sprintf("Now=%d Clock=%d FPS=%3.2f", now, clock, fps))
}

var consoleLines []string = []string{""}

func showConsole(ox, oy int, sx, sy int) {
	canvas.Color(255, 255, 255, 128)
	n := len(consoleLines)
	for i := 0; i < n; i++ {
		font.ASCII.DrawString(ox, oy+(-n+i)*sy, sx, sy, consoleLines[i])
	}
}

func addConsoleLine(s string) {
	lines := strings.Split(s, "\n")
	if len(lines) < 1 {
		return
	}
	consoleLines[len(consoleLines)-1] += lines[0]
	consoleLines = append(consoleLines, lines[1:]...)

	if len(consoleLines) > 8 {
		consoleLines = consoleLines[len(consoleLines)-8:]
	}
}

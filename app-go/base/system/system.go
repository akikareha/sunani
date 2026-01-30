package system

import (
	"fmt"
	"strings"

	"github.com/akikareha/sunani/app-go/base/canvas"
	"github.com/akikareha/sunani/app-go/base/console"
	"github.com/akikareha/sunani/app-go/base/font"
	"github.com/akikareha/sunani/app-go/base/kb"
	"github.com/akikareha/sunani/app-go/base/key"
	"github.com/akikareha/sunani/app-go/base/mouse"
	"github.com/akikareha/sunani/app-go/base/screen"
	"github.com/akikareha/sunani/lib"
)

func systemPrint(s string) {
	console.Print(s)
	addConsoleLine(s)
}

func consoleHandler(line string) {
	addConsoleLine(line + "\n")
	repl(line)
	systemPrint(prompt)
}

func keyHandler(key lib.Key, action lib.Action) {
	if action == lib.ActionPress {
		addConsoleLine(kb.Char(key))
	}
}

func mouseButtonHandler(button lib.Mouse, action lib.Action) {
	if button == lib.MouseLeft && action == lib.ActionPress {
		if virtualKey != lib.KeyUnknown {
			addConsoleLine(kb.Char(virtualKey))
		}
	}
}

func Init() {
	console.AddHandler(consoleHandler)
	screen.AddFrameHandler(frameHandler)
	key.AddHandler(keyHandler)
	mouse.AddButtonHandler(mouseButtonHandler)
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

	systemPrint(greeting)
	systemPrint(prompt)
}

//
// System
//

var mouseSignX int = 1
var mouseSignY int = 1

func showMouse() {
	width, height := screen.GetSize()
	mouseX, mouseY := mouse.GetPosition()

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
	now := screen.Now()
	width, height := screen.GetSize()
	clock := screen.Clock()
	fps := screen.FPS()
	mouseX, mouseY := mouse.GetPosition()
	mouseLeft, mouseRight, mouseMiddle := mouse.GetButtons()
	wheelX, wheelY := mouse.GetWheel()

	canvas.Color(255, 255, 255, 128)
	font.ASCII.DrawString(0, 0, 16, 16, fmt.Sprintf("Size=%dx%d Mouse=%d,%d Button=%d,%d,%d Wheel=%d,%d", width, height, mouseX, mouseY, btoi(mouseLeft), btoi(mouseRight), btoi(mouseMiddle), wheelX, wheelY))
	font.ASCII.DrawString(0, 16, 16, 16, fmt.Sprintf("Now=%d Clock=%d FPS=%d", now, clock, fps))
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

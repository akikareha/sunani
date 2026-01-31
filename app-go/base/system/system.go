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

func consoleInputHandler(line string) {
	addConsoleLine(line + "\n")
	repl(line)
	systemPrint(prompt)
}

func keyEventHandler(key lib.Key, action lib.Action) {
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
	console.AddInputHandler(consoleInputHandler)
	screen.AddFrameHandler(frameHandler)
	key.AddEventHandler(keyEventHandler)
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

var mouseSignX = 1
var mouseSignY = 1

var cursor = canvas.Polygon{
	0, 0,
	48, 24,
	24, 24,
	24, 48,
}

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

	denomW, denomH := canvas.GetDenoms()

	canvas.SetColor(0, 0, 255, 128)
	cursor.Draw(mouseX, mouseY, denomW*mouseSignX, denomH*mouseSignY)

	canvas.SetColor(255, 255, 0, 128)
	cursor.Fill(mouseX, mouseY, denomW*mouseSignX, denomH*mouseSignY)
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

	canvas.SetColor(255, 255, 255, 128)
	font.ASCII.DrawString(0, 0, 16, 16, fmt.Sprintf("Size=%dx%d Mouse=%d,%d Button=%d,%d,%d Wheel=%d,%d", width, height, mouseX, mouseY, btoi(mouseLeft), btoi(mouseRight), btoi(mouseMiddle), wheelX, wheelY))
	font.ASCII.DrawString(0, 16, 16, 16, fmt.Sprintf("Now=%d Clock=%d FPS=%d", now, clock, fps))
}

var consoleLines = []string{""}

func showConsole(ox, oy int, sx, sy int) {
	canvas.SetColor(255, 255, 255, 128)
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

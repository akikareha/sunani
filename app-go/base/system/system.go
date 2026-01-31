package system

import (
	"fmt"

	"github.com/akikareha/sunani/app-go/base/canvas"
	"github.com/akikareha/sunani/app-go/base/canvas/widget"
	"github.com/akikareha/sunani/app-go/base/fb"
	"github.com/akikareha/sunani/app-go/base/font"
	"github.com/akikareha/sunani/app-go/base/kb"
	"github.com/akikareha/sunani/app-go/base/key"
	"github.com/akikareha/sunani/app-go/base/mouse"
	"github.com/akikareha/sunani/app-go/base/repl"
	"github.com/akikareha/sunani/app-go/base/screen"
	"github.com/akikareha/sunani/lib"
)

var text = widget.NewText()
var input = widget.NewText()

var startCallback func()
var infoVisible bool

var virtualKey = lib.KeyUnknown

var mouseSignX = 1
var mouseSignY = 1

func Init() {
	screen.AddFrameHandler(frameHandler)
	key.AddEventHandler(keyEventHandler)
	mouse.AddButtonHandler(mouseButtonHandler)
	repl.Default.AddEchoHandler(echoHandler)

	text.SetColor(255, 255, 255, 128)
	input.SetColor(255, 255, 0, 192)
}

//export sunani_start
func start() {
	if startCallback != nil {
		startCallback()
	}

	repl.Default.Init()
}

func SetStart(callback func()) {
	startCallback = callback
}

func SetInfoVisible(v bool) {
	infoVisible = v
}

func frameHandler() {
	width, height := screen.GetSize()
	pitch := min(width/15, height/5)
	ox := (width - pitch*15) / 2
	oy := height - pitch*5
	mouseX, mouseY := mouse.GetPosition()
	virtualKey = kb.Default.Draw(
		int(ox), int(oy), int(pitch),
		int(mouseX), int(mouseY),
	)

	text.Draw(int(ox), int(oy), int(pitch)/2, int(pitch)/2)
	input.Draw(int(ox+pitch/2), int(oy), int(pitch)/2, int(pitch)/2)

	fb.Paint()

	showMouse()
	if infoVisible {
		showInfo()
	}
}

func putKey(key lib.Key) {
	char := kb.Char(key)
	if char == "\n" {
		line := input.Get()
		input.Clear()
		repl.Default.Input(line)
	} else {
		input.Add(char)
	}
}

func keyEventHandler(key lib.Key, action lib.Action) {
	if action == lib.ActionPress {
		putKey(key)
	}
}

func mouseButtonHandler(button lib.Mouse, action lib.Action) {
	if button == lib.MouseLeft && action == lib.ActionPress {
		if virtualKey != lib.KeyUnknown {
			putKey(virtualKey)
		}
	}
}

func echoHandler(r *repl.REPL, s string) {
	text.Add(s)
}

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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

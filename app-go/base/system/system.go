package system

import (
	"fmt"
	"unicode/utf8"

	"github.com/akikareha/sunani/app-go/base/canvas"
	"github.com/akikareha/sunani/app-go/base/canvas/widget"
	"github.com/akikareha/sunani/app-go/base/fb"
	"github.com/akikareha/sunani/app-go/base/font"
	"github.com/akikareha/sunani/app-go/base/kb"
	"github.com/akikareha/sunani/app-go/base/key"
	"github.com/akikareha/sunani/app-go/base/mouse"
	"github.com/akikareha/sunani/app-go/base/screen"
	"github.com/akikareha/sunani/lib"
)

var cursor = widget.NewCursor()
var input = widget.NewText()
var text = widget.NewText()

var startCallback func()
var infoVisible bool

var virtualKey = lib.KeyUnknown

func Init() {
	screen.AddFrameHandler(frameHandler)
	key.AddEventHandler(keyEventHandler)
	mouse.AddButtonHandler(mouseButtonHandler)

	text.SetColor(255, 255, 255, 128)
	input.SetColor(255, 255, 0, 192)
}

//export sunani_start
func start() {
	if startCallback != nil {
		startCallback()
	}

	initREPL()
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

	cursor.Draw()

	if infoVisible {
		showInfo()
	}
}

func putKey(key lib.Key) {
	if key == lib.KeyEnter {
		line := input.Get()
		input.Clear()
		REPL.Input(line)
	} else if key == lib.KeyBackspace {
		line := input.Get()
		input.Clear()
		_, size := utf8.DecodeLastRuneInString(line)
		input.Add(line[:len(line)-size])
	} else {
		char := kb.Char(key)
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

package system

import (
	"github.com/akikareha/sunani/app-go/base/fb"
	"github.com/akikareha/sunani/app-go/base/kb"
	"github.com/akikareha/sunani/app-go/base/mouse"
	"github.com/akikareha/sunani/app-go/base/screen"
	"github.com/akikareha/sunani/lib"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var virtualKey lib.Key = lib.KeyUnknown

var infoVisible bool

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

	showConsole(int(ox), int(oy), int(pitch)/2, int(pitch)/2)

	fb.Paint()

	showMouse()
	if infoVisible {
		showInfo()
	}
}

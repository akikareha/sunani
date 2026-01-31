package main

import (
	"github.com/akikareha/sunani/app-go/base/fb"
	"github.com/akikareha/sunani/app-go/base/mouse"
	"github.com/akikareha/sunani/app-go/base/screen"
	"github.com/akikareha/sunani/app-go/base/system"
)

const fbW, fbH = 256, 256

//export sunani_init
func sunaniInit() {
	system.Init()
	system.SetInfoVisible(true)

	screen.AddFrameHandler(frameHandler)
	fb.SetSize(fbW, fbH)
	fb.SetColor(255, 128, 64, 192)
}

func frameHandler() {
	x, y := mouse.GetPosition()
	ox, oy, rw, rh := fb.GetRect()
	if rw > 0 && rh > 0 {
		px := (x - ox) * fbW / rw
		py := (y - oy) * fbH / rh
		fb.DrawPixel(px, py)
	}
}

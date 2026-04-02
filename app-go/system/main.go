package main

import (
	"tea.kareha.org/loom/sunani/app-go/base/color"
	"tea.kareha.org/loom/sunani/app-go/base/fb"
	"tea.kareha.org/loom/sunani/app-go/base/mouse"
	"tea.kareha.org/loom/sunani/app-go/base/screen"
	"tea.kareha.org/loom/sunani/app-go/base/system"
)

const fbW, fbH = 256, 256

//export sunani_init
func sunaniInit() {
	system.Init()
	system.SetInfoVisible(true)

	screen.AddFrameHandler(frameHandler)
	fb.SetSize(fbW, fbH)
	fb.SetColor(color.New(255, 0, 0, 128))
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

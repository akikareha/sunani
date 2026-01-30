package main

import (
	"github.com/akikareha/sunani/app-go/base/fb"
	"github.com/akikareha/sunani/app-go/base/mouse"
	"github.com/akikareha/sunani/app-go/base/screen"
	"github.com/akikareha/sunani/app-go/base/system"
)

//export sunani_init
func sunaniInit() {
	system.Init()
	system.SetInfoVisible(true)

	screen.AddFrameHandler(frameHandler)
	fb.SetColor(255, 128, 64, 192)
}

func frameHandler() {
	x, y := mouse.GetPosition()
	fb.DrawPixel(x, y)
}

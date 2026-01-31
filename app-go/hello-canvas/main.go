package main

import (
	"github.com/akikareha/sunani/app-go/api/screen"
	"github.com/akikareha/sunani/app-go/base/canvas"
	"github.com/akikareha/sunani/app-go/base/font"
)

const x, y = 8, 8
const w, h = 32, 32
const message = "Hello, World!\n"

//export sunani_screen_init
func screenInit() {}

//export sunani_screen_frame
func screenFrame() {
	screen.Clear(0, 0, 0, 255)
	canvas.SetColor(255, 255, 255, 255)
	font.ASCII.DrawString(x, y, w, h, message)
}

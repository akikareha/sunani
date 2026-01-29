package main

import (
	"github.com/akikareha/sunani/app-go/api/canvas"
	"github.com/akikareha/sunani/app-go/api/runtime"
	"github.com/akikareha/sunani/app-go/base/font"
)

const x, y = 8, 8
const w, h = 32, 32
const message = "Hello, World!\n"

//export sunani_runtime_init
func runtimeInit() {}

//export sunani_canvas_init
func canvasInit() {}

//export sunani_runtime_frame
func runtimeFrame() {
	runtime.Clear(0, 0, 0, 255)
	canvas.Begin()
	canvas.Color(255, 255, 255, 255)
	font.ASCII.DrawString(x, y, w, h, message)
}

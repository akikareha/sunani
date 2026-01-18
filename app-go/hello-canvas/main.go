package main

import (
	"github.com/akikareha/sunani/app-go/api/canvas"
	"github.com/akikareha/sunani/app-go/base/font"
)

const title = "Hello Canvas"
const hello = "Hello, World!\n"
const size = 16

//export sunani_system_init
func systemInit() {}

//export sunani_canvas_init
func canvasInit() {}

//export sunani_system_frame
func systemFrame() {
	canvas.Begin()
	canvas.Clear(0, 0, 0, 255)

	canvas.Color(255, 255, 255, 255)
	i := 0
	for _, r := range hello {
		font.DrawRune(
			uint32(i)*size,
			0,
			size,
			size*2,
			r,
		)
		i++
	}
}

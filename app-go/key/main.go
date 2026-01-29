package main

import (
	"github.com/akikareha/sunani/app-go/api/canvas"
	"github.com/akikareha/sunani/app-go/api/runtime"
	"github.com/akikareha/sunani/app-go/base/font"
	"github.com/akikareha/sunani/lib"
)

const size = 16

//export sunani_runtime_init
func runtimeInit() {}

//export sunani_canvas_init
func canvasInit() {}

type Key struct {
	Code  lib.Key
	Label string
	// grid units
	X      int
	Y      int
	Width  int
	Height int
}

type Keyboard struct {
	Keys       []Key
	GridWidth  int
	GridHeight int
}

var RK61Keys = []Key{
	// row 1
	{lib.KeyEscape, "Esc", 0, 0, 12, 12},
	{lib.Key1, "1", 12, 0, 12, 12},
	{lib.Key2, "2", 24, 0, 12, 12},
	{lib.Key3, "3", 36, 0, 12, 12},
	{lib.Key4, "4", 48, 0, 12, 12},
	{lib.Key5, "5", 60, 0, 12, 12},
	{lib.Key6, "6", 72, 0, 12, 12},
	{lib.Key7, "7", 84, 0, 12, 12},
	{lib.Key8, "8", 96, 0, 12, 12},
	{lib.Key9, "9", 108, 0, 12, 12},
	{lib.Key0, "0", 120, 0, 12, 12},
	{lib.KeyUnknown, "-", 132, 0, 12, 12},
	{lib.KeyUnknown, "=", 144, 0, 12, 12},
	{lib.KeyBackspace, "BS", 156, 0, 24, 12},
	// row 2
	{lib.KeyTab, "Tab", 0, 12, 18, 12},
	{lib.KeyQ, "Q", 18, 12, 12, 12},
	{lib.KeyW, "W", 30, 12, 12, 12},
	{lib.KeyE, "E", 42, 12, 12, 12},
	{lib.KeyR, "R", 54, 12, 12, 12},
	{lib.KeyT, "T", 66, 12, 12, 12},
	{lib.KeyY, "Y", 78, 12, 12, 12},
	{lib.KeyU, "U", 90, 12, 12, 12},
	{lib.KeyI, "I", 102, 12, 12, 12},
	{lib.KeyO, "O", 114, 12, 12, 12},
	{lib.KeyP, "P", 126, 12, 12, 12},
	{lib.KeyUnknown, "[", 138, 12, 12, 12},
	{lib.KeyUnknown, "]", 150, 12, 12, 12},
	{lib.KeyUnknown, "\\", 162, 12, 18, 12},
	// row 3
	{lib.KeyUnknown, "Caps", 0, 24, 22, 12},
	{lib.KeyA, "A", 22, 24, 12, 12},
	{lib.KeyS, "S", 34, 24, 12, 12},
	{lib.KeyD, "D", 46, 24, 12, 12},
	{lib.KeyF, "F", 58, 24, 12, 12},
	{lib.KeyG, "G", 70, 24, 12, 12},
	{lib.KeyH, "H", 82, 24, 12, 12},
	{lib.KeyJ, "J", 94, 24, 12, 12},
	{lib.KeyK, "K", 106, 24, 12, 12},
	{lib.KeyL, "L", 118, 24, 12, 12},
	{lib.KeyUnknown, ";", 130, 24, 12, 12},
	{lib.KeyUnknown, "'", 142, 24, 12, 12},
	{lib.KeyEnter, "Enter", 154, 24, 26, 12},
	// row 4
	{lib.KeyUnknown, "Shift", 0, 36, 28, 12},
	{lib.KeyZ, "Z", 28, 36, 12, 12},
	{lib.KeyX, "X", 40, 36, 12, 12},
	{lib.KeyC, "C", 52, 36, 12, 12},
	{lib.KeyV, "V", 64, 36, 12, 12},
	{lib.KeyB, "B", 76, 36, 12, 12},
	{lib.KeyN, "N", 88, 36, 12, 12},
	{lib.KeyM, "M", 100, 36, 12, 12},
	{lib.KeyUnknown, ",", 112, 36, 12, 12},
	{lib.KeyUnknown, ".", 124, 36, 12, 12},
	{lib.KeyUnknown, "/", 136, 36, 12, 12},
	{lib.KeyUnknown, "Shift", 148, 36, 32, 12},
	// row 5
	{lib.KeyUnknown, "Ctrl", 0, 48, 15, 12},
	{lib.KeyUnknown, "OS", 15, 48, 15, 12},
	{lib.KeyUnknown, "Alt", 30, 48, 15, 12},
	{lib.KeySpace, "Space", 45, 48, 74, 12},
	{lib.KeyUnknown, "Alt", 119, 48, 15, 12},
	{lib.KeyUnknown, "Menu", 134, 48, 15, 12},
	{lib.KeyUnknown, "Ctrl", 149, 48, 15, 12},
	{lib.KeyUnknown, "Fn", 164, 48, 16, 12},
}

var RK61 = Keyboard{
	Keys:       RK61Keys,
	GridWidth:  12,
	GridHeight: 12,
}

//export sunani_runtime_frame
func runtimeFrame() {
	runtime.Clear(0, 0, 0, 255)

	canvas.Begin()

	kb := RK61
	ox, oy := 0, 0
	pitch := 32

	canvas.Color(64, 64, 64, 255)
	for i := 0; i < len(kb.Keys); i++ {
		k := kb.Keys[i]
		canvas.FillRect(
			int32(ox+k.X*pitch/kb.GridWidth+pitch/8),
			int32(oy+k.Y*pitch/kb.GridHeight+pitch/8),
			int32(k.Width*pitch/kb.GridWidth-pitch/4),
			int32(k.Height*pitch/kb.GridHeight-pitch/4),
		)
	}

	canvas.Color(255, 255, 255, 255)
	for i := 0; i < len(kb.Keys); i++ {
		k := kb.Keys[i]
		n := len(k.Label)
		width := pitch * k.Width / kb.GridWidth / n
		height := pitch * k.Height / kb.GridHeight / 2
		dx := (pitch*k.Width/kb.GridWidth - width*n/2) / 2
		dy := (pitch*k.Height/kb.GridHeight - height) / 2
		font.ASCII.DrawString(
			ox+k.X*pitch/kb.GridWidth+dx,
			oy+k.Y*pitch/kb.GridHeight+dy,
			width,
			height,
			k.Label,
		)
	}
}

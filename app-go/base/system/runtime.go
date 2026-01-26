package system

import (
	"unsafe"

	"github.com/akikareha/sunani/app-go/api/canvas"
	"github.com/akikareha/sunani/app-go/api/runtime"
	"github.com/akikareha/sunani/app-go/api/std"
	"github.com/akikareha/sunani/app-go/base/key"
	"github.com/akikareha/sunani/lib"
)

var runtimeInitialized bool

var title string = "Sunani System"

func GetTitle() string {
	return title
}

func SetTitle(s string) {
	title = s

	if !runtimeInitialized {
		return
	}

	b := []byte(title)
	length := len(b)
	if length < 1 {
		// cannot point null string
		b = []byte(dummyString)
	}
	runtime.Title(
		uint32(uintptr(unsafe.Pointer(&b[0]))),
		uint32(length),
	)
}

var cursorVisible bool

func IsCursorVisible() bool {
	return cursorVisible
}

func SetCursorVisible(v bool) {
	cursorVisible = v

	if !runtimeInitialized {
		return
	}

	if v {
		runtime.Cursor(1)
	} else {
		runtime.Cursor(0)
	}
}

//export sunani_runtime_init
func runtimeInit() {
	runtimeInitialized = true

	SetTitle(title)
	SetCursorVisible(cursorVisible)
}

var width, height int32

//export sunani_runtime_resize
func runtimeResize(w, h int32) {
	width = w
	height = h
}

func GetSize() (int, int) {
	return int(width), int(height)
}

var clock uint64
var now int64
var prev int64
var fps float64

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

var bgR, bgG, bgB, bgA int = 0, 0, 0, 255

func GetBg() (int, int, int, int) {
	return bgR, bgG, bgB, bgA
}

func SetBg(r, g, b, a int) {
	bgR, bgG, bgB, bgA = r, g, b, a
}

var virtualKey lib.Key = lib.KeyUnknown

//export sunani_runtime_frame
func runtimeFrame() {
	clock++
	now = std.Now()
	if clock%16 == 0 {
		elapsed := now - prev
		fps = 16 * 1000 / float64(elapsed)
		prev = now
	}

	runtime.Clear(uint32(bgR), uint32(bgG), uint32(bgB), uint32(bgA))
	canvas.Begin()

	pitch := min(width/15, height/5)
	ox := (width - pitch*15) / 2
	oy := height - pitch*5
	virtualKey = key.Default.Draw(
		int(ox), int(oy), int(pitch), keyTable,
		int(mouseX), int(mouseY),
	)

	showConsole(int(ox), int(oy), int(pitch)/4, int(pitch)/2)

	showMouse()
	showInfo()
}

func Halt() {
	runtime.Halt()
}

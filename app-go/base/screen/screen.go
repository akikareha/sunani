package screen

import (
	"unsafe"

	scr "tea.kareha.org/loom/sunani/app-go/api/screen"
	"tea.kareha.org/loom/sunani/app-go/base/color"
	"tea.kareha.org/loom/sunani/app-go/base/std"
)

type ResizeHandler func(int, int)
type FrameHandler func()

const dummy = " "

var initialized bool
var title string
var cursorVisible bool

var width, height int
var resizeHandlers = make([]ResizeHandler, 0)

var clock int64
var now int64
var tick int64 = 16
var prev int64
var fps int64
var bg = color.New(0, 0, 0, 255)
var frameHandlers = make([]FrameHandler, 0)

//export sunani_screen_init
func screenInit() {
	initialized = true

	SetTitle(title)
	SetCursorVisible(cursorVisible)
}

//export sunani_screen_resize
func screenResize(w, h int32) {
	width = int(w)
	height = int(h)

	for _, handler := range resizeHandlers {
		handler(width, height)
	}
}

//export sunani_screen_frame
func screenFrame() {
	clock++
	now = std.Now()
	if clock%tick == 0 {
		elapsed := now - prev
		fps = tick * 1000 / elapsed
		prev = now
	}

	Clear()

	for _, handler := range frameHandlers {
		handler()
	}
}

func AddFrameHandler(handler FrameHandler) {
	frameHandlers = append(frameHandlers, handler)
}

func GetTitle() string {
	return title
}

func SetTitle(s string) {
	title = s

	if !initialized {
		return
	}

	b := []byte(title)
	length := len(b)
	if length < 1 {
		// cannot point null string
		b = []byte(dummy)
	}
	scr.Title(
		uint32(uintptr(unsafe.Pointer(&b[0]))),
		uint32(length),
	)
}

func IsCursorVisible() bool {
	return cursorVisible
}

func SetCursorVisible(visible bool) {
	cursorVisible = visible

	if !initialized {
		return
	}

	if cursorVisible {
		scr.Cursor(1)
	} else {
		scr.Cursor(0)
	}
}

func GetSize() (int, int) {
	return width, height
}

func Clock() int64 {
	return clock
}

func Now() int64 {
	return now
}

func GetTick() int {
	return int(tick)
}

func SetTick(value int) {
	tick = int64(value)
}

func FPS() int {
	return int(fps)
}

func GetBg() color.Color {
	return bg
}

func SetBg(c color.Color) {
	bg = c
}

func Clear() {
	r, g, b, a := bg.Values()
	scr.Clear(uint32(r), uint32(g), uint32(b), uint32(a))
}

func Halt() {
	scr.Halt()
}

package fb

import (
	"unsafe"

	api "tea.kareha.org/loom/sunani/app-go/api/fb"
)

type RectHandler func(int, int, int, int)

var initialized bool
var w, h = 256, 256
var buffer []byte

var rx, ry int
var rw, rh int
var rectHandlers = make([]RectHandler, 0)

//export sunani_fb_init
func fbInit() {
	initialized = true

	SetSize(w, h)
}

//export sunani_fb_rect
func fbRect(x, y int32, w, h int32) {
	rx = int(x)
	ry = int(y)
	rw = int(w)
	rh = int(h)

	for _, handler := range rectHandlers {
		handler(rx, ry, rw, rh)
	}
}

func AddRectHandler(handler RectHandler) {
	rectHandlers = append(rectHandlers, handler)
}

func SetSize(width, height int) {
	w = width
	h = height
	buffer = make([]byte, w*h*4)

	if !initialized {
		return
	}
	api.Params(
		uint32(uintptr(unsafe.Pointer(&buffer[0]))),
		int32(w),
		int32(h),
	)
}

func GetRect() (int, int, int, int) {
	return rx, ry, rw, rh
}

func Paint() {
	api.Paint()
}

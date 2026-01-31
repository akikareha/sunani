package fb

import (
	"unsafe"

	fbuf "github.com/akikareha/sunani/app-go/api/fb"
)

var initialized bool
var w, h int = 256, 256
var buffer []byte
var ox, oy int
var rw, rh int

func SetSize(width, height int) {
	w = width
	h = height
	buffer = make([]byte, w*h*4)

	if !initialized {
		return
	}
	fbuf.Params(
		uint32(uintptr(unsafe.Pointer(&buffer[0]))),
		int32(w),
		int32(h),
	)
}

func Clear(r, g, b, a int) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			i := (y*w + x) * 4
			buffer[i] = uint8(r)
			buffer[i+1] = uint8(g)
			buffer[i+2] = uint8(b)
			buffer[i+3] = uint8(a)
		}
	}
}

//export sunani_fb_init
func fbInit() {
	initialized = true

	SetSize(w, h)
}

//export sunani_fb_rect
func fbRect(x, y int32, w, h int32) {
	ox = int(x)
	oy = int(y)
	rw = int(w)
	rh = int(h)
}

func GetRect() (int, int, int, int) {
	return ox, oy, rw, rh
}

func Paint() {
	fbuf.Paint()
}

var fgR, fgG, fgB, fgA int

func SetColor(r, g, b, a int) {
	fgR, fgG, fgB, fgA = r, g, b, a
}

func DrawPixel(x, y int) {
	if x < 0 || x >= w {
		return
	}
	if y < 0 || y >= h {
		return
	}
	i := (y*w + x) * 4
	buffer[i] = uint8(fgR)
	buffer[i+1] = uint8(fgG)
	buffer[i+2] = uint8(fgB)
	buffer[i+3] = uint8(fgA)
}

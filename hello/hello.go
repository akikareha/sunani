package main

import (
	"unsafe"

	"github.com/akikareha/sunani/hello/api"
)

var mouseEnabled bool
var mouseX float32
var mouseY float32
var mouseDX float32
var mouseDY float32
var mouseBlink int

var anchorEnabled bool
var anchorX float32
var anchorY float32

//go:wasmimport sunani system_quit
func systemQuit()

//go:wasmimport sunani canvas_clear
func canvasClear(r, g, b, a float32)

//go:wasmimport sunani canvas_color
func canvasSetColor(r, g, b, a float32)

//go:wasmimport sunani canvas_line
func canvasLine(x1, y1, x2, y2 float32)

//go:wasmimport sunani canvas_rect
func canvasRect(x, y, w, h float32, fill uint32)

//export sunani_canvas_draw
func draw() {
	canvasClear(0.10, 0.10, 0.15, 1.0)

	canvasSetColor(1, 1, 1, 1)
	canvasLine(50, 50, 300, 200)

	canvasSetColor(0.2, 0.8, 0.4, 1)
	canvasRect(100, 300, 200, 120, 1)

	canvasSetColor(1, 0.3, 0.3, 1)
	canvasRect(400, 100, 180, 180, 0)

	if anchorEnabled {
		canvasSetColor(0, 1, 0, 1)
		canvasRect(anchorX-8, anchorY-8, 16, 16, 1)
	}

	if mouseEnabled {
		if anchorEnabled {
			canvasSetColor(0, 1, 1, 1)
			canvasLine(anchorX, anchorY, mouseX, mouseY)
		}

		canvasSetColor(0.5, 0.5, 0.5, 1)
		canvasRect(mouseX-8-mouseDX, mouseY-8-mouseDY, 16, 16, 1)

		if mouseBlink&0x10 == 0 {
			canvasSetColor(1, 0, 0, 1)
		} else {
			canvasSetColor(1, 1, 0, 1)
		}
		canvasRect(mouseX-8, mouseY-8, 16, 16, 1)

		mouseBlink++
	}
}

const (
	AtlasW = 128
	AtlasH = 128
	GlyphW = 8
	GlyphH = 8

	// frame buffer
	FBW = 256
	FBH = 256
)

var framebuffer = make([]byte, FBW*FBH*4)

//export sunani_fb_ptr
func FBPtr() uint32 { return uint32(uintptr(unsafe.Pointer(&framebuffer[0]))) }

//export sunani_fb_width
func FBW_() uint32 { return FBW }

//export sunani_fb_height
func FBH_() uint32 { return FBH }

func Clear(r, g, b, a uint32) {
	for i := 0; i < len(framebuffer); i += 4 {
		framebuffer[i+0] = byte(r)
		framebuffer[i+1] = byte(g)
		framebuffer[i+2] = byte(b)
		framebuffer[i+3] = byte(a)
	}
}

func DrawText(x, y uint32, strPtr uint32, strLen uint32) {
	s := bytesFromWasm(strPtr, strLen)

	cx := int(x)
	cy := int(y)

	for _, ch := range s {
		drawGlyph(cx, cy, ch)
		cx += GlyphW
	}
}

func drawGlyph(dstX, dstY int, ch byte) {
	// ASCII 0..127, 16x8 tiles
	tx := int(ch%16) * GlyphW
	ty := int(ch/16) * GlyphH

	for gy := 0; gy < GlyphH; gy++ {
		sy := ty + gy
		dy := dstY + gy
		if dy < 0 || dy >= FBH {
			continue
		}
		for gx := 0; gx < GlyphW; gx++ {
			sx := tx + gx
			dx := dstX + gx
			if dx < 0 || dx >= FBW {
				continue
			}

			si := (sy*AtlasW + sx) * 4
			di := (dy*FBW + dx) * 4

			sr := FontAtlasRGBA[si+0]
			sg := FontAtlasRGBA[si+1]
			sb := FontAtlasRGBA[si+2]
			sa := FontAtlasRGBA[si+3]

			// skip transparent pixels
			if sa == 0 {
				continue
			}

			// alpha blending
			dr := framebuffer[di+0]
			dg := framebuffer[di+1]
			db := framebuffer[di+2]
			da := framebuffer[di+3]

			a := uint32(sa)
			inv := 255 - a

			framebuffer[di+0] = byte((uint32(sr)*a + uint32(dr)*inv) / 255)
			framebuffer[di+1] = byte((uint32(sg)*a + uint32(dg)*inv) / 255)
			framebuffer[di+2] = byte((uint32(sb)*a + uint32(db)*inv) / 255)
			framebuffer[di+3] = byte((a + uint32(da)*inv/255)) // ざっくり
		}
	}
}

func bytesFromWasm(ptr uint32, n uint32) []byte {
	if n == 0 {
		return nil
	}
	return unsafeSlice(ptr, n)
}

func unsafeSlice(ptr uint32, n uint32) []byte {
	return *(*[]byte)(unsafe.Pointer(&struct {
		addr uintptr
		len  int
		cap  int
	}{uintptr(ptr), int(n), int(n)}))
}

var textBuf = make([]byte, 256)

func TextBufPtr() uint32 {
	return uint32(uintptr(unsafe.Pointer(&textBuf[0])))
}

//export sunani_fb_draw
func FBDraw() {
	hello := "Hello, Sunani!"

	Clear(0, 0, 0, 0)
	for i, b := range []byte(hello) {
		textBuf[i] = b
	}
	DrawText(16, 32, TextBufPtr(), uint32(len(hello)))
}

//export sunani_input_key
func InputKey(key uint32, action uint32) {
	if key == uint32(api.KeyQ) && action == uint32(api.KeyPress) {
		systemQuit()
	}
}

//export sunani_input_mouse
func InputMouse(action uint32, x float32, y float32, dx float32, dy float32, wheelX float32, wheelY float32, button uint32) {
	if action == uint32(api.MouseMove) {
		mouseEnabled = true
		mouseX = x
		mouseY = y
		mouseDX = dx
		mouseDY = dy
	} else if action == uint32(api.MousePress) {
		if button == uint32(api.MouseLeft) {
			anchorEnabled = true
			anchorX = x
			anchorY = y
		} else if button == uint32(api.MouseRight) {
			anchorEnabled = false
		}
	}
}

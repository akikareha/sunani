package main

import (
	"unsafe"

	"github.com/akikareha/sunani/app-go/api/fb"
	"github.com/akikareha/sunani/app-go/api/runtime"
	"github.com/akikareha/sunani/resources/fonts"
)

const hello = "Hello, World!\n"
const size = 16

const (
	atlasWidth  = 128
	atlasHeight = 128
	glyphWidth  = 8
	glyphHeight = 8

	// frame buffer
	fbWidth  = 256
	fbHeight = 256
)

var framebuffer = make([]byte, fbWidth*fbHeight*4)

func fbClear(r, g, b, a uint32) {
	for i := 0; i < len(framebuffer); i += 4 {
		framebuffer[i+0] = byte(r)
		framebuffer[i+1] = byte(g)
		framebuffer[i+2] = byte(b)
		framebuffer[i+3] = byte(a)
	}
}

func drawText(x, y uint32, s string) {
	cx := int(x)
	cy := int(y)

	for _, r := range s {
		drawGlyph(cx, cy, byte(r))
		cx += glyphWidth
	}
}

func drawGlyph(dstX, dstY int, ch byte) {
	// ASCII 0..127, 16x8 tiles
	tx := int(ch%16) * glyphWidth
	ty := int(ch/16) * glyphHeight

	for gy := 0; gy < glyphHeight; gy++ {
		sy := ty + gy
		dy := dstY + gy
		if dy < 0 || dy >= fbHeight {
			continue
		}
		for gx := 0; gx < glyphWidth; gx++ {
			sx := tx + gx
			dx := dstX + gx
			if dx < 0 || dx >= fbWidth {
				continue
			}

			si := (sy*atlasWidth + sx) * 4
			di := (dy*fbWidth + dx) * 4

			sr := fonts.FontAtlasRGBA[si+0]
			sg := fonts.FontAtlasRGBA[si+1]
			sb := fonts.FontAtlasRGBA[si+2]
			sa := fonts.FontAtlasRGBA[si+3]

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
			framebuffer[di+3] = byte((a + uint32(da)*inv/255))
		}
	}
}

//export sunani_runtime_init
func runtimeInit() {}

//export sunani_fb_init
func fbInit() {
	fb.Params(
		uint32(uintptr(unsafe.Pointer(&framebuffer[0]))),
		int32(fbWidth),
		int32(fbHeight),
	)
}

//export sunani_runtime_frame
func runtimeFrame() {
	runtime.Clear(0, 0, 0, 255)
	fbClear(0, 0, 0, 0)
	drawText(0, 0, hello)
	fb.Paint()
}

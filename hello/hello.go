package main

import "unsafe"

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

//export FBPtr
func FBPtr() uint32 { return uint32(uintptr(unsafe.Pointer(&framebuffer[0]))) }

//export FBW
func FBW_() uint32 { return FBW }

//export FBH
func FBH_() uint32 { return FBH }

//export Clear
func Clear(r, g, b, a uint32) {
	for i := 0; i < len(framebuffer); i += 4 {
		framebuffer[i+0] = byte(r)
		framebuffer[i+1] = byte(g)
		framebuffer[i+2] = byte(b)
		framebuffer[i+3] = byte(a)
	}
}

//export DrawText
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

//export TextBufPtr
func TextBufPtr() uint32 {
	return uint32(uintptr(unsafe.Pointer(&textBuf[0])))
}

func main() {}

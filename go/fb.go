package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type FB struct {
	window *glfw.Window

	fbPtr uint32
	fbW   int
	fbH   int

	tex uint32
}

func NewFB(window *glfw.Window) *FB {
	return &FB{
		window: window,
	}
}

func (fb *FB) Init() {
	fb.fbPtr = callU32("FBPtr")
	fb.fbW = int(callU32("FBW"))
	fb.fbH = int(callU32("FBH"))

	gl.GenTextures(1, &fb.tex)
	gl.BindTexture(gl.TEXTURE_2D, fb.tex)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(fb.fbW),
		int32(fb.fbH),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		nil,
	)
}

func (fb *FB) Begin() {
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(-1, 1, -1, 1, -1, 1)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	gl.Enable(gl.TEXTURE_2D)

	fw, fh := fb.window.GetFramebufferSize()
	scale := min(fw/fb.fbW, fh/fb.fbH)
	drawW := fb.fbW * scale
	drawH := fb.fbH * scale

	ox := (fw - drawW) / 2
	oy := (fh - drawH) / 2

	gl.Viewport(int32(ox), int32(oy), int32(drawW), int32(drawH))
}

func (fb *FB) Draw() {
	fbSize := fb.fbW * fb.fbH * 4
	pix, ok := mem.Read(fb.fbPtr, uint32(fbSize))
	if !ok {
		panic("mem.Read failed")
	}

	gl.BindTexture(gl.TEXTURE_2D, fb.tex)
	gl.TexSubImage2D(
		gl.TEXTURE_2D,
		0,
		0, 0,
		int32(fb.fbW), int32(fb.fbH),
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(pix),
	)

	gl.Color4f(1, 1, 1, 1)
	gl.Begin(gl.QUADS)
	gl.TexCoord2f(0, 1)
	gl.Vertex2f(-1, -1)
	gl.TexCoord2f(1, 1)
	gl.Vertex2f(1, -1)
	gl.TexCoord2f(1, 0)
	gl.Vertex2f(1, 1)
	gl.TexCoord2f(0, 0)
	gl.Vertex2f(-1, 1)
	gl.End()
}

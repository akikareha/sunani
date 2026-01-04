package main

import (
	"log"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tetratelabs/wazero/api"
)

type FB struct {
	window *glfw.Window

	ptr    uint32
	width  int
	height int

	tex uint32

	draw api.Function
}

func NewFB() *FB {
	return &FB{}
}

func (fb *FB) Preinit() {
	fb.draw = mod.ExportedFunction("sunani_fb_draw")
}

func (fb *FB) IsEnabled() bool {
	return fb.draw != nil
}

func (fb *FB) Init(window *glfw.Window) {
	if !fb.IsEnabled() {
		return
	}

	fb.window = window

	fb.ptr = callU32("sunani_fb_ptr")
	fb.width = int(callU32("sunani_fb_width"))
	fb.height = int(callU32("sunani_fb_height"))

	gl.GenTextures(1, &fb.tex)
	gl.BindTexture(gl.TEXTURE_2D, fb.tex)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(fb.width),
		int32(fb.height),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		nil,
	)
}

func (fb *FB) Begin() {
	if !fb.IsEnabled() {
		return
	}

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(-1, 1, -1, 1, -1, 1)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	gl.Disable(gl.DEPTH_TEST)
	gl.Enable(gl.TEXTURE_2D)

	fw, fh := fb.window.GetFramebufferSize()
	scale := min(fw/fb.width, fh/fb.height)
	width := fb.width * scale
	height := fb.height * scale

	ox := (fw - width) / 2
	oy := (fh - height) / 2

	gl.Viewport(int32(ox), int32(oy), int32(width), int32(height))
}

func (fb *FB) Draw() {
	if !fb.IsEnabled() {
		return
	}

	_, err := fb.draw.Call(ctx)
	if err != nil {
		log.Fatalln("fb draw call failed:", err)
	}

	size := fb.width * fb.height * 4
	pix, ok := mem.Read(fb.ptr, uint32(size))
	if !ok {
		panic("mem.Read failed")
	}

	gl.BindTexture(gl.TEXTURE_2D, fb.tex)
	gl.TexSubImage2D(
		gl.TEXTURE_2D,
		0,
		0, 0,
		int32(fb.width), int32(fb.height),
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

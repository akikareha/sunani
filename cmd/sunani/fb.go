package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tetratelabs/wazero/api"
)

var mem api.Memory

type FB struct {
	init api.Function

	window *glfw.Window

	ptr    uint32
	width  int
	height int

	tex uint32

	prepared bool
}

func NewFB() *FB {
	return &FB{}
}

func (fb *FB) Preinit() {
	fb.init = mod.ExportedFunction("sunani_fb_init")
}

func (fb *FB) IsEnabled() bool {
	return fb.init != nil
}

func (fb *FB) Init(window *glfw.Window) {
	if !fb.IsEnabled() {
		return
	}
	if window == nil {
		panic("window is nil")
	}
	fb.window = window

	mem = mod.ExportedMemory("memory")
	if mem == nil {
		die("wasm exported memory not found")
	}

	if fb.init != nil {
		_, err := fb.init.Call(ctx)
		if err != nil {
			die("sunani_fb_init call failed:", err)
		}
	}
}

func (fb *FB) Params(ptr uint32, width, height int) {
	if !fb.IsEnabled() {
		errlog("sunani fb.params was called, but Framebuffer API is not enabled.\nExport snunani_fb_init to enable this API.")
		return
	}
	if fb.window == nil {
		panic("window is nil")
	}

	fb.ptr = ptr
	fb.width = width
	fb.height = height

	if fb.prepared {
		gl.DeleteTextures(1, &fb.tex)
	}
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

	fb.prepared = true
}

func (fb *FB) Paint() {
	if !fb.IsEnabled() {
		errlog("sunani fb.paint was called, but Framebuffer API is not enabled.\nExport snunani_fb_init to enable this API.")
		return
	}
	if !fb.prepared {
		return
	}
	if fb.window == nil {
		panic("window is nil")
	}

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(-1, 1, -1, 1, -1, 1)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	gl.Disable(gl.DEPTH_TEST)
	gl.Enable(gl.TEXTURE_2D)

	fw, fh := fb.window.GetFramebufferSize()
	scale := min(fw/int(fb.width), fh/int(fb.height))
	width := int(fb.width) * scale
	height := int(fb.height) * scale

	ox := (fw - width) / 2
	oy := (fh - height) / 2

	gl.Viewport(int32(ox), int32(oy), int32(width), int32(height))

	size := fb.width * fb.height * 4
	if size < 0 {
		size = -size
	}
	if size < 1 {
		return
	}
	pix, ok := mem.Read(fb.ptr, uint32(size))
	if !ok {
		errlog("mem.Read failed")
		return
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

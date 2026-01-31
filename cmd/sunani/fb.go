package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tetratelabs/wazero/api"
)

var mem api.Memory

type FB struct {
	init api.Function
	rect api.Function

	window *glfw.Window

	ptr    uint32
	width  int
	height int

	tex uint32

	prepared bool

	ox, oy int
	rw, rh int
}

func NewFB() *FB {
	return &FB{}
}

func (fb *FB) Preinit() {
	fb.init = mod.ExportedFunction("sunani_fb_init")
	fb.rect = mod.ExportedFunction("sunani_fb_rect")
}

func (fb *FB) IsEnabled() bool {
	return fb.init != nil
}

func (fb *FB) DoResize(width, height int) {
	fb.updateRect()
	if fb.rect == nil {
		return
	}
	_, err := fb.rect.Call(
		ctx,
		uint64(fb.ox),
		uint64(fb.oy),
		uint64(fb.rw),
		uint64(fb.rh),
	)
	if err != nil {
		die("sunani_fb_rect call failed:", err)
	}
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

	fbw, fbh := window.GetFramebufferSize()
	fb.DoResize(fbw, fbh)
}

func (fb *FB) updateRect() {
	fw, fh := fb.window.GetFramebufferSize()
	if fb.width == 0 || fb.height == 0 {
		fb.rw = 0
		fb.rh = 0
	} else {
		scale := min(fw/int(fb.width), fh/int(fb.height))
		fb.rw = int(fb.width) * scale
		fb.rh = int(fb.height) * scale
	}
	fb.ox = (fw - fb.rw) / 2
	fb.oy = (fh - fb.rh) / 2
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

	fb.updateRect()

	fb.prepared = true
}

func (fb *FB) Begin() {
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

	gl.Viewport(int32(fb.ox), int32(fb.oy), int32(fb.rw), int32(fb.rh))
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

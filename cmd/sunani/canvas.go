package main

import (
	"log"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tetratelabs/wazero/api"
)

type Canvas struct {
	window *glfw.Window

	r, g, b, a float32

	init api.Function
}

func NewCanvas() *Canvas {
	return &Canvas{}
}

func (c *Canvas) Preinit() {
	c.init = mod.ExportedFunction("sunani_canvas_init")
}

func (c *Canvas) IsEnabled() bool {
	return c.init != nil
}

func (c *Canvas) Init(window *glfw.Window) {
	if !c.IsEnabled() {
		return
	}

	c.window = window

	_, err := c.init.Call(ctx)
	if err != nil {
		log.Fatalln("canvas init call failed:", err)
	}
}

func (c *Canvas) Begin() {
	if !c.IsEnabled() {
		return
	}

	fw, fh := c.window.GetFramebufferSize()

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(fw), float64(fh), 0, -1, 1)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	gl.Disable(gl.DEPTH_TEST)
	gl.Disable(gl.TEXTURE_2D)

	gl.Viewport(0, 0, int32(fw), int32(fh))
}

func (c *Canvas) Clear(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (c *Canvas) SetColor(r, g, b, a float32) {
	c.r, c.g, c.b, c.a = r, g, b, a
	gl.Color4f(r, g, b, a)
}

func (c *Canvas) Line(x1, y1, x2, y2 float32) {
	gl.Color4f(c.r, c.g, c.b, c.a)
	gl.Begin(gl.LINES)
	gl.Vertex2f(x1, y1)
	gl.Vertex2f(x2, y2)
	gl.End()
}

func (c *Canvas) Rect(x, y, w, h float32, fill bool) {
	gl.Color4f(c.r, c.g, c.b, c.a)
	if fill {
		gl.Begin(gl.QUADS)
	} else {
		gl.Begin(gl.LINE_LOOP)
	}
	gl.Vertex2f(x, y)
	gl.Vertex2f(x+w, y)
	gl.Vertex2f(x+w, y+h)
	gl.Vertex2f(x, y+h)
	gl.End()
}

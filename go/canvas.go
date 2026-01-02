package main

import "github.com/go-gl/gl/v2.1/gl"

type Canvas struct {
	width, height  int
	cr, cg, cb, ca float32
}

func NewCanvas(w, h int) *Canvas {
	return &Canvas{
		width:  w,
		height: h,
		cr:     1, cg: 1, cb: 1, ca: 1,
	}
}

func (c *Canvas) Init() {
	gl.Viewport(0, 0, int32(c.width), int32(c.height))

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(
		0, float64(c.width),
		float64(c.height), 0,
		-1, 1,
	)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	gl.Disable(gl.DEPTH_TEST)
}

func (c *Canvas) Clear(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (c *Canvas) SetColor(r, g, b, a float32) {
	c.cr, c.cg, c.cb, c.ca = r, g, b, a
	gl.Color4f(r, g, b, a)
}

func (c *Canvas) Line(x1, y1, x2, y2 float32) {
	gl.Color4f(c.cr, c.cg, c.cb, c.ca)
	gl.Begin(gl.LINES)
	gl.Vertex2f(x1, y1)
	gl.Vertex2f(x2, y2)
	gl.End()
}

func (c *Canvas) Rect(x, y, w, h float32, fill bool) {
	gl.Color4f(c.cr, c.cg, c.cb, c.ca)
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

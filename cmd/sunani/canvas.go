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
	points     []float32

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

func (c *Canvas) polygon(points []float32) {
	n := len(points) / 2
	if n < 2 {
		return
	}

	for i := 0; i < n-1; i++ {
		c.Line(
			points[i*2], points[i*2+1],
			points[(i+1)*2], points[(i+1)*2+1],
		)
	}

	// close
	c.Line(
		points[(n-1)*2], points[(n-1)*2+1],
		points[0], points[1],
	)
}

func cross(ax, ay, bx, by, cx, cy float32) float32 {
	return (bx-ax)*(cy-ay) - (by-ay)*(cx-ax)
}

func pointInTriangle(px, py float32,
	ax, ay, bx, by, cx, cy float32) bool {

	c1 := cross(px, py, ax, ay, bx, by)
	c2 := cross(px, py, bx, by, cx, cy)
	c3 := cross(px, py, cx, cy, ax, ay)

	return (c1 >= 0 && c2 >= 0 && c3 >= 0) ||
		(c1 <= 0 && c2 <= 0 && c3 <= 0)
}

func triangulatePolygon(points []float32) []float32 {
	tris := []float32{}
	verts := append([]float32{}, points...)

	for len(verts) >= 6 {
		n := len(verts) / 2
		earFound := false

		for i := 0; i < n; i++ {
			i0 := (i + n - 1) % n
			i1 := i
			i2 := (i + 1) % n

			ax, ay := verts[i0*2], verts[i0*2+1]
			bx, by := verts[i1*2], verts[i1*2+1]
			cx, cy := verts[i2*2], verts[i2*2+1]

			if cross(ax, ay, bx, by, cx, cy) <= 0 {
				continue
			}

			isEar := true
			for j := 0; j < n; j++ {
				if j == i0 || j == i1 || j == i2 {
					continue
				}
				px, py := verts[j*2], verts[j*2+1]
				if pointInTriangle(px, py, ax, ay, bx, by, cx, cy) {
					isEar = false
					break
				}
			}

			if isEar {
				tris = append(tris,
					ax, ay, bx, by, cx, cy,
				)

				verts = append(
					verts[:i1*2],
					verts[(i1+1)*2:]...,
				)
				earFound = true
				break
			}
		}

		if !earFound {
			break
		}
	}
	return tris
}

func (c *Canvas) triangle(x1, y1, x2, y2, x3, y3 float32) {
	gl.Color4f(c.r, c.g, c.b, c.a)
	gl.Begin(gl.TRIANGLES)
	gl.Vertex2f(x1, y1)
	gl.Vertex2f(x2, y2)
	gl.Vertex2f(x3, y3)
	gl.End()
}

func (c *Canvas) fillPolygon(points []float32) {
	tris := triangulatePolygon(points)

	for i := 0; i < len(tris); i += 6 {
		c.triangle(
			tris[i], tris[i+1],
			tris[i+2], tris[i+3],
			tris[i+4], tris[i+5],
		)
	}
}

func (c *Canvas) Path(x, y float32) {
	c.points = []float32{x, y}
}

func (c *Canvas) Vertex(x, y float32) {
	c.points = append(c.points, x, y)
}

func (c *Canvas) Polygon() {
	c.polygon(c.points)
}

func (c *Canvas) FillPolygon() {
	c.fillPolygon(c.points)
}

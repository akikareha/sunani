package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tetratelabs/wazero/api"
)

type Canvas struct {
	init api.Function

	window *glfw.Window

	r, g, b, a uint32
	points     []uint32
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
	if window == nil {
		panic("window is nil")
	}
	c.window = window

	if c.init != nil {
		_, err := c.init.Call(ctx)
		if err != nil {
			die("sunani_canvas_init call failed:", err)
		}
	}
}

func (c *Canvas) Begin() {
	if !c.IsEnabled() {
		errlog("sunani canvas.begin was called, but Canvas API is not enabled.\nExport snunani_canvas_init to enable this API.")
		return
	}
	if c.window == nil {
		panic("window is nil")
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

func (c *Canvas) Clear(r, g, b, a uint32) {
	if !c.IsEnabled() {
		errlog("sunani canvas.clear was called, but Canvas API is not enabled.\nExport snunani_canvas_init to enable this API.")
		return
	}
	if c.window == nil {
		panic("window is nil")
	}

	gl.ClearColor(
		float32(r)/255,
		float32(g)/255,
		float32(b)/255,
		float32(a)/255,
	)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (c *Canvas) Color(r, g, b, a uint32) {
	if !c.IsEnabled() {
		errlog("sunani canvas.color was called, but Canvas API is not enabled.\nExport snunani_canvas_init to enable this API.")
		return
	}
	if c.window == nil {
		panic("window is nil")
	}

	c.r, c.g, c.b, c.a = r, g, b, a
	gl.Color4ub(uint8(r), uint8(g), uint8(b), uint8(a))
}

func (c *Canvas) Line(x1, y1 uint32, x2, y2 uint32) {
	if !c.IsEnabled() {
		errlog("sunani canvas.line was called, but Canvas API is not enabled.\nExport snunani_canvas_init to enable this API.")
		return
	}
	if c.window == nil {
		panic("window is nil")
	}

	gl.Color4ub(uint8(c.r), uint8(c.g), uint8(c.b), uint8(c.a))
	gl.Begin(gl.LINES)
	gl.Vertex2f(float32(x1)+0.5, float32(y1)+0.5)
	gl.Vertex2f(float32(x2)+0.5, float32(y2)+0.5)
	gl.End()
}

func (c *Canvas) Rect(x, y uint32, w, h uint32) {
	if !c.IsEnabled() {
		errlog("sunani canvas.rect was called, but Canvas API is not enabled.\nExport snunani_canvas_init to enable this API.")
		return
	}
	if c.window == nil {
		panic("window is nil")
	}

	gl.Color4ub(uint8(c.r), uint8(c.g), uint8(c.b), uint8(c.a))
	gl.Begin(gl.LINE_LOOP)
	gl.Vertex2f(float32(x)+0.5, float32(y)+0.5)
	gl.Vertex2f(float32(x+w)+0.5, float32(y)+0.5)
	gl.Vertex2f(float32(x+w)+0.5, float32(y+h)+0.5)
	gl.Vertex2f(float32(x)+0.5, float32(y+h)+0.5)
	gl.End()
}

func (c *Canvas) FillRect(x, y uint32, w, h uint32) {
	if !c.IsEnabled() {
		errlog("sunani canvas.fill_rect was called, but Canvas API is not enabled.\nExport snunani_canvas_init to enable this API.")
		return
	}
	if c.window == nil {
		panic("window is nil")
	}

	gl.Color4ub(uint8(c.r), uint8(c.g), uint8(c.b), uint8(c.a))
	gl.Begin(gl.QUADS)
	gl.Vertex2i(int32(x), int32(y))
	gl.Vertex2i(int32(x+w), int32(y))
	gl.Vertex2i(int32(x+w), int32(y+h))
	gl.Vertex2i(int32(x), int32(y+h))
	gl.End()
}

func (c *Canvas) polygon(points []uint32) {
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

func cross(ax, ay, bx, by, cx, cy uint32) int32 {
	return int32((bx-ax)*(cy-ay) - (by-ay)*(cx-ax))
}

func pointInTriangle(px, py uint32,
	ax, ay, bx, by, cx, cy uint32) bool {

	c1 := cross(px, py, ax, ay, bx, by)
	c2 := cross(px, py, bx, by, cx, cy)
	c3 := cross(px, py, cx, cy, ax, ay)

	return (c1 >= 0 && c2 >= 0 && c3 >= 0) ||
		(c1 <= 0 && c2 <= 0 && c3 <= 0)
}

func triangulatePolygon(points []uint32) []uint32 {
	tris := []uint32{}
	verts := append([]uint32{}, points...)

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

func (c *Canvas) triangle(x1, y1, x2, y2, x3, y3 uint32) {
	gl.Color4ub(uint8(c.r), uint8(c.g), uint8(c.b), uint8(c.a))
	gl.Begin(gl.TRIANGLES)
	gl.Vertex2i(int32(x1), int32(y1))
	gl.Vertex2i(int32(x2), int32(y2))
	gl.Vertex2i(int32(x3), int32(y3))
	gl.End()
}

func (c *Canvas) fillPolygon(points []uint32) {
	tris := triangulatePolygon(points)

	for i := 0; i < len(tris); i += 6 {
		c.triangle(
			tris[i], tris[i+1],
			tris[i+2], tris[i+3],
			tris[i+4], tris[i+5],
		)
	}
}

func (c *Canvas) Path(x, y uint32) {
	if !c.IsEnabled() {
		errlog("sunani canvas.path was called, but Canvas API is not enabled.\nExport snunani_canvas_init to enable this API.")
		return
	}
	if c.window == nil {
		panic("window is nil")
	}

	c.points = []uint32{x, y}
}

func (c *Canvas) Vertex(x, y uint32) {
	if !c.IsEnabled() {
		errlog("sunani canvas.vertex was called, but Canvas API is not enabled.\nExport snunani_canvas_init to enable this API.")
		return
	}
	if c.window == nil {
		panic("window is nil")
	}

	c.points = append(c.points, x, y)
}

func (c *Canvas) Polygon() {
	if !c.IsEnabled() {
		errlog("sunani canvas.polygon was called, but Canvas API is not enabled.\nExport snunani_canvas_init to enable this API.")
		return
	}
	if c.window == nil {
		panic("window is nil")
	}

	c.polygon(c.points)
}

func (c *Canvas) FillPolygon() {
	if !c.IsEnabled() {
		errlog("sunani canvas.fill_polygon was called, but Canvas API is not enabled.\nExport snunani_canvas_init to enable this API.")
		return
	}
	if c.window == nil {
		panic("window is nil")
	}

	c.fillPolygon(c.points)
}

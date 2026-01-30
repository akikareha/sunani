package canvas

import (
	c "github.com/akikareha/sunani/app-go/api/canvas"
)

//export sunani_canvas_init
func canvasInit() {}

func Color(r, g, b, a int) {
	c.Color(uint32(r), uint32(g), uint32(b), uint32(a))
}

func Line(x1, y1 int, x2, y2 int) {
	c.Line(int32(x1), int32(y1), int32(x2), int32(y2))
}

func Rect(x, y int, w, h int) {
	c.Rect(int32(x), int32(y), int32(w), int32(h))
}

func FillRect(x, y int, w, h int) {
	c.FillRect(int32(x), int32(y), int32(w), int32(h))
}

func Path() {
	c.Path()
}

func Vertex(x, y int) {
	c.Vertex(int32(x), int32(y))
}

func Polygon() {
	c.Polygon()
}

func FillPolygon() {
	c.FillPolygon()
}

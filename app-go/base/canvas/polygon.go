package canvas

import (
	api "github.com/akikareha/sunani/app-go/api/canvas"
)

type Polygon []int

func (poly Polygon) Draw(x, y int, w, h int) {
	path()
	for i := 0; i < len(poly); i += 2 {
		if i+1 >= len(poly) {
			break
		}
		vertex(x+poly[i]*w/denomW, y+poly[i+1]*h/denomH)
	}
	polygon()
}

func (poly Polygon) Fill(x, y int, w, h int) {
	path()
	for i := 0; i < len(poly); i += 2 {
		if i+1 >= len(poly) {
			break
		}
		vertex(x+poly[i]*w/denomW, y+poly[i+1]*h/denomH)
	}
	fillPolygon()
}

func path() {
	api.Path()
}

func vertex(x, y int) {
	api.Vertex(int32(x), int32(y))
}

func polygon() {
	api.Polygon()
}

func fillPolygon() {
	api.FillPolygon()
}

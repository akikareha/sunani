package widget

import (
	"github.com/akikareha/sunani/app-go/base/canvas"
	"github.com/akikareha/sunani/app-go/base/mouse"
	"github.com/akikareha/sunani/app-go/base/screen"
)

type Cursor struct {
	polygon    canvas.Polygon
	r, g, b, a int
	mw, mh     int
	dx, dy     int
}

func NewCursor() *Cursor {
	cur := Cursor{}
	cur.SetColor(255, 255, 0, 128)
	cur.SetSize(48)
	cur.SetMargin(64, 64)
	return &cur
}

func (cur *Cursor) Draw() {
	width, height := screen.GetSize()
	x, y := mouse.GetPosition()

	if x < cur.mw {
		cur.dx = 1
	} else if x >= width-cur.mw {
		cur.dx = -1
	}
	if y < cur.mh {
		cur.dy = 1
	} else if y >= height-cur.mh {
		cur.dy = -1
	}

	denomW, denomH := canvas.GetDenoms()

	canvas.SetColor(cur.r, cur.g, cur.b, cur.a)
	cur.polygon.Fill(x, y, denomW*cur.dx, denomH*cur.dy)

	canvas.SetColor(255-cur.r, 255-cur.g, 255-cur.b, cur.a)
	cur.polygon.Draw(x, y, denomW*cur.dx, denomH*cur.dy)
}

func (cur *Cursor) SetColor(r, g, b, a int) {
	cur.r, cur.g, cur.b, cur.a = r, g, b, a
}

func createCursorPolygon(size int) canvas.Polygon {
	return canvas.Polygon{
		0, 0,
		size, size / 2,
		size / 2, size / 2,
		size / 2, size,
	}
}

func (cur *Cursor) SetSize(size int) {
	cur.polygon = createCursorPolygon(size)
}

func (cur *Cursor) SetMargin(w, h int) {
	cur.mw, cur.mh = w, h
}

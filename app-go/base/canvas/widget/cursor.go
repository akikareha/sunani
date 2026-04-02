package widget

import (
	"tea.kareha.org/loom/sunani/app-go/base/canvas"
	"tea.kareha.org/loom/sunani/app-go/base/color"
	"tea.kareha.org/loom/sunani/app-go/base/mouse"
	"tea.kareha.org/loom/sunani/app-go/base/screen"
)

type Cursor struct {
	polygon canvas.Polygon
	color   color.Color
	mw, mh  int
	dx, dy  int
}

func NewCursor() *Cursor {
	cur := Cursor{}
	cur.SetColor(color.New(255, 255, 0, 128))
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

	canvas.SetColor(cur.color)
	cur.polygon.Fill(x, y, denomW*cur.dx, denomH*cur.dy)

	canvas.SetColor(cur.color.Complement())
	cur.polygon.Draw(x, y, denomW*cur.dx, denomH*cur.dy)
}

func (cur *Cursor) SetColor(c color.Color) {
	cur.color = c
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

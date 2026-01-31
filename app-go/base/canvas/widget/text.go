package widget

import (
	"strings"

	"github.com/akikareha/sunani/app-go/base/canvas"
	"github.com/akikareha/sunani/app-go/base/font"
)

type Text struct {
	s          string
	r, g, b, a int
}

func NewText() *Text {
	return &Text{
		r: 255,
		g: 255,
		b: 255,
		a: 255,
	}
}

func (t *Text) Draw(x, y int, w, h int) {
	lines := strings.Split(t.s, "\n")
	n := len(lines)
	canvas.SetColor(t.r, t.g, t.b, t.a)
	for i := 0; i < n; i++ {
		font.ASCII.DrawString(x, y+(-n+i)*h, w, h, lines[i])
	}
}

func (t *Text) Get() string {
	return t.s
}

func (t *Text) Clear() {
	t.s = ""
}

func (t *Text) Add(s string) {
	t.s += s
}

func (t *Text) SetColor(r, g, b, a int) {
	t.r, t.g, t.b, t.a = r, g, b, a
}

package widget

import (
	"strings"

	"github.com/akikareha/sunani/app-go/base/canvas"
	"github.com/akikareha/sunani/app-go/base/color"
	"github.com/akikareha/sunani/app-go/base/font"
)

type Text struct {
	s          string
	color color.Color
}

func NewText() *Text {
	return &Text{
		color: color.New(255, 255, 255, 255),
	}
}

func (t *Text) Draw(x, y int, w, h int) {
	lines := strings.Split(t.s, "\n")
	n := len(lines)
	canvas.SetColor(t.color)
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

func (t *Text) SetColor(c color.Color) {
	t.color = c
}

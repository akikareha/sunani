package canvas

import (
	api "tea.kareha.org/loom/sunani/app-go/api/canvas"
	"tea.kareha.org/loom/sunani/app-go/base/color"
)

var fg = color.New(255, 255, 255, 255)

//export sunani_canvas_init
func canvasInit() {
	SetColor(fg)
}

func GetColor() color.Color {
	return fg
}

func SetColor(c color.Color) {
	fg = c
	r, g, b, a := c.Values()
	api.Color(uint32(r), uint32(g), uint32(b), uint32(a))
}

func DrawLine(x1, y1 int, x2, y2 int) {
	api.Line(int32(x1), int32(y1), int32(x2), int32(y2))
}

func DrawRect(x, y int, w, h int) {
	api.Rect(int32(x), int32(y), int32(w), int32(h))
}

func FillRect(x, y int, w, h int) {
	api.FillRect(int32(x), int32(y), int32(w), int32(h))
}

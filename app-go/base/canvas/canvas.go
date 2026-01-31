package canvas

import (
	api "github.com/akikareha/sunani/app-go/api/canvas"
)

var fgR, fgG, fgB, fgA = 255, 255, 255, 255

//export sunani_canvas_init
func canvasInit() {
	SetColor(fgR, fgG, fgB, fgA)
}

func GetColor() (int, int, int, int) {
	return fgR, fgG, fgB, fgA
}

func SetColor(r, g, b, a int) {
	fgR, fgG, fgB, fgA = r, g, b, a
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

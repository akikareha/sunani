package canvas

//go:wasmimport sunani canvas.begin
func Begin()

//go:wasmimport sunani canvas.clear
func Clear(r, g, b, a float32)

//go:wasmimport sunani canvas.color
func Color(r, g, b, a float32)

//go:wasmimport sunani canvas.line
func Line(x1, y1, x2, y2 float32)

//go:wasmimport sunani canvas.rect
func Rect(x, y, w, h float32)

//go:wasmimport sunani canvas.fill_rect
func FillRect(x, y, w, h float32)

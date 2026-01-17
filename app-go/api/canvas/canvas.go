package canvas

//go:wasmimport sunani canvas.begin
func Begin()

//go:wasmimport sunani canvas.clear
func Clear(r, g, b, a uint32)

//go:wasmimport sunani canvas.color
func Color(r, g, b, a uint32)

//go:wasmimport sunani canvas.line
func Line(x1, y1 uint32, x2, y2 uint32)

//go:wasmimport sunani canvas.rect
func Rect(x, y uint32, w, h uint32)

//go:wasmimport sunani canvas.fill_rect
func FillRect(x, y uint32, w, h uint32)

//go:wasmimport sunani canvas.path
func Path()

//go:wasmimport sunani canvas.vertex
func Vertex(x, y uint32)

//go:wasmimport sunani canvas.polygon
func Polygon()

//go:wasmimport sunani canvas.fill_polygon
func FillPolygon()

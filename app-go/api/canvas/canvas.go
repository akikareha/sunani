package canvas

//go:wasmimport sunani canvas.color
func Color(r, g, b, a uint32)

//go:wasmimport sunani canvas.line
func Line(x1, y1 int32, x2, y2 int32)

//go:wasmimport sunani canvas.rect
func Rect(x, y int32, w, h int32)

//go:wasmimport sunani canvas.fill_rect
func FillRect(x, y int32, w, h int32)

//go:wasmimport sunani canvas.path
func Path()

//go:wasmimport sunani canvas.vertex
func Vertex(x, y int32)

//go:wasmimport sunani canvas.polygon
func Polygon()

//go:wasmimport sunani canvas.fill_polygon
func FillPolygon()

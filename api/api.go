package api

//go:wasmimport sunani system_quit
func SystemQuit()

//go:wasmimport sunani canvas_clear
func CanvasClear(r, g, b, a float32)

//go:wasmimport sunani canvas_color
func CanvasColor(r, g, b, a float32)

//go:wasmimport sunani canvas_line
func CanvasLine(x1, y1, x2, y2 float32)

//go:wasmimport sunani canvas_rect
func CanvasRect(x, y, w, h float32, fill uint32)

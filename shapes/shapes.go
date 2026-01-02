package main

//go:wasmimport canvas clear
func canvasClear(r, g, b, a float32)

//go:wasmimport canvas setColor
func canvasSetColor(r, g, b, a float32)

//go:wasmimport canvas line
func canvasLine(x1, y1, x2, y2 float32)

//go:wasmimport canvas rect
func canvasRect(x, y, w, h float32, fill uint32)

//export draw
func draw() {
	canvasClear(0.10, 0.10, 0.15, 1.0)

	canvasSetColor(1, 1, 1, 1)
	canvasLine(50, 50, 300, 200)

	canvasSetColor(0.2, 0.8, 0.4, 1)
	canvasRect(100, 300, 200, 120, 1)

	canvasSetColor(1, 0.3, 0.3, 1)
	canvasRect(400, 100, 180, 180, 0)
}

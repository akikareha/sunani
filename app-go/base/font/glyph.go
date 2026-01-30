package font

import (
	"github.com/akikareha/sunani/app-go/base/canvas"
)

// 0: grid width
// 1: grid height
// 2: grid advance width
// 3: multi-polyline data
type Glyph []int8

func (g Glyph) Advance(w int) int {
	if len(g) < 3 {
		// invalid glyph
		return w
	}
	gw := int(g[0])
	ga := int(g[2])
	return w * ga / gw
}

func (g Glyph) Draw(x, y int, w, h int) {
	if len(g) < 4 {
		// invalid glyph
		return
	}
	gw := int(g[0])
	gh := int(g[1])

	i := 3
	polylines := int(g[i])
	i++
	for n := 0; n < polylines; n++ {
		if i >= len(g) {
			// invalid glyph
			return
		}
		vertices := int(g[i])
		i++
		if vertices < 1 {
			// invalid glyph
			continue
		}
		if i+vertices*2 > len(g) {
			// invalid glyph
			return
		}
		x1 := int(g[i])
		i++
		y1 := int(g[i])
		i++
		for m := 1; m < vertices; m++ {
			if i+1 >= len(g) {
				return
			}
			x2 := int(g[i])
			i++
			y2 := int(g[i])
			i++
			canvas.Line(
				x+x1*w/gw,
				y+y1*h/gh,
				x+x2*w/gw,
				y+y2*h/gh,
			)
			x1 = x2
			y1 = y2
		}
	}
}

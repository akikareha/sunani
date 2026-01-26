package font

import (
	"github.com/akikareha/sunani/app-go/api/canvas"
)

type Font struct {
	GridWidth  int8
	GridHeight int8
	Glyphs     [][]int8
	Offset     rune
}

func (f *Font) Draw(x, y int, width, height int, r rune) {
	if r < f.Offset || r > f.Offset+rune(len(f.Glyphs)) {
		return
	}
	glyph := f.Glyphs[r-f.Offset]

	if len(glyph) < 1 {
		return
	}
	i := 0
	polylines := int(glyph[i])
	i++
	for k := 0; k < polylines; k++ {
		if i >= len(glyph) {
			return
		}
		segments := int(glyph[i])
		i++
		if segments < 1 {
			continue
		}
		if i+1 >= len(glyph) {
			return
		}
		x1 := int(glyph[i])
		i++
		y1 := int(glyph[i])
		i++
		for j := 1; j < segments; j++ {
			if i+1 >= len(glyph) {
				return
			}
			x2 := int(glyph[i])
			i++
			y2 := int(glyph[i])
			i++
			canvas.Line(
				int32(x+x1*width/int(f.GridWidth)),
				int32(y+y1*height/int(f.GridHeight)),
				int32(x+x2*width/int(f.GridWidth)),
				int32(y+y2*height/int(f.GridHeight)),
			)
			x1 = x2
			y1 = y2
		}
	}
}

func (f *Font) DrawString(x, y int, sx, sy int, s string) {
	i := 0
	for _, r := range s {
		f.Draw(x+i*sx, y, sx, sy, r)
		i++
	}
}

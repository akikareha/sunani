package font

import (
	"github.com/akikareha/sunani/app-go/api/canvas"
)

type Glyphs struct {
	start  rune
	width  int8
	height int8
	data   [][]int8
}

func NewGlyphs(start rune, width, height int8, data [][]int8) Glyphs {
	return Glyphs{
		start:  start,
		width:  width,
		height: height,
		data:   data,
	}
}

func (g *Glyphs) Has(r rune) bool {
	return r >= g.start && r < g.start+rune(len(g.data))
}

func (g *Glyphs) Draw(x, y int, width, height int, r rune) {
	if !g.Has(r) {
		return
	}
	glyph := g.data[r-g.start]

	if len(glyph) < 1 {
		// invalid glyph
		return
	}
	i := 0
	polylines := int(glyph[i])
	i++
	for n := 0; n < polylines; n++ {
		if i >= len(glyph) {
			// invalid glyph
			return
		}
		vertices := int(glyph[i])
		i++
		if vertices < 1 {
			// invalid glyph
			continue
		}
		if i+vertices*2 > len(glyph) {
			// invalid glyph
			return
		}
		x1 := int(glyph[i])
		i++
		y1 := int(glyph[i])
		i++
		for m := 1; m < vertices; m++ {
			if i+1 >= len(glyph) {
				return
			}
			x2 := int(glyph[i])
			i++
			y2 := int(glyph[i])
			i++
			canvas.Line(
				int32(x+x1*width/int(g.width)),
				int32(y+y1*height/int(g.height)),
				int32(x+x2*width/int(g.width)),
				int32(y+y2*height/int(g.height)),
			)
			x1 = x2
			y1 = y2
		}
	}
}

func (g *Glyphs) DrawString(x, y int, width, height int, s string) {
	for i, r := range s {
		g.Draw(x+i*width, y, width, height, r)
	}
}

type Font struct {
	glyphs []Glyphs
}

func New(glyphs []Glyphs) Font {
	return Font{
		glyphs: glyphs,
	}
}

func (f *Font) Draw(x, y int, width, height int, r rune) {
	for _, g := range f.glyphs {
		if g.Has(r) {
			g.Draw(x, y, width, height, r)
			return
		}
	}
}

func (f *Font) DrawString(x, y int, width, height int, s string) {
	for i, r := range s {
		f.Draw(x+i*width, y, width, height, r)
	}
}

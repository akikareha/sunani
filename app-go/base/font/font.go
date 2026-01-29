package font

type Glyphs map[rune]Glyph

type Advance struct{
	start rune
	end rune
	ga int8
	gw int8
}

func (a Advance) Contains(r rune) bool {
	return r >= a.start && r <= a.end
}

func (a Advance) Amount(w int) int {
	return w * int(a.ga) / int(a.gw)
}

type Font struct{
	glyphs   []Glyphs
	advances []Advance
}

func New(glyphs []Glyphs, advances []Advance) Font {
	return Font{
		glyphs: glyphs,
		advances: advances,
	}
}

func (f *Font) Advance(w int, r rune) int {
	for _, glyphs := range f.glyphs {
		g := glyphs[r]
		if g != nil {
			return g.Advance(w)
		}
	}
	for _, advance := range f.advances {
		if advance.Contains(r) {
			return advance.Amount(w)
		}
	}
	return w
}

func (f *Font) Draw(x, y int, w, h int, r rune) {
	for _, glyphs := range f.glyphs {
		g := glyphs[r]
		if g != nil {
			g.Draw(x, y, w, h)
			return
		}
	}
}

func (f *Font) DrawString(x, y int, w, h int, s string) {
	for _, r := range s {
		f.Draw(x, y, w, h, r)
		x += f.Advance(w, r)
	}
}

func (f *Font) StringWidth(w int, s string) int {
	sum := 0
	for _, r := range s {
		sum += f.Advance(w, r)
	}
	return sum
}

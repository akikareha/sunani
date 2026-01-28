package font

type Font struct {
	glyphs []Glyphs
	kutens []KutenGlyphs
}

func New(glyphs []Glyphs, kutens []KutenGlyphs) Font {
	return Font{
		glyphs: glyphs,
		kutens: kutens,
	}
}

func (f *Font) IsWide(r rune) bool {
	for _, g := range f.glyphs {
		if g.Has(r) {
			return g.IsWide()
		}
	}
	return false
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
	for _, r := range s {
		f.Draw(x, y, width, height, r)
		if f.IsWide(r) {
			x += width * 2
		} else {
			x += width
		}
	}
}

func (f *Font) StringWidth(width int, s string) int {
	sum := 0
	for _, r := range s {
		if f.IsWide(r) {
			sum += width * 2
		} else {
			sum += width
		}
	}
	return sum
}

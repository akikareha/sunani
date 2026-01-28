package font

type Kuten struct {
	Ku  uint8
	Ten uint8
}

type KutenGlyphs struct {
	width  int8
	height int8
	toGlyph map[rune][]int8
}

type kutenKey uint16

func toKutenKey(ku, ten uint8) kutenKey {
	return kutenKey(uint16(ku) << 8 | uint16(ten))
}

func NewKutenGlyphs(
	width,
	height int8,
	data []Glyph,
	toRune map[Kuten]rune,
) KutenGlyphs {
	glyphByKuten := make(map[kutenKey][]int8)
	for _, g := range data {
		if len(g) < 2 {
			// invalid glyph
			continue
		}
		glyphByKuten[toKutenKey(uint8(g[0]), uint8(g[1]))] = g[2:]
	}

	toGlyph := make(map[rune][]int8)
	for k := range toRune {
		r := toRune[k]
		glyph := glyphByKuten[toKutenKey(k.Ku, k.Ten)]
		if glyph == nil {
			continue
		}
		toGlyph[r] = glyph
	}
	return KutenGlyphs{
		width:  width,
		height: height,
		toGlyph: toGlyph,
	}
}

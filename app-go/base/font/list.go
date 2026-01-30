package font

var ASCIIAdvance = NewAdvance(0, 127, 8, 16)

var ASCII = New(
	[]Glyphs{
		ASCIIGlyphs,
	},
	[]Advance{
		ASCIIAdvance,
	},
)

var Kana = New(
	[]Glyphs{
		ASCIIGlyphs,
		KanaGlyphs,
	},
	[]Advance{
		ASCIIAdvance,
	},
)

var Kanji = New(
	[]Glyphs{
		ASCIIGlyphs,
		KanaGlyphs,
		KanjiGlyphs,
	},
	[]Advance{
		ASCIIAdvance,
	},
)

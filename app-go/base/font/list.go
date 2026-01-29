package font

var ASCII = New(
	[]Glyphs{
		ASCIIGlyphs,
	},
	[]Advance{},
)

var Kana = New(
	[]Glyphs{
		ASCIIGlyphs,
		KanaGlyphs,
	},
	[]Advance{},
)

var Kanji = New(
	[]Glyphs{
		ASCIIGlyphs,
		KanaGlyphs,
		KanjiGlyphs,
	},
	[]Advance{},
)

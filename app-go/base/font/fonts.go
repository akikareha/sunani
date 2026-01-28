package font

var Default = New(
	[]Glyphs{
		ASCIIGlyphs,
		KanaGlyphs,
	},
	[]KutenGlyphs{
		Kanji1Glyphs,
	},
)

var ASCII = New(
	[]Glyphs{
		ASCIIGlyphs,
	},
	[]KutenGlyphs{},
)

var Full = New(
	[]Glyphs{
		ASCIIGlyphs,
		KanaGlyphs,
	},
	[]KutenGlyphs{
		Kanji1Glyphs,
	},
)

var Kana = New(
	[]Glyphs{
		ASCIIGlyphs,
		KanaGlyphs,
	},
	[]KutenGlyphs{},
)

var Kanji = New(
	[]Glyphs{
		ASCIIGlyphs,
		KanaGlyphs,
	},
	[]KutenGlyphs{
		Kanji1Glyphs,
	},
)

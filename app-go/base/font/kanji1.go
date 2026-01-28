package font

var Kanji1Glyphs = NewKutenGlyphs(8, 16, kanji1Data, Kanji1ToRune)

var kanji1Data = []Glyph{
    // 亜
	{
		16, 1,
		0, // dummy
	},
    // 唖
	{
		16, 2,
		0, // dummy
	},
}

var Kanji1ToRune = map[Kuten]rune{
    {16, 1}: '亜',
    {16, 2}: '唖',
    // ...
}

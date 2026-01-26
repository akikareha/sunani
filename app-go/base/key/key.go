package key

import (
	"github.com/akikareha/sunani/app-go/api/canvas"
	"github.com/akikareha/sunani/app-go/base/font"
	"github.com/akikareha/sunani/lib"
)

type Key struct {
	Code  lib.Key
	Label string
	// grid units
	X      int16
	Y      int16
	Width  int16
	Height int16
}

type Keyboard struct {
	Keys       []Key
	GridWidth  int8
	GridHeight int8
}

func (kb *Keyboard) Draw(ox, oy int, pitch int) {
	canvas.Color(64, 64, 64, 18)
	for i := 0; i < len(kb.Keys); i++ {
		k := kb.Keys[i]
		canvas.FillRect(
			int32(ox+int(k.X)*pitch/int(kb.GridWidth)+pitch/8),
			int32(oy+int(k.Y)*pitch/int(kb.GridHeight)+pitch/8),
			int32(int(k.Width)*pitch/int(kb.GridWidth)-pitch/4),
			int32(int(k.Height)*pitch/int(kb.GridHeight)-pitch/4),
		)
	}

	canvas.Color(255, 255, 255, 128)
	for i := 0; i < len(kb.Keys); i++ {
		k := kb.Keys[i]
		n := len(k.Label)
		width := pitch * int(k.Width) / int(kb.GridWidth) / n / 2
		height := pitch * int(k.Height) / int(kb.GridHeight) / 2
		dx := (pitch*int(k.Width)/int(kb.GridWidth) - width*n) / 2
		dy := (pitch*int(k.Height)/int(kb.GridHeight) - height) / 2
		font.Default.DrawString(
			ox+int(k.X)*pitch/int(kb.GridWidth)+dx,
			oy+int(k.Y)*pitch/int(kb.GridHeight)+dy,
			width,
			height,
			k.Label,
		)
	}
}

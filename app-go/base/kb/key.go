package kb

import (
	"github.com/akikareha/sunani/app-go/base/canvas"
	"github.com/akikareha/sunani/app-go/base/font"
	"github.com/akikareha/sunani/app-go/base/key"
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

func (kb *Keyboard) Draw(ox, oy int, pitch int, mouseX, mouseY int) lib.Key {
	found := lib.KeyUnknown

	for i := 0; i < len(kb.Keys); i++ {
		k := kb.Keys[i]
		if key.IsDown(k.Code) {
			canvas.SetColor(128, 128, 128, 128)
		} else {
			canvas.SetColor(64, 64, 64, 64)
		}
		canvas.FillRect(
			ox+int(k.X)*pitch/int(kb.GridWidth)+pitch/8,
			oy+int(k.Y)*pitch/int(kb.GridHeight)+pitch/8,
			int(k.Width)*pitch/int(kb.GridWidth)-pitch/4,
			int(k.Height)*pitch/int(kb.GridHeight)-pitch/4,
		)
		if mouseX >= ox+int(k.X)*pitch/int(kb.GridWidth) &&
			mouseX < ox+int(k.X)*pitch/int(kb.GridWidth)+int(k.Width)*pitch/int(kb.GridWidth) &&
			mouseY >= oy+int(k.Y)*pitch/int(kb.GridHeight) &&
			mouseY < oy+int(k.Y)*pitch/int(kb.GridHeight)+int(k.Height)*pitch/int(kb.GridHeight) {
			canvas.SetColor(128, 128, 0, 128)
			canvas.DrawRect(
				ox+int(k.X)*pitch/int(kb.GridWidth),
				oy+int(k.Y)*pitch/int(kb.GridHeight),
				int(k.Width)*pitch/int(kb.GridWidth),
				int(k.Height)*pitch/int(kb.GridHeight),
			)

			found = k.Code
		}
	}

	for i := 0; i < len(kb.Keys); i++ {
		k := kb.Keys[i]
		if key.IsDown(k.Code) {
			canvas.SetColor(0, 0, 0, 255)
		} else {
			canvas.SetColor(255, 255, 255, 128)
		}
		n := len(k.Label)
		width := pitch * int(k.Width) / int(kb.GridWidth) / n
		height := pitch * int(k.Height) / int(kb.GridHeight) / 2
		dx := (pitch*int(k.Width)/int(kb.GridWidth) - width*n/2) / 2
		dy := (pitch*int(k.Height)/int(kb.GridHeight) - height) / 2
		font.ASCII.DrawString(
			ox+int(k.X)*pitch/int(kb.GridWidth)+dx,
			oy+int(k.Y)*pitch/int(kb.GridHeight)+dy,
			width,
			height,
			k.Label,
		)
	}

	return found
}

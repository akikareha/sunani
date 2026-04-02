package fb

import (
	"tea.kareha.org/loom/sunani/app-go/base/color"
)

var clear = color.New(0, 0, 0, 0)

var bg = color.New(0, 0, 0, 0)
var fg = color.New(255, 255, 255, 255)

func GetBgColor() color.Color {
	return bg
}

func SetBgColor(c color.Color) {
	bg = c
}

func GetColor() color.Color {
	return fg
}

func SetColor(c color.Color) {
	fg = c
}

func Clear() {
	bgr, bgg, bgb, bga := bg.Values()
	r, g, b, a := uint8(bgr), uint8(bgg), uint8(bgb), uint8(bga)
	for i := 0; i < len(buffer); i += 4 {
		buffer[i] = r
		buffer[i+1] = g
		buffer[i+2] = b
		buffer[i+3] = a
	}
}

func GetPixel(x, y int) color.Color {
	if x < 0 || x >= w {
		return clear
	}
	if y < 0 || y >= h {
		return clear
	}
	i := (y*w + x) * 4
	return color.New(
		int(buffer[i]),
		int(buffer[i+1]),
		int(buffer[i+2]),
		int(buffer[i+3]),
	)
}

func DrawPixel(x, y int) {
	if x < 0 || x >= w {
		return
	}
	if y < 0 || y >= h {
		return
	}
	i := (y*w + x) * 4
	r, g, b, a := fg.Values()
	buffer[i] = uint8(r)
	buffer[i+1] = uint8(g)
	buffer[i+2] = uint8(b)
	buffer[i+3] = uint8(a)
}

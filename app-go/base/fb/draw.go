package fb

var bgR, bgG, bgB, bgA uint8 = 0, 0, 0, 0
var fgR, fgG, fgB, fgA uint8 = 255, 255, 255, 255

func GetBgColor() (int, int, int, int) {
	return int(bgR), int(bgG), int(bgB), int(bgA)
}

func SetBgColor(r, g, b, a int) {
	bgR, bgG, bgB, bgA = uint8(r), uint8(g), uint8(b), uint8(a)
}

func GetColor() (int, int, int, int) {
	return int(fgR), int(fgG), int(fgB), int(fgA)
}

func SetColor(r, g, b, a int) {
	fgR, fgG, fgB, fgA = uint8(r), uint8(g), uint8(b), uint8(a)
}

func Clear() {
	for i := 0; i < len(buffer); i += 4 {
		buffer[i] = bgR
		buffer[i+1] = bgG
		buffer[i+2] = bgB
		buffer[i+3] = bgA
	}
}

func GetPixel(x, y int) (int, int, int, int) {
	if x < 0 || x >= w {
		return 0, 0, 0, 0
	}
	if y < 0 || y >= h {
		return 0, 0, 0, 0
	}
	i := (y*w + x) * 4
	return int(buffer[i]),
		 int(buffer[i+1]),
		 int(buffer[i+2]),
		 int(buffer[i+3])
}

func DrawPixel(x, y int) {
	if x < 0 || x >= w {
		return
	}
	if y < 0 || y >= h {
		return
	}
	i := (y*w + x) * 4
	buffer[i] = fgR
	buffer[i+1] = fgG
	buffer[i+2] = fgB
	buffer[i+3] = fgA
}

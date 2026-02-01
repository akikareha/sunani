package color

type Color struct {
	r, g, b, a uint8
}

func New(r, g, b, a int) Color {
	return Color{
		r: uint8(r),
		g: uint8(g),
		b: uint8(b),
		a: uint8(a),
	}
}

func (c Color) Values() (int, int, int, int) {
	return int(c.r), int(c.g), int(c.b), int(c.a)
}

func (c Color) Complement() Color {
	return Color{
		r: 255 - c.r,
		g: 255 - c.g,
		b: 255 - c.b,
		a: c.a,
	}
}

func (c Color) rotate() Color {
	return Color{
		r: c.g,
		g: c.b,
		b: c.r,
		a: c.a,
	}
}

func (c Color) invRotate() Color {
	return Color{
		r: c.b,
		g: c.r,
		b: c.g,
		a: c.a,
	}
}

func (c Color) ToFg() Color {
	average := (int(c.r) + int(c.g) + int(c.b)) * int(c.a) / 255 / 3
	if average < 64 {
		return New(255, 255, 255, 255)
	} else if average < 128 {
		return c.rotate()
	} else if average < 192 {
		return c.invRotate()
	} else {
		return New(0, 0, 0, 255)
	}
}

func (c Color) WithAlpha(a int) Color {
	return Color{
		r: c.r,
		g: c.g,
		b: c.b,
		a: uint8(a),
	}
}

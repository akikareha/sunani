package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"
)

func main() {
	in := flag.String("in", "ascii8x8.png", "input 128x128 PNG")
	out := flag.String("out", "wasm/font_atlas.go", "output .go file")
	pkg := flag.String("pkg", "main", "package name for output")
	varName := flag.String("var", "FontAtlasRGBA", "var name for atlas")
	flag.Parse()

	b, err := os.ReadFile(*in)
	must(err)

	img, err := png.Decode(bytes.NewReader(b))
	must(err)

	bounds := img.Bounds()
	if bounds.Dx() != 128 || bounds.Dy() != 128 {
		panic(fmt.Sprintf("expected 128x128, got %dx%d", bounds.Dx(), bounds.Dy()))
	}

	// RGBA: 4 bytes * 128 * 128 = 65536
	atlas := make([]byte, 128*128*4)
	for y := 0; y < 128; y++ {
		for x := 0; x < 128; x++ {
			r, g, bb, a := img.At(x, y).RGBA()
			// Convert 16-bit to 8-bit
			i := (y*128 + x) * 4
			atlas[i+0] = byte(r >> 8)
			atlas[i+1] = byte(g >> 8)
			atlas[i+2] = byte(bb >> 8)
			atlas[i+3] = byte(a >> 8)
		}
	}

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "package %s\n\n", *pkg)
	fmt.Fprintf(&buf, "// %s is 128x128 RGBA atlas: 65536 bytes.\n", *varName)
	fmt.Fprintf(&buf, "var %s = [%d]byte{\n", *varName, len(atlas))
	for i := 0; i < len(atlas); i++ {
		if i%16 == 0 {
			buf.WriteString("\t")
		}
		fmt.Fprintf(&buf, "0x%02x,", atlas[i])
		if i%16 == 15 {
			buf.WriteString("\n")
		}
	}
	if len(atlas)%16 != 0 {
		buf.WriteString("\n")
	}
	buf.WriteString("}\n")

	must(os.WriteFile(*out, buf.Bytes(), 0o644))
	fmt.Println("wrote:", *out)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// Ensure image/png is linked
var _ image.Image

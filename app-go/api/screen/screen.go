package screen

//go:wasmimport sunani screen.halt
func Halt()

//go:wasmimport sunani screen.title
func Title(ptr uint32, length uint32)

//go:wasmimport sunani screen.cursor
func Cursor(visible uint32)

//go:wasmimport sunani screen.clear
func Clear(r, g, b, a uint32)

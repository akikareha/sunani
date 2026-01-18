package runtime

//go:wasmimport sunani runtime.halt
func Halt()

//go:wasmimport sunani runtime.title
func Title(ptr uint32, length uint32)

//go:wasmimport sunani runtime.cursor
func Cursor(enabled uint32)

//go:wasmimport sunani runtime.clear
func Clear(r, g, b, a uint32)

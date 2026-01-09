package system

//go:wasmimport sunani system.halt
func Halt()

//go:wasmimport sunani system.title
func Title(ptr uint32, len uint32)

//go:wasmimport sunani system.cursor
func Cursor(enabled uint32)

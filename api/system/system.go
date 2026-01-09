package system

//go:wasmimport sunani system.halt
func Halt()

//go:wasmimport sunani system.cursor
func Cursor(enabled uint32)

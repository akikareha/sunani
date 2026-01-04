package fb

//go:wasmimport sunani fb.params
func Params(ptr, width, height uint32)

//go:wasmimport sunani fb.paint
func Paint()

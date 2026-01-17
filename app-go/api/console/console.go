package console

//go:wasmimport sunani console.params
func Params(ptr uint32, length uint32)

//go:wasmimport sunani console.put
func Put(ptr uint32, length uint32)

package main

import (
	"unsafe"

	"tea.kareha.org/loom/sunani/app-go/api/console"
)

const hello = "Hello, World!\n"

//export sunani_console_init
func consoleInit() {
	b := []byte(hello)
	console.Put(
		uint32(uintptr(unsafe.Pointer(&b[0]))),
		uint32(len(hello)),
	)
}

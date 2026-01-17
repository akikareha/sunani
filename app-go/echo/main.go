package main

import (
	"fmt"
	"unsafe"

	"github.com/akikareha/sunani/app-go/api/console"
)

const consoleBufferLength = 1024

var consoleBuffer = make([]byte, consoleBufferLength)

const prompt = "> "

func putString(s string) {
	b := []byte(s)
	console.Put(
		uint32(uintptr(unsafe.Pointer(&b[0]))),
		uint32(len(s)),
	)
}

//export sunani_console_init
func consoleInit() {
	console.Params(
		uint32(uintptr(unsafe.Pointer(&consoleBuffer[0]))),
		uint32(consoleBufferLength),
	)

	putString("Type any text and I will echo it.\n")
	putString("Type \"quit\" to leave.\n")

	putString(prompt)

	console.Wait()
}

//export sunani_console_get
func consoleGet(ptr uint32, length uint32) {
	b := unsafe.Slice((*byte)(unsafe.Pointer(uintptr(ptr))), length)
	line := string(b)

	reply := fmt.Sprintf("You typed: %s\n", line)
	putString(reply)

	if line == "quit" {
		putString("Bye!\n")

		console.Leave()
		return
	}

	putString(prompt)
}

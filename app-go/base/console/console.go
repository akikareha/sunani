package console

import (
	"unsafe"

	con "github.com/akikareha/sunani/app-go/api/console"
)

type InputHandler func(line string)

const dummy = " "

var initialized bool
var bufferLength = 256
var buffer []byte
var inputHandlers = make([]InputHandler, 0)

//export sunani_console_init
func consoleInit() {
	initialized = true

	SetBufferLength(bufferLength)
}

//export sunani_console_get
func consoleGet(ptr uint32, length uint32) {
	b := unsafe.Slice((*byte)(unsafe.Pointer(uintptr(ptr))), length)

	line := string(b)
	for _, handler := range inputHandlers {
		handler(line)
	}
}

func AddInputHandler(handler InputHandler) {
	inputHandlers = append(inputHandlers, handler)
}

func SetBufferLength(length int) {
	bufferLength = length

	if !initialized {
		return
	}

	buffer = make([]byte, bufferLength)
	con.Params(
		uint32(uintptr(unsafe.Pointer(&buffer[0]))),
		uint32(bufferLength),
	)
}

func Print(s string) {
	if !initialized {
		return
	}
	Output([]byte(s))
}

func Output(b []byte) {
	if !initialized {
		return
	}
	length := len(b)
	if length < 1 {
		// cannot point null string
		b = []byte(dummy)
	}
	con.Put(
		uint32(uintptr(unsafe.Pointer(&b[0]))),
		uint32(length),
	)
}

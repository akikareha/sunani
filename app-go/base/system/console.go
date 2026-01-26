package system

import (
	"unsafe"

	"github.com/akikareha/sunani/app-go/api/console"
)

var consoleInitialized bool

var greeting string = "Hello, World!\n"

func SetGreeting(s string) {
	greeting = s
}

var prompt string = "> "

func SetPrompt(s string) {
	prompt = s
}

func Print(s string) {
	b := []byte(s)
	length := len(b)
	if length < 1 {
		// cannot point null string
		b = []byte(dummyString)
	}
	console.Put(
		uint32(uintptr(unsafe.Pointer(&b[0]))),
		uint32(length),
	)
	addConsoleLine(s)
}

var consoleBufferLength int = 1024
var consoleBuffer []byte

func GetConsoleBufferLength() int {
	return consoleBufferLength
}

func SetConsoleBufferLength(length int) {
	consoleBufferLength = length

	if !consoleInitialized {
		return
	}

	consoleBuffer = make([]byte, consoleBufferLength)
	console.Params(
		uint32(uintptr(unsafe.Pointer(&consoleBuffer[0]))),
		uint32(consoleBufferLength),
	)
}

//export sunani_console_init
func consoleInit() {
	consoleInitialized = true

	SetConsoleBufferLength(consoleBufferLength)

	Print(greeting)
	Print(prompt)
}

var consoleGetCallback func(string) bool = repl

func SetConsoleGet(callback func(string) bool) {
	consoleGetCallback = callback
}

//export sunani_console_get
func consoleGet(ptr uint32, length uint32) {
	if consoleGetCallback == nil {
		return
	}

	b := unsafe.Slice((*byte)(unsafe.Pointer(uintptr(ptr))), length)
	line := string(b)

	addConsoleLine(line + "\n")

	if consoleGetCallback(line) {
		Print(prompt)
	}
}

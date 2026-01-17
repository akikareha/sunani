package main

import (
	"unsafe"

	"github.com/akikareha/sunani/app-go/api/console"
	"github.com/akikareha/sunani/app-go/api/fb"
	"github.com/akikareha/sunani/app-go/api/system"
)

const title = "Demo"
const hello = "Hello, World!\n"

//export sunani_system_init
func systemInit() {
	b := []byte(title)
	system.Title(
		uint32(uintptr(unsafe.Pointer(&b[0]))),
		uint32(len(title)),
	)

	system.Cursor(0)
}

//export sunani_console_init
func consoleInit() {
	console.Params(
		uint32(uintptr(unsafe.Pointer(&consoleBuffer[0]))),
		uint32(consoleBufferLength),
	)

	b := []byte(hello)
	console.Put(
		uint32(uintptr(unsafe.Pointer(&b[0]))),
		uint32(len(hello)),
	)
}

//export sunani_canvas_init
func canvasInit() {}

//export sunani_fb_init
func fbInit() {
	fb.Params(
		uint32(uintptr(unsafe.Pointer(&framebuffer[0]))),
		uint32(fbWidth),
		uint32(fbHeight),
	)
}

//export sunani_key_init
func keyInit() {}

//export sunani_mouse_init
func mouseInit() {}

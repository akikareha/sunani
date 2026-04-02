package main

import (
	"unsafe"

	"tea.kareha.org/loom/sunani/app-go/api/console"
	"tea.kareha.org/loom/sunani/app-go/api/fb"
	"tea.kareha.org/loom/sunani/app-go/api/screen"
)

const title = "Demo"
const hello = "Hello, World!\n"

//export sunani_screen_init
func screenInit() {
	b := []byte(title)
	screen.Title(
		uint32(uintptr(unsafe.Pointer(&b[0]))),
		uint32(len(title)),
	)

	screen.Cursor(0)
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

//export sunani_fb_init
func fbInit() {
	fb.Params(
		uint32(uintptr(unsafe.Pointer(&framebuffer[0]))),
		int32(fbWidth),
		int32(fbHeight),
	)
}

//export sunani_key_init
func keyInit() {}

//export sunani_mouse_init
func mouseInit() {}

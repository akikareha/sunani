package main

import (
	"unsafe"

	"github.com/akikareha/sunani/api/fb"
	"github.com/akikareha/sunani/api/system"
)

const title = "Demo"

//export sunani_system_init
func systemInit() {
	b := []byte(title)
	system.Title(
		uint32(uintptr(unsafe.Pointer(&b[0]))),
		uint32(len(title)),
	)

	system.Cursor(0)
}

//export sunani_canvas_init
func canvasInit() {}

//export sunani_fb_init
func fbInit() {
	fb.Params(FBPtr(), FBW_(), FBH_())
}

//export sunani_key_init
func keyInit() {}

//export sunani_mouse_init
func mouseInit() {}

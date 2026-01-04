package main

import (
	"github.com/akikareha/sunani/api/fb"
)

//export sunani_system_init
func systemInit() {}

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

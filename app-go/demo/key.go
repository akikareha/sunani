package main

import (
	"github.com/akikareha/sunani/app-go/api/screen"
	"github.com/akikareha/sunani/lib"
)

//export sunani_key_event
func keyEvent(key uint32, action uint32) {
	k := lib.Key(key)
	a := lib.Action(action)

	if k == lib.KeyEscape && a == lib.ActionPress {
		screen.Halt()
	}
}

package main

import (
	"github.com/akikareha/sunani/api/system"
	"github.com/akikareha/sunani/lib"
)

//export sunani_key_event
func KeyEvent(key uint32, action uint32) {
	k := lib.Key(key)
	a := lib.Action(action)

	if k == lib.KeyQ && a == lib.ActionPress {
		system.Quit()
	}
}

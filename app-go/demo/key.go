package main

import (
	"tea.kareha.org/loom/sunani/app-go/api/screen"
	"tea.kareha.org/loom/sunani/lib"
)

//export sunani_key_event
func keyEvent(key uint32, action uint32) {
	k := lib.Key(key)
	a := lib.Action(action)

	if k == lib.KeyEscape && a == lib.ActionPress {
		screen.Halt()
	}
}

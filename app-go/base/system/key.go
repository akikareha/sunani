package system

import (
	bkey "github.com/akikareha/sunani/app-go/base/key"
	"github.com/akikareha/sunani/lib"
)

var keyTable []bool

//export sunani_key_init
func keyInit() {
	keyTable = make([]bool, lib.KeyCount)
}

//export sunani_key_event
func keyEvent(key uint32, action uint32) {
	k := lib.Key(key)
	a := lib.Action(action)

	if a == lib.ActionPress {
		keyTable[k] = true
		addConsoleLine(bkey.Char(k))
	} else if a == lib.ActionRelease {
		keyTable[k] = false
	}
}

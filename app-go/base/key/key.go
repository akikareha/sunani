package key

import (
	"github.com/akikareha/sunani/lib"
)

//export sunani_key_init
func keyInit() {}

var table []bool = make([]bool, lib.KeyCount)

func IsDown(key lib.Key) bool {
	// key cannot be negative since lib.Key is uint16.
	if key >= lib.KeyCount {
		return false
	}
	return table[key]
}

type Handler func(key lib.Key, action lib.Action)

var handlers []Handler = make([]Handler, 0)

func AddHandler(handler Handler) {
	handlers = append(handlers, handler)
}

//export sunani_key_event
func keyEvent(key uint32, action uint32) {
	k := lib.Key(key)
	a := lib.Action(action)

	if a == lib.ActionPress {
		table[k] = true
	} else if a == lib.ActionRelease {
		table[k] = false
	}

	for _, handler := range handlers {
		handler(k, a)
	}
}

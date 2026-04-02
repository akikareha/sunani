package key

import (
	"tea.kareha.org/loom/sunani/lib"
)

type EventHandler func(key lib.Key, action lib.Action)

var table = make([]bool, lib.KeyCount)
var eventHandlers = make([]EventHandler, 0)

//export sunani_key_init
func keyInit() {}

//export sunani_key_event
func keyEvent(key uint32, action uint32) {
	k := lib.Key(key)
	a := lib.Action(action)

	if a == lib.ActionPress {
		table[k] = true
	} else if a == lib.ActionRelease {
		table[k] = false
	}

	for _, handler := range eventHandlers {
		handler(k, a)
	}
}

func AddEventHandler(handler EventHandler) {
	eventHandlers = append(eventHandlers, handler)
}

func IsDown(key lib.Key) bool {
	// key cannot be negative since lib.Key is uint16.
	if key >= lib.KeyCount {
		return false
	}
	return table[key]
}

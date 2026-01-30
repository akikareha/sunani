package kb

import (
	"github.com/akikareha/sunani/lib"
)

func Char(k lib.Key) string {
	switch k {

	// Letters
	case lib.KeyA:
		return "a"
	case lib.KeyB:
		return "b"
	case lib.KeyC:
		return "c"
	case lib.KeyD:
		return "d"
	case lib.KeyE:
		return "e"
	case lib.KeyF:
		return "f"
	case lib.KeyG:
		return "g"
	case lib.KeyH:
		return "h"
	case lib.KeyI:
		return "i"
	case lib.KeyJ:
		return "j"
	case lib.KeyK:
		return "k"
	case lib.KeyL:
		return "l"
	case lib.KeyM:
		return "m"
	case lib.KeyN:
		return "n"
	case lib.KeyO:
		return "o"
	case lib.KeyP:
		return "p"
	case lib.KeyQ:
		return "q"
	case lib.KeyR:
		return "r"
	case lib.KeyS:
		return "s"
	case lib.KeyT:
		return "t"
	case lib.KeyU:
		return "u"
	case lib.KeyV:
		return "v"
	case lib.KeyW:
		return "w"
	case lib.KeyX:
		return "x"
	case lib.KeyY:
		return "y"
	case lib.KeyZ:
		return "z"

	// Digits
	case lib.Key0:
		return "0"
	case lib.Key1:
		return "1"
	case lib.Key2:
		return "2"
	case lib.Key3:
		return "3"
	case lib.Key4:
		return "4"
	case lib.Key5:
		return "5"
	case lib.Key6:
		return "6"
	case lib.Key7:
		return "7"
	case lib.Key8:
		return "8"
	case lib.Key9:
		return "9"

	// Control
	case lib.KeyEscape:
		return "\x7f"
	case lib.KeyEnter:
		return "\n"
	case lib.KeySpace:
		return " "
	case lib.KeyTab:
		return "\t"
	case lib.KeyBackspace:
		return "\b"

	// Arrows
	case lib.KeyUp:
		return ""
	case lib.KeyDown:
		return ""
	case lib.KeyLeft:
		return ""
	case lib.KeyRight:
		return ""

	// Unknown
	default:
		return ""
	}
}

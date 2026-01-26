package lib

type Key uint16

const (
	KeyUnknown Key = iota

	// Letters
	KeyA
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
	KeyG
	KeyH
	KeyI
	KeyJ
	KeyK
	KeyL
	KeyM
	KeyN
	KeyO
	KeyP
	KeyQ
	KeyR
	KeyS
	KeyT
	KeyU
	KeyV
	KeyW
	KeyX
	KeyY
	KeyZ

	// Digits
	Key0
	Key1
	Key2
	Key3
	Key4
	Key5
	Key6
	Key7
	Key8
	Key9

	// Control
	KeyEscape
	KeyEnter
	KeySpace
	KeyTab
	KeyBackspace

	// Arrows
	KeyUp
	KeyDown
	KeyLeft
	KeyRight

	// Count
	KeyCount
)

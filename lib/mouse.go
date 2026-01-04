package lib

type Mouse uint16

const (
	MouseUnknown Mouse = iota

	MouseLeft
	MouseRight
	MouseMiddle
)

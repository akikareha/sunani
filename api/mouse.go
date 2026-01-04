package api

type Mouse uint16

const (
	MouseUnknown Mouse = iota

	MouseLeft
	MouseRight
	MouseMiddle
)

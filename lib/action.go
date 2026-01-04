package lib

type Action uint16

const (
	ActionUnknown Action = iota

	ActionPress
	ActionRelease
)

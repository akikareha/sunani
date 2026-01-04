package api

type MouseButton uint16

const (
	MouseButtonUnknown MouseButton = iota

	MouseLeft
	MouseMiddle
	MouseRight
)

type MouseAction uint16

const (
	MouseActionUnknown MouseAction = iota

	MouseMove
	MousePress
	MouseRelease
	MouseWheel
)

type MouseEvent struct {
	Action         MouseAction
	X, Y           float64
	DX, DY         float64
	WheelX, WheelY float64
	Button         MouseButton
}

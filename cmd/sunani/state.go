package main

type Draw int

const (
	DrawNone = iota
	DrawCanvas
	DrawFB
)

type State struct {
	canvas *Canvas
	fb     *FB

	draw Draw
}

func NewState(canvas *Canvas, fb *FB) *State {
	return &State{
		canvas: canvas,
		fb:     fb,
	}
}

func (state *State) EnsureNone() {
	state.draw = DrawNone
}

func (state *State) EnsureCanvas() {
	if state.draw == DrawCanvas {
		return
	}
	if state.canvas == nil {
		return
	}
	state.canvas.Begin()
	state.draw = DrawCanvas
}

func (state *State) EnsureFB() {
	if state.draw == DrawFB {
		return
	}
	if state.fb == nil {
		return
	}
	state.fb.Begin()
	state.draw = DrawFB
}

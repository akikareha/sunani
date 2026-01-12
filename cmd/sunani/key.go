package main

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tetratelabs/wazero/api"
)

type Key struct {
	init  api.Function
	event api.Function

	window *glfw.Window
}

func NewKey() *Key {
	return &Key{}
}

func (k *Key) Preinit() {
	k.init = mod.ExportedFunction("sunani_key_init")
	k.event = mod.ExportedFunction("sunani_key_event")
}

func (k *Key) IsEnabled() bool {
	return k.init != nil
}

func (k *Key) Init(window *glfw.Window) {
	if !k.IsEnabled() {
		return
	}
	if window == nil {
		panic("window is nil")
	}
	k.window = window

	if k.init != nil {
		_, err := k.init.Call(ctx)
		if err != nil {
			die("sunani_key_init call failed:", err)
		}
	}

	if k.event != nil {
		window.SetKeyCallback(func(
			w *glfw.Window,
			key glfw.Key,
			scancode int,
			action glfw.Action,
			mods glfw.ModifierKey,
		) {
			kcode := mapGLFWKey(key)
			a := mapGLFWAction(action)

			_, err := k.event.Call(ctx, uint64(kcode), uint64(a))
			if err != nil {
				die("sunani_key_event call failed:", err)
			}
		})
	}
}

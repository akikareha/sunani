package main

import (
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tetratelabs/wazero/api"
)

type Key struct {
	window *glfw.Window

	init  api.Function
	event api.Function
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

	k.window = window

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
			log.Fatalln("key event call failed:", err)
		}
	})
}

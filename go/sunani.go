package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	//"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

var (
	ctx = context.Background()
	mod api.Module
	mem api.Memory
)

func init() {
	// Main thread is required for OpenGL.
	runtime.LockOSThread()
}

func mustRead(path string) []byte {
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	return b
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(
			os.Stderr,
			//"usage: %s program.wasm [args...]\n",
			"usage: %s program.wasm\n",
			os.Args[0],
		)
		os.Exit(1)
	}

	wasmPath := os.Args[1]
	//wasmArgs := os.Args[1:] // argv for WASI

	// --- wazero init ---
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)

	// for Go(wasip1) runtime
	//wasi_snapshot_preview1.MustInstantiate(ctx, r)

	canvas := NewCanvas()
	fb := NewFB()

	// Host module "canvas": expose Clear/SetColor/Line/Rect
	_, err := r.NewHostModuleBuilder("sunani").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context, r, g, b, a float32) {
			canvas.Clear(r, g, b, a)
		}).Export("canvas_clear").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context, r, g, b, a float32) {
			canvas.SetColor(r, g, b, a)
		}).Export("canvas_color").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context, x1, y1, x2, y2 float32) {
			canvas.Line(x1, y1, x2, y2)
		}).Export("canvas_line").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context, x, y, w, h float32, fill uint32) {
			canvas.Rect(x, y, w, h, fill != 0)
		}).Export("canvas_rect").
		Instantiate(ctx)
	if err != nil {
		log.Fatalln("instantiate host sunani module:", err)
	}

	wasmBytes := mustRead(wasmPath)

	// Don't call _start automatically.
	// draw() must be called from each frame.
	mod, err = r.InstantiateWithConfig(ctx, wasmBytes, wazero.NewModuleConfig()/*.WithArgs(wasmArgs...)*/.WithStartFunctions())
	if err != nil {
		log.Fatalln("instantiate guest:", err)
	}

	canvas.Preinit()
	fb.Preinit()
	var window *glfw.Window
	if canvas.IsEnabled() || fb.IsEnabled() {
		// --- GLFW/GL init ---
		if err := glfw.Init(); err != nil {
			log.Fatalln("glfw init failed:", err)
		}
		defer glfw.Terminate()

		glfw.WindowHint(glfw.ContextVersionMajor, 2)
		glfw.WindowHint(glfw.ContextVersionMinor, 1)

		if err := gl.Init(); err != nil {
			log.Fatalln("gl init failed:", err)
		}

		window, err = glfw.CreateWindow(512, 512, "wazero + GLFW", nil, nil)
		if err != nil {
			log.Fatalln(err)
		}
		window.MakeContextCurrent()

		canvas.Init(window)
		fb.Init(window)
	}

	if fb.IsEnabled() {
		mem = mod.ExportedMemory("memory")
		if mem == nil {
			log.Fatal("wasm exported memory not found")
		}
	}

	if window != nil {
		gl.Enable(gl.BLEND)
		gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
		gl.Enable(gl.TEXTURE_2D)

		// font orientation
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

		gl.PixelStorei(gl.UNPACK_ALIGNMENT, 1)

		// --- main loop ---
		for !window.ShouldClose() {
			gl.Clear(gl.COLOR_BUFFER_BIT)

			canvas.Begin()
			canvas.Draw()

			fb.Begin()
			fb.Draw()

			window.SwapBuffers()
			glfw.PollEvents()
		}
	}

	_ = mod.Close(ctx)
}

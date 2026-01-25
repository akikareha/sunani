package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

var (
	ctx = context.Background()
	mod api.Module
)

func init() {
	// Main thread is required for OpenGL.
	runtime.LockOSThread()
}

func mustRead(path string) []byte {
	b, err := os.ReadFile(path)
	if err != nil {
		die(err)
	}
	return b
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(
			os.Stderr,
			"usage: %s program.wasm [args...]\n",
			os.Args[0],
		)
		os.Exit(1)
	}

	wasmPath := os.Args[1]
	wasmArgs := os.Args[1:] // argv for WASI

	// --- wazero init ---
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)

	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	runtime := NewRuntime()
	console := NewConsole()
	canvas := NewCanvas()
	fb := NewFB()
	key := NewKey()
	mouse := NewMouse()

	_, err := r.NewHostModuleBuilder("sunani").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context) int64 {
			return time.Now().UnixMilli()
		}).Export("std.now").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context) {
			runtime.Halt()
		}).Export("runtime.halt").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context, ptr uint32, length uint32) {
			runtime.Title(ptr, length)
		}).Export("runtime.title").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context, enabled uint32) {
			runtime.Cursor(enabled)
		}).Export("runtime.cursor").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context, r, g, b, a uint32) {
			runtime.Clear(r, g, b, a)
		}).Export("runtime.clear").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context, ptr uint32, length uint32) {
			console.Params(ptr, int(length))
		}).Export("console.params").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context, ptr uint32, length uint32) {
			console.Put(ptr, length)
		}).Export("console.put").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context) {
			console.Wait()
		}).Export("console.wait").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context) {
			console.Leave()
		}).Export("console.leave").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context) {
			canvas.Begin()
		}).Export("canvas.begin").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context, r, g, b, a uint32) {
			canvas.Color(r, g, b, a)
		}).Export("canvas.color").
		NewFunctionBuilder().
		WithFunc(func(
			ctx context.Context,
			x1, y1 uint32,
			x2, y2 uint32,
		) {
			canvas.Line(x1, y1, x2, y2)
		}).Export("canvas.line").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context, x, y uint32, w, h uint32) {
			canvas.Rect(x, y, w, h)
		}).Export("canvas.rect").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context, x, y uint32, w, h uint32) {
			canvas.FillRect(x, y, w, h)
		}).Export("canvas.fill_rect").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context) {
			canvas.Path()
		}).Export("canvas.path").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context, x, y uint32) {
			canvas.Vertex(x, y)
		}).Export("canvas.vertex").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context) {
			canvas.Polygon()
		}).Export("canvas.polygon").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context) {
			canvas.FillPolygon()
		}).Export("canvas.fill_polygon").
		NewFunctionBuilder().
		WithFunc(func(
			ctx context.Context,
			ptr uint32,
			width, height uint32,
		) {
			fb.Params(ptr, int(width), int(height))
		}).Export("fb.params").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context) {
			fb.Paint()
		}).Export("fb.paint").
		Instantiate(ctx)
	if err != nil {
		die("instantiate host sunani module:", err)
	}

	wasmBytes := mustRead(wasmPath)

	config := wazero.NewModuleConfig().
		WithArgs(wasmArgs...).
		WithStdout(os.Stdout).
		WithStderr(os.Stderr).
		WithStdin(os.Stdin)

	// Don't call _start automatically.
	// runtime.frame must be called from each frame.
	mod, err = r.InstantiateWithConfig(ctx, wasmBytes, config)
	if err != nil {
		die("instantiate guest:", err)
	}

	runtime.Preinit()
	console.Preinit()
	canvas.Preinit()
	fb.Preinit()
	key.Preinit()
	mouse.Preinit()

	console.Init()
	var window *glfw.Window
	if canvas.IsEnabled() || fb.IsEnabled() {
		// --- GLFW/GL init ---
		if err := glfw.Init(); err != nil {
			die("glfw init failed:", err)
		}
		defer glfw.Terminate()

		glfw.WindowHint(glfw.ContextVersionMajor, 2)
		glfw.WindowHint(glfw.ContextVersionMinor, 1)

		if err := gl.Init(); err != nil {
			die("gl init failed:", err)
		}

		glfw.WindowHint(glfw.Samples, 4)
		window, err = glfw.CreateWindow(512, 512, "Sunani", nil, nil)
		if err != nil {
			die(err)
		}
		window.MakeContextCurrent()

		runtime.Init(window)
		canvas.Init(window)
		fb.Init(window)
		key.Init(window)
		mouse.Init(window)
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
			runtime.Frame()

			window.SwapBuffers()
			glfw.PollEvents()
		}
	}

	_ = mod.Close(ctx)
}

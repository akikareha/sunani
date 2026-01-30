package main

import (
	"context"
	"fmt"
	"os"
	"runtime"

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

	std := NewStd()
	console := NewConsole()
	screen := NewScreen()
	canvas := NewCanvas()
	fb := NewFB()
	key := NewKey()
	mouse := NewMouse()

	state := NewState(canvas, fb)

	_, err := r.NewHostModuleBuilder("sunani").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context) int64 {
			return std.Now()
		}).Export("std.now").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context, ptr uint32, length uint32) {
			console.Params(ptr, length)
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
			screen.Halt()
		}).Export("screen.halt").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context, ptr uint32, length uint32) {
			screen.Title(ptr, length)
		}).Export("screen.title").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context, visible uint32) {
			screen.Cursor(visible != 0)
		}).Export("screen.cursor").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context, r, g, b, a uint32) {
			screen.Clear(int(r), int(g), int(b), int(a))
		}).Export("screen.clear").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context, r, g, b, a uint32) {
			state.EnsureCanvas()
			canvas.Color(int(r), int(g), int(b), int(a))
		}).Export("canvas.color").
		NewFunctionBuilder().
		WithFunc(func(
			ctx context.Context,
			x1, y1 int32,
			x2, y2 int32,
		) {
			state.EnsureCanvas()
			canvas.Line(int(x1), int(y1), int(x2), int(y2))
		}).Export("canvas.line").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context, x, y int32, w, h int32) {
			state.EnsureCanvas()
			canvas.Rect(int(x), int(y), int(w), int(h))
		}).Export("canvas.rect").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context, x, y int32, w, h int32) {
			state.EnsureCanvas()
			canvas.FillRect(int(x), int(y), int(w), int(h))
		}).Export("canvas.fill_rect").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context) {
			state.EnsureCanvas()
			canvas.Path()
		}).Export("canvas.path").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context, x, y int32) {
			state.EnsureCanvas()
			canvas.Vertex(int(x), int(y))
		}).Export("canvas.vertex").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context) {
			state.EnsureCanvas()
			canvas.Polygon()
		}).Export("canvas.polygon").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context) {
			state.EnsureCanvas()
			canvas.FillPolygon()
		}).Export("canvas.fill_polygon").
		NewFunctionBuilder().
		WithFunc(func(
			ctx context.Context,
			ptr uint32,
			width, height int32,
		) {
			fb.Params(ptr, int(width), int(height))
		}).Export("fb.params").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context) {
			state.EnsureFB()
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
	// screen.frame must be called from each frame.
	mod, err = r.InstantiateWithConfig(ctx, wasmBytes, config)
	if err != nil {
		die("instantiate guest:", err)
	}

	std.Preinit()
	console.Preinit()
	screen.Preinit()
	canvas.Preinit()
	fb.Preinit()
	key.Preinit()
	mouse.Preinit()

	init := mod.ExportedFunction("sunani_init")
	if init != nil {
		init.Call(ctx)
	}

	std.Init()
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

		screen.Init(window)
		canvas.Init(window)
		fb.Init(window)
		key.Init(window)
		mouse.Init(window)
	}

	start := mod.ExportedFunction("sunani_start")
	if start != nil {
		start.Call(ctx)
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
			state.EnsureNone()

			screen.Frame()

			window.SwapBuffers()
			glfw.PollEvents()
		}
	}

	_ = mod.Close(ctx)
}

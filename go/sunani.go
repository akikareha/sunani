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
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
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
		fmt.Fprintf(os.Stderr, "usage: %s program.wasm [args...]\n", os.Args[0])
		os.Exit(1)
	}

	wasmPath := os.Args[1]
	wasmArgs := os.Args[1:] // argv for WASI

	// --- GLFW/GL init ---
	if err := glfw.Init(); err != nil {
		log.Fatalln("glfw init failed:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	const W, H = 512, 512
	window, err := glfw.CreateWindow(W, H, "wazero + GLFW", nil, nil)
	if err != nil {
		log.Fatalln(err)
	}
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		log.Fatalln("gl init failed:", err)
	}

	canvas := NewCanvas(W, H)
	canvas.Init()

	// --- wazero init ---
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)

	// for Go(wasip1) runtime
	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	// Host module "canvas": expose Clear/SetColor/Line/Rect
	_, err = r.NewHostModuleBuilder("canvas").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context, r, g, b, a float32) {
			canvas.Clear(r, g, b, a)
		}).Export("clear").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context, r, g, b, a float32) {
			canvas.SetColor(r, g, b, a)
		}).Export("setColor").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context, x1, y1, x2, y2 float32) {
			canvas.Line(x1, y1, x2, y2)
		}).Export("line").
		NewFunctionBuilder().
		WithFunc(func(ctx context.Context, x, y, w, h float32, fill uint32) {
			canvas.Rect(x, y, w, h, fill != 0)
		}).Export("rect").
		Instantiate(ctx)
	if err != nil {
		log.Fatalln("instantiate host canvas module:", err)
	}

	wasmBytes := mustRead(wasmPath)

	// Don't call _start automatically.
	// draw() must be called from each frame.
	mod, err = r.InstantiateWithConfig(ctx, wasmBytes, wazero.NewModuleConfig().WithArgs(wasmArgs...).WithStartFunctions())
	if err != nil {
		log.Fatalln("instantiate guest:", err)
	}
	drawFn := mod.ExportedFunction("draw")
	drawFnFound := drawFn != nil

	mem = mod.ExportedMemory("memory")
	if mem == nil {
		log.Fatal("wasm exported memory not found")
	}

	fbFound := isFnExist("FBPtr")
	var fbPtr uint32
	var fbW int
	var fbH int
	var fbSize int
	var strPtr uint32
	if fbFound {
		fbPtr = callU32("FBPtr")
		fbW = int(callU32("FBW"))
		fbH = int(callU32("FBH"))
		fbSize = fbW * fbH * 4
		strPtr = callU32("TextBufPtr")
	}

	if fbFound {
		gl.MatrixMode(gl.PROJECTION)
		gl.LoadIdentity()
		gl.Ortho(-1, 1, -1, 1, -1, 1)

		gl.MatrixMode(gl.MODELVIEW)
		gl.LoadIdentity()
	}

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.TEXTURE_2D)

	var tex uint32
	gl.GenTextures(1, &tex)
	gl.BindTexture(gl.TEXTURE_2D, tex)

	// font orientation
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	if fbFound {
		gl.TexImage2D(
			gl.TEXTURE_2D,
			0,
			gl.RGBA,
			int32(fbW),
			int32(fbH),
			0,
			gl.RGBA,
			gl.UNSIGNED_BYTE,
			nil,
		)
	}

	gl.PixelStorei(gl.UNPACK_ALIGNMENT, 1)

	// --- main loop ---
	for !window.ShouldClose() {
		if drawFnFound {
			_, callErr := drawFn.Call(ctx)
			if callErr != nil {
				log.Fatalln("draw call failed:", callErr)
			}
		}

		if fbFound {
			msg := "Hello, Sunani!"
			call("Clear", 0, 0, 0, 255)
			writeStringToWasm(mem, strPtr, msg)
			call("DrawText", 16, 32, uint64(strPtr), uint64(len(msg)))

			pix, ok := mem.Read(fbPtr, uint32(fbSize))
			if !ok {
				panic("mem.Read failed")
			}

			gl.Clear(gl.COLOR_BUFFER_BIT)
			fw, fh := window.GetFramebufferSize()
			scale := min(fw/fbW, fh/fbH)
			drawW := fbW * scale
			drawH := fbH * scale

			ox := (fw - drawW) / 2
			oy := (fh - drawH) / 2

			gl.Viewport(int32(ox), int32(oy), int32(drawW), int32(drawH))

			gl.BindTexture(gl.TEXTURE_2D, tex)
			gl.TexSubImage2D(
				gl.TEXTURE_2D,
				0,
				0, 0,
				int32(fbW), int32(fbH),
				gl.RGBA,
				gl.UNSIGNED_BYTE,
				gl.Ptr(pix),
			)

			gl.Begin(gl.QUADS)
			gl.TexCoord2f(0, 1)
			gl.Vertex2f(-1, -1)
			gl.TexCoord2f(1, 1)
			gl.Vertex2f(1, -1)
			gl.TexCoord2f(1, 0)
			gl.Vertex2f(1, 1)
			gl.TexCoord2f(0, 0)
			gl.Vertex2f(-1, 1)
			gl.End()
		}

		window.SwapBuffers()
		glfw.PollEvents()
	}

	_ = mod.Close(ctx)
	//_ = api.ErrInvalid
}

func isFnExist(name string) bool {
	fn := mod.ExportedFunction(name)
	return fn != nil
}

func callU32(name string, args ...uint64) uint32 {
	fn := mod.ExportedFunction(name)
	if fn == nil {
		log.Fatalf("function not found: %s", name)
	}
	res, err := fn.Call(ctx, args...)
	if err != nil {
		log.Fatal(err)
	}
	return uint32(res[0])
}

func call(name string, args ...uint64) {
	fn := mod.ExportedFunction(name)
	if fn == nil {
		log.Fatalf("function not found: %s", name)
	}
	_, err := fn.Call(ctx, args...)
	if err != nil {
		log.Fatal(err)
	}
}

func writeStringToWasm(mem api.Memory, ptr uint32, s string) {
	b := []byte(s)
	ok := mem.Write(ptr, b)
	if !ok {
		log.Fatal("mem.Write failed")
	}
}

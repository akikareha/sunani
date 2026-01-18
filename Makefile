.PHONY: all build docs clean fmt

all: build docs

build: resources/fonts/ascii8x8.go
	go build -o sunani ./cmd/sunani

docs: build/demo.wasm build/hello.wasm build/echo.wasm build/hello-canvas.wasm build/hello-fb.wasm build/key.wasm
	cp web/* docs
	cp build/*.wasm docs

build/demo.wasm: app-go/demo/*.go
	tinygo build -target=wasm-unknown -scheduler=none -gc=conservative -o build/demo.wasm ./app-go/demo

./resources/fonts/ascii8x8.go: resources/fonts/ascii8x8.png png2rgba
	./png2rgba -in resources/fonts/ascii8x8.png -out resources/fonts/ascii8x8.go -pkg fonts

./png2rgba: tools/png2rgba/*.go
	go build -o png2rgba ./tools/png2rgba

build/hello.wasm: app-go/hello/*.go
	tinygo build -target=wasm-unknown -scheduler=none -gc=conservative -o build/hello.wasm ./app-go/hello

build/echo.wasm: app-go/echo/*.go
	tinygo build -target=wasm-unknown -scheduler=none -gc=conservative -o build/echo.wasm ./app-go/echo

build/hello-canvas.wasm: app-go/hello-canvas/*.go
	tinygo build -target=wasm-unknown -scheduler=none -gc=conservative -o build/hello-canvas.wasm ./app-go/hello-canvas

build/hello-fb.wasm: app-go/hello-fb/*.go
	tinygo build -target=wasm-unknown -scheduler=none -gc=conservative -o build/hello-fb.wasm ./app-go/hello-fb

build/key.wasm: app-go/key/*.go
	tinygo build -target=wasm-unknown -scheduler=none -gc=conservative -o build/key.wasm ./app-go/key

clean:
	rm -f sunani png2rgba build/*.wasm

fmt:
	go fmt ./...

key: sunani build/key.wasm
	./sunani build/key.wasm

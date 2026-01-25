.PHONY: all prepare build docs clean fmt app-go

all: prepare build docs

prepare:
	mkdir -p build/app-go
	mkdir -p build/app-wat
	mkdir -p build/app-grain

build: resources/fonts/ascii8x8.go
	go build -o sunani ./cmd/sunani

docs: app-go app-wat app-grain
	rm -r docs/*
	cp -r web/* docs
	cp -r build/* docs

./resources/fonts/ascii8x8.go: resources/fonts/ascii8x8.png png2rgba
	./png2rgba -in resources/fonts/ascii8x8.png \
	-out resources/fonts/ascii8x8.go -pkg fonts

./png2rgba: tools/png2rgba/*.go
	go build -o png2rgba ./tools/png2rgba

#
# app-go
#

app-go: build/app-go/demo.wasm \
        build/app-go/hello.wasm \
        build/app-go/echo.wasm \
        build/app-go/hello-canvas.wasm \
        build/app-go/hello-fb.wasm \
        build/app-go/key.wasm \
        build/app-go/system.wasm

build/app-go/demo.wasm: app-go/demo/*.go
	tinygo build -target=wasm-unknown -scheduler=none -gc=conservative \
	-o build/app-go/demo.wasm ./app-go/demo

build/app-go/hello.wasm: app-go/hello/*.go
	tinygo build -target=wasm-unknown -scheduler=none -gc=conservative \
	-o build/app-go/hello.wasm ./app-go/hello

build/app-go/echo.wasm: app-go/echo/*.go
	tinygo build -target=wasm-unknown -scheduler=none -gc=conservative \
	-o build/app-go/echo.wasm ./app-go/echo

build/app-go/hello-canvas.wasm: app-go/hello-canvas/*.go
	tinygo build -target=wasm-unknown -scheduler=none -gc=conservative \
	-o build/app-go/hello-canvas.wasm ./app-go/hello-canvas

build/app-go/hello-fb.wasm: app-go/hello-fb/*.go
	tinygo build -target=wasm-unknown -scheduler=none -gc=conservative \
	-o build/app-go/hello-fb.wasm ./app-go/hello-fb

build/app-go/key.wasm: app-go/key/*.go
	tinygo build -target=wasm-unknown -scheduler=none -gc=conservative \
	-o build/app-go/key.wasm ./app-go/key

build/app-go/system.wasm: app-go/system/*.go
	tinygo build -target=wasm-unknown -scheduler=none -gc=conservative \
	-o build/app-go/system.wasm ./app-go/system

#
# app-wat
#

app-wat: build/app-wat/hello.wasm

build/app-wat/hello.wasm: app-wat/hello/*.wat
	wat2wasm -o build/app-wat/hello.wasm app-wat/hello/hello.wat

#
# app-grain
#

app-grain: build/app-grain/hello.wasm

build/app-grain/hello.wasm: app-grain/hello/*.gr
	grain compile app-grain/hello/hello.gr --no-wasm-tail-call \
	-o build/app-grain/hello.wasm

#
#
#

clean:
	rm -f sunani png2rgba
	rm -rf build/*

fmt:
	go fmt ./...

#
#
#

dev: sunani build/app-go/system.wasm
	./sunani build/app-go/system.wasm

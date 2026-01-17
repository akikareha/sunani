.PHONY: all build clean fmt

all: build web

build: resources/fonts/ascii8x8.go
	go build -o sunani ./cmd/sunani

web: build/demo.wasm
	cp web/* docs
	cp build/demo.wasm docs

build/demo.wasm: app-go/demo/*.go
	tinygo build -target=wasm-unknown -scheduler=none -gc=conservative -o build/demo.wasm ./app-go/demo

./resources/fonts/ascii8x8.go: resources/fonts/ascii8x8.png png2rgba
	./png2rgba -in resources/fonts/ascii8x8.png -out resources/fonts/ascii8x8.go -pkg fonts

./png2rgba: tools/png2rgba/*.go
	go build -o png2rgba ./tools/png2rgba

clean:
	rm -f sunani png2rgba build/demo.wasm

fmt:
	go fmt ./...

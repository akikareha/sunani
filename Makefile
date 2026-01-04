all: build

build: ./resources/fonts/ascii8x8.go
	go build -o sunani ./cmd/sunani
	tinygo build -target=wasm-unknown -scheduler=none -gc=conservative -o demo.wasm ./demo

./resources/fonts/ascii8x8.go: ./resources/fonts/ascii8x8.png ./png2rgba
	./png2rgba -in ./resources/fonts/ascii8x8.png -out ./resources/fonts/ascii8x8.go -pkg fonts

./png2rgba:
	go build -o png2rgba ./tools/png2rgba

clean:
	rm -f sunani png2rgba demo.wasm

fmt:
	go fmt ./...

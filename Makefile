all: build

build:
	go build -o sunani ./cmd/sunani
	go build -o png2rgba ./tools/png2rgba
	./png2rgba -in ./resources/fonts/ascii8x8.png -out ./demo/ascii8x8.go -pkg main
	tinygo build -target=wasm-unknown -scheduler=none -o demo.wasm ./demo

clean:
	rm -f sunani png2rgba demo.wasm
	rm -f demo/ascii8x8.go

fmt:
	go fmt ./...

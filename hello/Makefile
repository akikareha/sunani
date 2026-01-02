all: hello.wasm

hello.wasm: hello.go
	cat hello.go font_ascii8x8._go > all.go
	tinygo build -target=wasm-unknown -scheduler=none -o hello.wasm all.go

clean:
	rm -f hello.wasm
	rm -f all.go
	rm -f font_ascii8x8._go

fmt:
	go fmt .

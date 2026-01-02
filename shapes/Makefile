all: shapes.wasm

shapes.wasm: shapes.go
	tinygo build -target=wasm-unknown -scheduler=none -o shapes.wasm shapes.go

clean:
	rm -f shapes.wasm

fmt:
	go fmt .

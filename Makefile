all:
	make -C tools
	make -C go
	make -C shapes
	./tools/png2rgba/png2rgba -in ./assets/font_ascii8x8.png -out ./hello/font_ascii8x8._go
	make -C hello

clean:
	make -C tools clean
	make -C go clean
	make -C shapes clean
	make -C hello clean

fmt:
	make -C tools fmt
	make -C go fmt
	make -C shapes fmt
	make -C hello fmt

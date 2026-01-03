package main

import (
	"log"
)

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

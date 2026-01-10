package main

import (
	"unsafe"
)

const consoleBufferSize = 4096

var consoleBuffer = make([]byte, consoleBufferSize)

func consoleBufferPtr() uint32 { return uint32(uintptr(unsafe.Pointer(&consoleBuffer[0]))) }

func consoleBufferSize_() uint32 { return consoleBufferSize }

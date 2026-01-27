package main

import (
	"bufio"
	"os"

	"github.com/tetratelabs/wazero/api"
)

type Console struct {
	init api.Function
	get  api.Function

	ptr       uint32
	length    uint32
	paramsSet bool

	done chan struct{}
}

func NewConsole() *Console {
	return &Console{}
}

func (con *Console) Preinit() {
	con.init = mod.ExportedFunction("sunani_console_init")
	con.get = mod.ExportedFunction("sunani_console_get")
}

func (con *Console) IsEnabled() bool {
	return con.init != nil
}

func (con *Console) Init() {
	if !con.IsEnabled() {
		return
	}

	if con.init != nil {
		_, err := con.init.Call(ctx)
		if err != nil {
			die("sunani_console_init call failed:", err)
		}
	}
}

func (con *Console) Params(ptr uint32, length uint32) {
	if !con.IsEnabled() {
		errlog("sunani console.params was called, but Console API is not enabled.\nExport snunani_console_init to enable this API.")
		return
	}

	if con.get == nil {
		errlog("sunani console.params was called, but console input is not enabled.\nExport snunani_console_get to enable this function.")
		return
	}

	con.ptr = ptr
	con.length = length
	con.paramsSet = true

	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		for scanner.Scan() {
			bytes := scanner.Bytes()

			if len(bytes) > int(con.length) {
				errlog("Console input buffer overflowed")
				bytes = bytes[:con.length]
			}

			mem := mod.Memory()
			ok := mem.Write(con.ptr, bytes)
			if !ok {
				errlog("mem.Write failed")
				break
			}
			_, err := con.get.Call(ctx, uint64(con.ptr), uint64(len(bytes)))
			if err != nil {
				die("sunani_console_get call failed:", err)
			}
		}
		if err := scanner.Err(); err != nil {
			die("Failed to scan line:", err)
		}
	}()
}

func (con *Console) Put(ptr uint32, length uint32) {
	if !con.IsEnabled() {
		errlog("sunani console.put was called, but Console API is not enabled.\nExport snunani_console_init to enable this API.")
		return
	}

	if length < 1 {
		return
	}
	mem := mod.Memory()
	buf, ok := mem.Read(ptr, length)
	if !ok {
		errlog("mem.Read failed")
		return
	}

	os.Stdout.Write(buf)
}

func (con *Console) Wait() {
	con.done = make(chan struct{})
	<-con.done
}

func (con *Console) Leave() {
	close(con.done)
}

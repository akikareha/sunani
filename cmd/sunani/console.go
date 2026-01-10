package main

import (
	"log"
	"os"

	"github.com/tetratelabs/wazero/api"
)

type Console struct {
	init api.Function
}

func NewConsole() *Console {
	return &Console{}
}

func (con *Console) Preinit() {
	con.init = mod.ExportedFunction("sunani_console_init")
}

func (con *Console) IsEnabled() bool {
	return con.init != nil
}

func (con *Console) Init() {
	if !con.IsEnabled() {
		return
	}

	_, err := con.init.Call(ctx)
	if err != nil {
		log.Fatalln("console init call failed:", err)
	}
}

func (con *Console) Put(ptr uint32, length uint32) {
	if !con.IsEnabled() {
		return
	}

	mem := mod.Memory()
	buf, ok := mem.Read(ptr, length)
	if !ok {
		return
	}
	s := string(buf)

	os.Stdout.WriteString(s)
}

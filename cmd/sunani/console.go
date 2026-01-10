package main

import (
	"bufio"
	"log"
	"os"

	"github.com/tetratelabs/wazero/api"
)

type Console struct {
	ptr    uint32
	length int

	paramsSet bool

	init api.Function
	get  api.Function
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

	_, err := con.init.Call(ctx)
	if err != nil {
		log.Fatalln("console init call failed:", err)
	}
}

func (con *Console) Params(ptr uint32, length int) {
	con.ptr = ptr
	con.length = length

	con.paramsSet = true

	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		for scanner.Scan() {
			line := scanner.Text()

			mem := mod.Memory()
			ok := mem.Write(con.ptr, []byte(line))
			if !ok {
				panic("mem.Write failed")
			}
			_, err := con.get.Call(ctx, uint64(con.ptr), uint64(len(line)))
			if err != nil {
				log.Fatalln("console init call failed:", err)
			}
		}
		if err := scanner.Err(); err != nil {
			// TODO handle error
		}
	}()
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

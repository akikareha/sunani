package main

import (
	"time"

	"github.com/tetratelabs/wazero/api"
)

type Std struct {
	init api.Function
}

func NewStd() *Std {
	return &Std{}
}

func (std *Std) Preinit() {
	std.init = mod.ExportedFunction("sunani_std_init")
}

func (std *Std) IsEnabled() bool {
	return std.init != nil
}

func (std *Std) Init() {
	if !std.IsEnabled() {
		return
	}

	if std.init != nil {
		_, err := std.init.Call(ctx)
		if err != nil {
			die("sunani_std_init call failed:", err)
		}
	}
}

func (std *Std) Now() int64 {
	if !std.IsEnabled() {
		return 0
	}

	return time.Now().UnixMilli()
}

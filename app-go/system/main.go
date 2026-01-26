package main

import (
	"github.com/akikareha/sunani/app-go/base/system"
)

//export sunani_init
func init() {
	system.Run()
}

package std

import (
	standard "github.com/akikareha/sunani/app-go/api/std"
)

//export sunani_std_init
func stdInit() {}

func Now() int64 {
	return standard.Now()
}

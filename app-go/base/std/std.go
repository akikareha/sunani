package std

import (
	standard "tea.kareha.org/loom/sunani/app-go/api/std"
)

//export sunani_std_init
func stdInit() {}

func Now() int64 {
	return standard.Now()
}

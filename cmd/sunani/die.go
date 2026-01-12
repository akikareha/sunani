package main

import (
	"log"
)

func die(v ...any) {
	log.Fatalln(v...)
}

func errlog(v ...any) {
	log.Println(v...)
}

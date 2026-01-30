package system

import (
	"fmt"
	"strings"

	"github.com/akikareha/sunani/app-go/base/console"
	"github.com/akikareha/sunani/app-go/base/screen"
)

func repl(line string) bool {
	args := strings.Split(line, " ")
	if len(args) < 1 {
		return false
	}
	cmd := args[0]

	if cmd == "quit" || cmd == "q" {
		console.Print("Bye!\n")
		screen.Halt()
		return false
	}
	if cmd == "bg" {
		if len(args) < 2 {
		} else {
			arg := args[1]
			if arg == "red" {
				screen.SetBg(255, 0, 0, 255)
			} else if arg == "green" {
				screen.SetBg(0, 255, 0, 255)
			} else if arg == "blue" {
				screen.SetBg(0, 0, 255, 255)
			} else if arg == "black" {
				screen.SetBg(0, 0, 0, 255)
			}
		}
		return true
	}
	reply := fmt.Sprintf("Unknown command: %s\n", line)
	console.Print(reply)
	return true
}

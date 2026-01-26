package system

import (
	"fmt"
	"strings"
)

func repl(line string) bool {
	args := strings.Split(line, " ")
	if len(args) < 1 {
		return false
	}
	cmd := args[0]

	if cmd == "quit" || cmd == "q" {
		Print("Bye!\n")
		Halt()
		return false
	}
	if cmd == "bg" {
		if len(args) < 2 {
		} else {
			arg := args[1]
			if arg == "red" {
				SetBg(255, 0, 0, 255)
			} else if arg == "green" {
				SetBg(0, 255, 0, 255)
			} else if arg == "blue" {
				SetBg(0, 0, 255, 255)
			} else if arg == "black" {
				SetBg(0, 0, 0, 255)
			}
		}
		return true
	}
	reply := fmt.Sprintf("Unknown command: %s\n", line)
	Print(reply)
	return true
}

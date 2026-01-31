package system

import (
	"fmt"

	"github.com/akikareha/sunani/app-go/base/repl"
	"github.com/akikareha/sunani/app-go/base/screen"
)

var REPL *repl.REPL

func initREPL() {
	REPL = repl.Default()

	REPL.AddEchoHandler(echoHandler)

	background := repl.NewCommand(
		"background",
		"bg",
		bgFn,
		`Set background color.
Try > bg blue
Available colors:
  black white
  red green blue
  cyan magenta yellow
`,
	)
	REPL.AddCommand(background)

	REPL.Init()
}

func echoHandler(r *repl.REPL, s string) {
	text.Add(s)
}

func bgFn(r *repl.REPL, args []string) bool {
	if len(args) < 2 {
		r.Print(fmt.Sprintf("Usage: %s COLOR_NAME\n", args[0]))
		r.Print("COLOR_NAME: black white red green blue cyan magenta yellow\n")
		return false
	}
	switch args[1] {
	case "black":
		screen.SetBg(0, 0, 0, 255)
	case "white":
		screen.SetBg(255, 255, 255, 255)
	case "red":
		screen.SetBg(255, 0, 0, 255)
	case "green":
		screen.SetBg(0, 255, 0, 255)
	case "blue":
		screen.SetBg(0, 0, 255, 255)
	case "cyan":
		screen.SetBg(0, 255, 255, 255)
	case "magenta":
		screen.SetBg(255, 0, 255, 255)
	case "yellow":
		screen.SetBg(255, 255, 0, 255)
	default:
		r.Print(fmt.Sprintf("Unknown color: %s\n", args[1]))
		bgFn(r, []string{args[0]}) // show usage
	}
	return false
}

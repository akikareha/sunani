package system

import (
	"fmt"

	"github.com/akikareha/sunani/app-go/base/color"
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
		r.Print("COLOR_NAME:")
		for _, name := range color.Names() {
			r.Print(" " + name)
		}
		r.Print("\n")
		return false
	}
	c, ok := color.ByName(args[1])
	if !ok {
		r.Print(fmt.Sprintf("Unknown color: %s\n", args[1]))
		bgFn(r, []string{args[0]}) // show usage
		return false
	}
	screen.SetBg(c)
	input.SetColor(c.ToFg().WithAlpha(192))
	text.SetColor(c.ToFg().WithAlpha(128))
	return false
}

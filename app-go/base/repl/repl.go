package repl

import (
	"fmt"
	"strings"

	"github.com/akikareha/sunani/app-go/base/console"
	"github.com/akikareha/sunani/app-go/base/screen"
)

type Command func(*REPL, []string) bool
type EchoHandler func(*REPL, string)

type REPL struct {
	greeting     string
	prompt       string
	commands     map[string]Command
	echoHandlers []EchoHandler
}

const greeting = `Hello, World!
This is Sunani REPL System.
Type h for help or q to quit.
`
const prompt = "> "

var commands = map[string]Command{
	"h":  hCommand,
	"q":  qCommand,
	"bg": bgCommand,
}

var Default = New(greeting, prompt, commands)

func New(
	greeting string,
	prompt string,
	commands map[string]Command,
) *REPL {
	cmds := make(map[string]Command, len(commands))
	for k, v := range commands {
		cmds[k] = v
	}
	return &REPL{
		greeting:     greeting,
		prompt:       prompt,
		commands:     cmds,
		echoHandlers: make([]EchoHandler, 0),
	}
}

func (r *REPL) Init() {
	console.AddInputHandler(r.consoleInputHandler)
	r.Print(r.greeting)
	r.Print(r.prompt)
}

func (r *REPL) AddEchoHandler(handler EchoHandler) {
	r.echoHandlers = append(r.echoHandlers, handler)
}

func (r *REPL) Print(s string) {
	console.Print(s)
	r.echo(s)
}

func (r *REPL) echo(s string) {
	for _, handler := range r.echoHandlers {
		handler(r, s)
	}
}

func (r *REPL) consoleInputHandler(line string) {
	r.echo(line + "\n")
	args := strings.Split(line, " ")
	c := args[0]
	command := r.commands[c]
	if command == nil {
		r.Print(fmt.Sprintf("Unknown command: %s\n", c))
	} else if command(r, args) {
		return
	}
	r.Print(r.prompt)
}

func (r *REPL) Input(line string) {
	console.Print(line + "\n")
	r.consoleInputHandler(line)
}

func hCommand(r *REPL, args []string) bool {
	r.Print("Commands:\n")
	r.Print("h q bg\n")
	return false
}

func qCommand(r *REPL, args []string) bool {
	r.Print("See you.\n")
	screen.Halt()
	return true
}

func bgCommand(r *REPL, args []string) bool {
	if len(args) < 2 {
		r.Print(fmt.Sprintf("Usage: %s COLOR_NAME\n", args[0]))
		r.Print("COLOR_NAME: black white red green blue cyan purple yellow\n")
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
	case "purple":
		screen.SetBg(255, 0, 255, 255)
	case "yellow":
		screen.SetBg(255, 255, 0, 255)
	default:
		r.Print(fmt.Sprintf("Unknown color: %s\n", args[1]))
		bgCommand(r, []string{args[0]}) // show usage
	}
	return false
}

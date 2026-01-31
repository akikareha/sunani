package repl

import (
	"fmt"
	"strings"

	"github.com/akikareha/sunani/app-go/base/console"
	"github.com/akikareha/sunani/app-go/base/screen"
)

type Fn func(*REPL, []string) bool

type Command struct {
	name     string
	shortcut string
	fn       Fn
	help     string
}

func NewCommand(
	name string,
	shortcut string,
	fn Fn,
	help string,
) *Command {
	return &Command{
		name:     name,
		shortcut: shortcut,
		fn:       fn,
		help:     help,
	}
}

type EchoHandler func(*REPL, string)

type REPL struct {
	greeting     string
	prompt       string
	commands     map[string]*Command
	shortcuts    map[string]*Command
	list         []*Command
	echoHandlers []EchoHandler
}

func Default() *REPL {
	r := New()

	r.SetGreeting(`Hello, World!
This is Sunani REPL System.
Type h for help or q to quit.
`)

	help := NewCommand(
		"help",
		"h",
		hFn,
		"Show help message.\n",
	)
	r.AddCommand(help)

	quit := NewCommand(
		"quit",
		"q",
		qFn,
		"Quit this system.\n",
	)
	r.AddCommand(quit)

	return r
}

func New() *REPL {
	r := REPL{
		greeting:     "Ready.",
		prompt:       "> ",
		commands:     make(map[string]*Command, 0),
		shortcuts:    make(map[string]*Command, 0),
		list:         make([]*Command, 0),
		echoHandlers: make([]EchoHandler, 0),
	}
	return &r
}

func (r *REPL) Init() {
	console.AddInputHandler(r.consoleInputHandler)
	r.Print(r.greeting)
	r.Print(r.prompt)
}

func (r *REPL) AddCommand(c *Command) {
	r.commands[c.name] = c
	r.shortcuts[c.shortcut] = c
	r.list = append(r.list, c)
}

func (r *REPL) AddEchoHandler(handler EchoHandler) {
	r.echoHandlers = append(r.echoHandlers, handler)
}

func (r *REPL) GetGreeting() string {
	return r.greeting
}

func (r *REPL) SetGreeting(s string) {
	r.greeting = s
}

func (r *REPL) GetPrompt() string {
	return r.prompt
}

func (r *REPL) SetPrompt(s string) {
	r.prompt = s
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
	name := args[0]
	c := r.commands[name]
	if c == nil {
		c = r.shortcuts[name]
	}
	if c == nil {
		r.Print(fmt.Sprintf("Unknown command: %s\n", name))
	} else if c.fn(r, args) {
		return
	}
	r.Print(r.prompt)
}

func (r *REPL) Input(line string) {
	console.Print(line + "\n")
	r.consoleInputHandler(line)
}

func hFn(r *REPL, args []string) bool {
	if len(args) < 2 {
		r.Print("List of commands:\n")
		for _, c := range r.list {
			r.Print(fmt.Sprintf("* %s (%s)\n", c.name, c.shortcut))
		}
		r.Print("Try > help [command]\n")
		r.Print("for more help for [command].\n")
		return false
	}
	name := args[1]
	c := r.commands[name]
	if c == nil {
		c = r.shortcuts[name]
	}
	if c == nil {
		r.Print(fmt.Sprintf("Help: Unknown command: %s\n", c))
		return false
	}
	r.Print(c.help)
	return false
}

func qFn(r *REPL, args []string) bool {
	r.Print("See you.\n")
	screen.Halt()
	return true
}

package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func NewGreetCommand() *GreetCommand {
	gc := &GreetCommand{
		fs: flag.NewFlagSet("greet", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.name, "name", "World", "name to greet")
	return gc
}

type GreetCommand struct {
	fs *flag.FlagSet

	name string
}

func (g *GreetCommand) Name() string {
	return g.fs.Name()
}

func (g *GreetCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *GreetCommand) Run() error {
	fmt.Println("Hello ", g.name)
	return nil
}

type Runner interface {
	Name() string
	Init(args []string) error
	Run() error
}

// root - to read the sub-command
func root(args []string) error {
	if len(args) < 1 {
		return errors.New("you must pass a sub-command")
	}

	cmds := []Runner{
		NewGreetCommand(),
	}

	subCommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subCommand {
			cmd.Init(os.Args[2:])
			return cmd.Run()
		}
	}

	return fmt.Errorf("unknown sub-command")
}

func main() {
	if err := root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

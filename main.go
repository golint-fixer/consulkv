package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
	"github.com/spiritloose/consulkv/command"
)

const (
	VERSION = "0.0.1"
)

func main() {
	c := cli.NewCLI("consulkv", VERSION)
	c.Args = os.Args[1:]
	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}
	c.Commands = map[string]cli.CommandFactory{
		"list": func() (cli.Command, error) {
			return &command.List{Ui: ui}, nil
		},
		"cat": func() (cli.Command, error) {
			return &command.Cat{Ui: ui}, nil
		},
		"put": func() (cli.Command, error) {
			return &command.Put{Ui: ui}, nil
		},
		"edit": func() (cli.Command, error) {
			return &command.Edit{Ui: ui}, nil
		},
		"delete": func() (cli.Command, error) {
			return &command.Delete{Ui: ui}, nil
		},
		"flags": func() (cli.Command, error) {
			return &command.Flags{Ui: ui}, nil
		},
	}
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}

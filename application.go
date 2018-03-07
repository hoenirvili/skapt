// Package skapt provides a tiny interface
// to create and manage your command line applications
package skapt

import (
	"fmt"

	"github.com/hoenirvili/skapt/command"
	"github.com/hoenirvili/skapt/context"
	"github.com/hoenirvili/skapt/flag"
	"github.com/hoenirvili/skapt/parser"
)

// Application will hold all the information for creating
// and parsing the command line
type Application struct {
	// Name of the Command line application
	Name string
	// Usage is the usage of the command line
	Usage string
	// Description holds the description of the command line
	Description string
	// Author holds the author name
	Author string
	// Version is the version of the application
	Version string
	// Flags holds list of the root command
	Flags flag.Flags
	// Commands holds list of subcommands
	Commands command.Commands
	// Handler is the root main handler
	Handler context.Handler
}

// validate if the holds valid information
// in order to be executed by Exec
func (a Application) validate() error {
	if a.Name == "" {
		return fmt.Errorf("skapt: Empty application name")
	}
	if a.Handler == nil {
		return fmt.Errorf("skapt: Empty application handler")
	}

	if a.Flags != nil {
		if err := a.Flags.Validate(); err != nil {
			return err
		}
	}

	if a.Commands != nil {
		if err := a.Commands.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// Exec executes the command line based on the args provided
func (a Application) Exec(args []string) error {
	if err := a.validate(); err != nil {
		return err
	}

	switch len(args) {
	case 0:
		return fmt.Errorf("skapt: No arguments to execute")
	}

	root := command.Command{
		Name:    args[0],
		Flags:   a.Flags,
		Handler: a.Handler,
	}

	// if default flags are not set, set them
	root.Flags.AppendDefault()
	a.Commands.AppendDefault()

	parser := parser.New(root, a.Commands)
	handlers, contexts, err := parser.Parse(args)
	if err != nil {
		return err
	}

	for key, handler := range handlers {
		context := contexts[key]
		if err := handler(context); err != nil {
			return err
		}
	}

	return nil
}

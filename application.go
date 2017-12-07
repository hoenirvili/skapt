// Package skapt provides a tiny interface
// to create and manage your command line aplications
package skapt

import (
	"fmt"

	"github.com/hoenirvili/skapt/argument"
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
	// Type holds the type of the root command value
	Type argument.Type
	// Handler is the root main handler
	Handler context.Handler
	// Required is set to true when we want to
	// specify that a value should be entered after the
	// main root command is executed
	Required bool
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
	if a.Type == argument.Bool && a.Required {
		return fmt.Errorf("skapt: Cannot have type Bool and requried true")
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
	case 1:
		if a.Required {
			return fmt.Errorf("skapt: Command %s requires a value", args[0])
		}

		ctx := context.New(a.Flags, nil)
		return a.Handler(ctx)
	}

	root := command.Command{
		Name:     args[0],
		Type:     a.Type,
		Flags:    a.Flags,
		Handler:  a.Handler,
		Required: a.Required,
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

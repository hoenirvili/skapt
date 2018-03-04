package command

import (
	"fmt"

	"github.com/hoenirvili/skapt/argument"
	"github.com/hoenirvili/skapt/context"
	"github.com/hoenirvili/skapt/flag"
)

// Command holds information in order to construct and parse
// commands in command line arguments
type Command struct {
	// Name is the name of the command
	Name string
	// Usage describes how to use the command
	Usage string
	// Description holds formal information about
	// how the comamnd behaves and what cand do
	Description string
	// Flags aditional flags of the command
	// to change behaviour or to add sub-features
	Flags flag.Flags
	// Type is the command type value of the command
	Type argument.Type
	// Handler is the action that will be
	// executed after the command will be parsed
	Handler context.Handler
	// Required ture when the value is required
	Required bool

	// parseed is true when the comamnd has value has
	// been parsed sucessfully
	parsed bool
	// value hoolds the underlying value
	value interface{}
}

// Validate returns nil if the command form is valid
func (c Command) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("skapt: Empty command name")
	}

	if c.Handler == nil {
		return fmt.Errorf("skapt: Command has no handler attached")
	}

	if c.Flags != nil {
		if err := c.Flags.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// ParseValue parses the value given, if the value
// does not contain a form of command.Type this will return
// an error
func (c *Command) ParseValue(value string) error {
	var err error
	c.value, err = argument.ParseValue(c.Type, value)
	if err != nil {
		return err
	}
	c.parsed = true
	return nil
}

// StringValue return the value as string
func (c Command) StringValue() string {
	value, _ := c.value.(string)
	return value
}

// IntValue return the value as int
func (c Command) IntValue() int {
	value, _ := c.value.(int)
	return value
}

// BoolValue returns true if the command has been parsed
func (c Command) BoolValue() bool {
	if c.parsed {
		return true
	}

	return false
}

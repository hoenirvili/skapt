// Package skapt provides a tiny interface
// to create and manage your command line aplications
package flag

import (
	"fmt"

	"github.com/hoenirvili/skapt/argument"
)

// Flag type holds all information
// for a flag to be parsed
type Flag struct {
	// Short is the short name of the flag
	Short string
	// Long is the long name of the flag
	Long string
	// Description holds the description of the flag
	Description string
	// Type is the flag type
	Type argument.Type
	// Required is true when the flag need to be passed
	Required bool

	value interface{}
	// parsed is set to true when the argument has been parsed
	parsed bool
}

func (f Flag) Parsed() bool {
	return f.parsed
}

func (f *Flag) ParseValue(value string) error {
	var err error
	f.value, err = argument.ParseValue(f.Type, value)
	if err != nil {
		return err
	}

	f.parsed = true
	return nil
}

func (f Flag) StringValue() string {
	value, _ := f.value.(string)
	return value
}

func (f Flag) IntValue() int {
	value, _ := f.value.(int)
	return value
}

func (f Flag) BoolValue() bool {
	if f.parsed {
		return true
	}

	return false
}

// String returns the flag as string format
func (f Flag) String() string {
	str := ""
	if f.Short != "" {
		str += "-" + f.Short
	}
	if f.Long != "" {
		if f.Short != "" {
			str += " "
		}
		str += "--" + f.Long
	}

	return str
}

var _ fmt.Stringer = (*Flag)(nil)

// Validate validates if the flag definitions are valid
func (f Flag) Validate() error {
	if f.Short == "" && f.Long == "" {
		return fmt.Errorf("skapt: Empty flag name")
	}
	return nil
}

func (f Flag) Eq(name string) bool {
	if name == "" {
		return false
	}
	if f.Short == name || f.Long == name {
		return true
	}
	return false
}

// Valid retrun true if the arg is a valid
// short or long format flag
func Valid(arg string) bool {
	return argument.Short(arg) || argument.Long(arg)
}

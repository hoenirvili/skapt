// Package flag provides types and functions
// to create and manage your command line applications
package flag

import (
	"fmt"

	"github.com/hoenirvili/skapt/argument"
)

// Flag type holds all information for a flag
// to be parsed on the command line
type Flag struct {
	// Short is the short name of the flag
	Short string
	// Long is the long name of the flag
	Long string
	// Description holds information describing the
	// behavior of the flag can have
	Description string
	// Type is the value type of the flag
	Type argument.Type
	// Required is true when the user is required to
	// pass this flag in the command line
	Required bool
	// value holds the underlying value of the flag
	// if the flag is parsed this will  be != nil
	value *argument.Value
}

// Parsed return true if the flag is parsed
func (f Flag) Parsed() bool {
	return f.value != nil
}

var _ fmt.Stringer = (*Flag)(nil)

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

// Validate validates if the flag definitions are valid
func (f Flag) Validate() error {
	if f.Short == "" && f.Long == "" {
		return fmt.Errorf("empty flag name")
	}
	if f.Short == f.Long {
		return fmt.Errorf("short and long names are the same")
	}
	return nil
}

// Is returns true if the argument name is present
// in the short or long name of the flag
func (f Flag) Is(arg string) bool {
	return (arg != "" &&
		(f.Short == arg || f.Long == arg))
}

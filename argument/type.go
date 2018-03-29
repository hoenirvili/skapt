// Package argument provides small set of types
// to parse and interpret command line arguments
package argument

import "fmt"

// Type defines the type of an argument
type Type uint8

const (
	// Bool is the flag value of type bool
	Bool Type = iota
	// Int is the flag value of type int
	Int
	// String is the flag value of type string
	String
	// Float is the flag value of type float
	Float
)

var _ fmt.Stringer = (*Type)(nil)

// String returns the type in an outputted format
func (t Type) String() string {
	switch t {
	case String:
		return "string"
	case Bool:
		return "bool"
	case Int:
		return "int"
	case Float:
		return "float"
	default:
		return "unknown type"
	}
}

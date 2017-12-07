package argument

import "fmt"

// Type defines the type that a argument can be
type Type uint8

const (
	// Bool is the type of flag that
	// has no value and needs to be present
	Bool Type = iota
	// Int is the type of flag that
	// receives a value of type int
	Int
	// String is the type of flag that
	// receives a value of type string
	String
)

var _ fmt.Stringer = (*Type)(nil)

func (t Type) String() string {
	switch t {
	case String:
		return "string"
	case Bool:
		return "bool"
	case Int:
		return "int"
	default:
		return "undefined type"
	}
}

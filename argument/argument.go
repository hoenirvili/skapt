// Package argument provides small set of types
// to parse and interpret command line arguments
package argument

import (
	"fmt"
	"strconv"
	"strings"
)

// Short return true given the arg is a short one
func Short(arg string) bool {
	return (len(arg) == 2 &&
		arg[0] == byte('-') && arg[1] != byte('-') &&
		!strings.Contains(arg, " "))
}

// Long return true given the arg is a long one
func Long(arg string) bool {
	return (len(arg) > 2 &&
		arg[:2] == "--" &&
		!strings.Contains(arg, " "))
}

// ShortTrim returns the argument without it's short prefix
func ShortTrim(arg string) string {
	if Short(arg) {
		return arg[1:]
	}

	return arg
}

// LongTrim returns the argument without it's long prefix and
// his target value if he has one
func LongTrim(arg string) (string, string) {
	if Long(arg) {
		if !strings.Contains(arg, "=") {
			return arg[2:], ""
		}

		args := strings.SplitN(arg[2:], "=", 2)
		return args[0], args[1]
	}

	return arg, ""
}

// Value can hold any argument of type argument.Type
type Value struct {
	sv string
	t  Type
	v  interface{}
}

// String returns the value as type string
func (v Value) String() string {
	value, _ := v.v.(string)
	return value
}

// Int returns the value as type int
func (v Value) Int() int {
	value, _ := v.v.(int)
	return value
}

// Bool returns the value as type bool
func (v Value) Bool() bool {
	value, _ := v.v.(bool)
	return value
}

// NewValue takes a command line string argument and his desired type
// and returns a new Value that can convert to t Type
func NewValue(arg string, t Type) *Value {
	return &Value{
		sv: arg,
		t:  t,
	}
}

// Parse parses the value as a given t Type given
// if the value is not a valid t Type it will return an error
func (v *Value) Parse() error {
	var err error
	switch v.t {
	case Bool:
		v.v = true
	case String:
		v.v = v.sv
	case Int:
		vint, err := strconv.ParseInt(v.sv, 10, 32)
		if err != nil {
			return fmt.Errorf("cannot parse %s as integer", v.sv)
		}
		v.v = int(vint)
	}

	return err
}

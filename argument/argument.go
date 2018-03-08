package argument

import (
	"strconv"
	"strings"
)

// Short return true given the arg is a short one
func Short(arg string) bool {
	return (len(arg) > 1 &&
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
		args := strings.SplitN(arg[2:], "=", 2)
		return args[0], args[1]
	}

	return arg, ""
}

// Value can hold an argument value
type Value struct {
	sv string
	t  Type
	v  interface{}
}

// NewValue based on the string arg and his type returns new Value
// tha can be easly be parsed
func NewValue(arg string, t Type) *Value {
	return &Value{
		sv: arg,
		t:  t,
	}
}

// ParseValue parses the value as a given t Type given
// if the value is not a valid t Type it will return an error
func ParseValue(t Type, value string) (interface{}, error) {
	var (
		v   interface{}
		err error
	)
	switch t {
	case Bool:
	case String:
		v = value
	case Int:
		v, err = strconv.ParseInt(value, 10, 32)
		if err != nil {
			return nil, err
		}
	}

	return v, err
}

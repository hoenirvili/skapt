package argument

import (
	"strconv"
	"strings"
)

func Short(arg string) bool {
	return (len(arg) > 1 &&
		arg[0] == byte('-') && arg[1] != byte('-') &&
		!strings.Contains(arg, " "))
}

func Long(arg string) bool {
	return (len(arg) > 2 &&
		arg[:2] == "--" &&
		!strings.Contains(arg, " "))
}

func ShortTrim(arg string) string {
	if Short(arg) {
		return arg[1:]
	}

	return arg
}

func LongTrim(arg string) (string, string) {
	if Long(arg) {
		args := strings.SplitN(arg[2:], "=", 2)
		return args[0], args[1]
	}

	return arg, ""
}

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

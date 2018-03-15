package flag

import (
	"fmt"

	"github.com/hoenirvili/skapt/argument"
)

// Flags type is a list of command line flags
type Flags []Flag

// RequiredAreParsed checks if all the required flags has been parsed
func (f Flags) RequiredAreParsed() error {
	for _, flag := range f {
		if flag.Required && flag.value == nil {
			return fmt.Errorf("flag: Flag %s is not parsed", flag)
		}
	}

	return nil
}

// Validate checks all flags to be validated
func (f Flags) Validate() error {
	m := len(f)
	if m == 0 {
		return nil
	}

	n := m - 1
	for i := 0; i < n; i++ {
		if err := f[i].Validate(); err != nil {
			return err
		}
		for j := i + 1; j < m; j++ {
			if f[i].Is(f[j].Short) ||
				f[i].Is(f[j].Long) {
				return fmt.Errorf("flag: Every flag should be unique")
			}
		}
	}
	// don't forget to check the integrity of the last flag
	return f[n].Validate()
}

// Flag returns the flag that mathes the name provided
func (f Flags) Flag(name string) *Flag {
	for key, flag := range f {
		if flag.Is(name) {
			return &f[key]
		}
	}
	return nil
}

//Bool return the bool value of the arg
func (f Flags) Bool(arg string) bool {
	flag := f.Flag(arg)
	if flag == nil {
		return false
	}

	if flag.value == nil {
		return false
	}

	return flag.value.Bool()
}

// Int return the int value of the flag
func (f Flags) Int(arg string) int {
	flag := f.Flag(arg)
	if flag == nil {
		return 0
	}

	if flag.value == nil {
		return 0
	}

	return flag.value.Int()
}

// String return the string value of the arg
func (f Flags) String(arg string) string {
	flag := f.Flag(arg)
	if flag == nil {
		return ""
	}

	if flag.value == nil {
		return ""
	}

	return flag.value.String()
}

// Parse parses the command line arguments and returns
// the one that has are not flags
func (f Flags) Parse(args []string) ([]string, error) {
	if len(args) == 0 {
		return args, nil
	}

	var unparsed []string
	n := len(args)
	for i := 0; i < n; i++ {
		// skip empty flags
		if args[i] == "" {
			continue
		}

		value, arg := "", ""

		// decide which type of argument we are dealing
		// and trim their prefixes extracting only their names
		switch {
		case argument.Short(args[i]):
			arg = argument.ShortTrim(args[i])
		case argument.Long(args[i]):
			arg, value = argument.LongTrim(args[i])
		default:
			arg = args[i]
		}

		flag := f.Flag(arg)
		if flag == nil {
			unparsed = append(unparsed, arg)
			continue
		}

		if flag.Parsed() {
			return nil, fmt.Errorf("flag: Flag %s is already parsed", arg)
		}

		switch flag.Type {
		case argument.Bool:
			if value != "" {
				return nil, fmt.Errorf("flag: Flag %s does not require a value", arg)
			}
		case argument.String, argument.Int:
			if value == "" {
				if i+1 > n || argument.Long(args[i]) {
					return nil, fmt.Errorf("flag: Flag %s requires a value", arg)
				}
			}
			if i+1 < n && value == "" {
				value = args[i+1]
				if argument.Short(value) || argument.Long(value) {
					return nil, fmt.Errorf(
						"flag: Invalid value for %s, need value of type %s ",
						arg, flag.Type,
					)
				}
				i++
			}
		default:
			return nil, fmt.Errorf("flag: Can't parse flag of type %s", flag.Type)
		}

		v := argument.NewValue(value, flag.Type)
		if err := v.Parse(); err != nil {
			return nil, err
		}

		flag.value = v
	}

	// if we have any required flags and their were not parsed
	if err := f.RequiredAreParsed(); err != nil {
		return nil, err
	}

	return unparsed, nil
}

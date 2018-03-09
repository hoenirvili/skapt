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
	for _, flag := range f {
		if err := flag.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// Flag returns the flag that mathes the name provided
func (f Flags) Flag(name string) Flag {
	for _, flag := range f {
		if flag.Is(name) {
			return flag
		}
	}

	return Flag{}

}

// Bool return the bool value of the arg
func (f Flags) Bool(arg string) bool {
	return f.Flag(arg).value.Bool()
}

// Int return the int value of the flag
func (f Flags) Int(arg string) int {
	return f.Flag(arg).value.Int()
}

// String return the string value of the arg
func (f Flags) String(arg string) string {
	return f.Flag(arg).value.String()
}

// Parse parses the command line arguments and returns
// the one that has are not flags
func (f Flags) Parse(args []string) ([]string, error) {
	n := len(args)
	if n == 0 {
		return args, nil
	}

	args = argument.Strip(args)

	unparsed := make([]string, 0)
	for i := 0; i < n; i++ {
		value, parsed := "", false
		for j := range f {
			if !f[j].Is(args[i]) {
				continue
			}

			if f[j].Type != argument.Bool && i+1 < n {
				value = args[i+1]
				i++
			}

			v := argument.NewValue(value, f[i].Type)
			if err := v.Parse(); err != nil {
				return nil, err
			}
			f[i].value = v
			parsed = true
			break
		}

		if !parsed {
			unparsed = append(unparsed, args[i])
		}
	}

	if err := f.RequiredAreParsed(); err != nil {
		return nil, err
	}

	return unparsed, nil
}

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
func (f Flags) Flag(name string) Flag {
	var empty Flag
	for _, flag := range f {
		if flag.Is(name) {
			return flag
		}
	}

	return empty

}

// Bool return the bool value of the arg
func (f Flags) Bool(arg string) bool {
	if value := f.Flag(arg).value; value != nil {
		return value.Bool()
	}

	return false
}

// Int return the int value of the flag
func (f Flags) Int(arg string) int {
	if value := f.Flag(arg).value; value != nil {
		return value.Int()
	}
	return 0
}

// String return the string value of the arg
func (f Flags) String(arg string) string {
	if value := f.Flag(arg).value; value != nil {
		return value.String()
	}
	return ""
}

// Parse parses the command line arguments and returns
// the one that has are not flags
func (f Flags) Parse(args []string) ([]string, error) {
	if len(args) == 0 {
		return args, nil
	}

	args = argument.Strip(args)
	var unparsed []string
	n := len(args)
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
			f[j].value = v
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

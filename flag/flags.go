// Package flag provides types and functions
// to create and manage your command line applications
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
			return fmt.Errorf("Option %s is required", flag)
		}
	}

	return nil
}

// AppendHelpIfNotPresent if the help flag is not present
// in the underlying slice append the default one
func (f *Flags) AppendHelpIfNotPresent() {
	if h, hl := f.Flag("h"), f.Flag("help"); h != nil || hl != nil {
		return
	}

	*f = append(*f, Flag{
		Short:       "h",
		Long:        "help",
		Description: "Print out the help menu",
	})
}

// AppendVersionIfNotPreset if the version flag is not present
// in the underlying slice append the default one
func (f *Flags) AppendVersionIfNotPreset() {
	if h, hl := f.Flag("v"), f.Flag("version"); h != nil || hl != nil {
		return
	}

	*f = append(*f, Flag{
		Short:       "v",
		Long:        "version",
		Description: "Print out the version of the program",
	})
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
				return fmt.Errorf("every flag should be unique")
			}
		}
	}
	// don't forget to check the integrity of the last flag
	return f[n].Validate()
}

// Flag returns the flag that matches the name provided
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
			unparsed = append(unparsed, args[i])
			continue
		}

		flag := f.Flag(arg)
		if flag == nil {
			unparsed = append(unparsed, arg)
			continue
		}

		if flag.Parsed() {
			return nil, fmt.Errorf("Multiple occurrences of %s", arg)
		}

		switch flag.Type {
		case argument.Bool:
			if value != "" {
				return nil, fmt.Errorf("Option %s does not require a value", arg)
			}
		case argument.String, argument.Int:
			if value == "" {
				if i+1 >= n || argument.Long(args[i]) {
					return nil, fmt.Errorf("Option %s requires a value", arg)
				}
			}
			if i+1 < n && value == "" {
				value = args[i+1]
				i++
			}
		default:
			return nil, fmt.Errorf("Cannot parse flag of type %s", flag.Type)
		}

		v := argument.NewValue(value, flag.Type)
		if err := v.Parse(); err != nil {
			return nil, err
		}

		flag.value = v
	}

	return unparsed, nil
}

package flag

import (
	"fmt"
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

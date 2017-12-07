package flag

import (
	"fmt"
)

type Flags []Flag

func (f Flags) Flag(name string) Flag {
	for _, flag := range f {
		if flag.Eq(name) {
			return flag
		}
	}

	return Flag{}
}

func (f Flags) AllRequiredParsed() error {
	for _, flag := range f {
		if flag.Required && !flag.parsed {
			return fmt.Errorf("Flag %s is required", flag)
		}
	}

	return nil
}

func (f Flags) Set(name string, value string) {
	for key := range f {
		if f[key].Eq(name) {
			f[key].value, f[key].parsed = value, true
			break
		}
	}
}

func (f Flags) Validate() error {
	for _, flag := range f {
		if err := flag.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (f Flags) Contains(name string) bool {
	for _, flag := range f {
		if flag.Eq(name) {
			return true
		}
	}

	return false
}

func (f *Flags) AppendDefault() {
	for _, flg := range []Flag{
		{Short: "h", Long: "help"},
		{Short: "v", Long: "version"},
	} {
		if !f.Contains(flg.Short) {
			*f = append(*f, flg)
		}
	}
}

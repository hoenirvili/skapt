package context

import (
	"github.com/hoenirvili/skapt/flag"
)

// Handler type used to hold the hole main
// execution of every command
type Handler func(ctx Context) error

type Context struct {
	flags flag.Flags
	share map[string]interface{}
}

func New(flags flag.Flags, share map[string]interface{}) Context {
	return Context{
		flags: flags,
		share: share,
	}
}

func (c Context) String(name string) string {
	flag := c.flags.Flag(name)
	return flag.StringValue()
}

func (c Context) Bool(name string) bool {
	flag := c.flags.Flag(name)
	return flag.BoolValue()
}

func (c Context) Int(name string) int {
	flag := c.flags.Flag(name)
	return flag.IntValue()
}

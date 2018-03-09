package flag_test

import (
	"github.com/hoenirvili/skapt/flag"
	gc "gopkg.in/check.v1"
)

type flagsSuite struct{}

var _ = gc.Suite(&flagsSuite{})

func (f flagsSuite) newFlags() flag.Flags {
	return flag.Flags{
		{},
		{Short: "u", Long: "url"},
		{Short: "k", Long: "specialk"},
		{Short: "l"},
		{Long: "full"},
	}
}

func (f flagsSuite) TestValidate(c *gc.C) {
	flags := f.newFlags()
	err := flags.Validate()
	c.Assert(err, gc.NotNil)

	flags = flags[1:]
	err = flags.Validate()
	c.Assert(err, gc.IsNil)
}

var empty flag.Flag

func (f flagsSuite) TestFlag(c *gc.C) {
	flags := f.newFlags()
	fl := flags.Flag("")
	c.Assert(fl, gc.DeepEquals, empty)

	fl = flags.Flag("u")
	c.Assert(fl, gc.DeepEquals, flag.Flag{
		Short: "u", Long: "url"})
}

func (f flagsSuite) TestParse(c *gc.C) {
	flags := f.newFlags()[1:]
	args := []string{"-u", "k", "full", "somevalue", "someothervalue"}
	unparsed, err := flags.Parse(args)
	c.Assert(err, gc.IsNil)
	c.Assert(unparsed, gc.NotNil)
	c.Assert(unparsed, gc.DeepEquals,
		[]string{"somevalue", "someothervalue"})
}

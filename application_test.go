package skapt_test

import (
	"github.com/hoenirvili/skapt"
	"github.com/hoenirvili/skapt/argument"
	"github.com/hoenirvili/skapt/flag"
	gc "gopkg.in/check.v1"
)

type appSuite struct{}

var _ = gc.Suite(&appSuite{})

func handler(flags flag.Flags, args []string) error {
	return nil

}

func (a appSuite) TestExecValidateWithErrors(c *gc.C) {
	apps := []skapt.Application{
		{},
		{Name: "test"},
		{Name: "test", Handler: handler},
		{Name: "test", Handler: handler, Flags: flag.Flags{
			{},
		}},
	}

	for _, app := range apps {
		err := app.Exec(nil)
		c.Assert(err, gc.NotNil)
	}
}

func (a appSuite) TestExec(c *gc.C) {
	args := []string{"./test"}
	app := skapt.Application{
		Name: "test",
		Handler: func(flags flag.Flags, args []string) error {
			s, l := flags.Bool("u"), flags.Bool("url")
			c.Assert(s, gc.Equals, false)
			c.Assert(l, gc.Equals, false)
			c.Assert(args, gc.DeepEquals, args)
			return nil
		},
		Flags: flag.Flags{
			{Short: "u", Long: "url", Type: argument.Bool},
		},
	}

	err := app.Exec(args)
	c.Assert(err, gc.IsNil)
}

func (a appSuite) TestExecWithErrors(c *gc.C) {
	args := []string{"./test"}
	app := skapt.Application{
		Name: "test",
		Handler: func(flags flag.Flags, args []string) error {
			s, l := flags.Bool("u"), flags.Bool("url")
			c.Assert(s, gc.Equals, false)
			c.Assert(l, gc.Equals, false)
			c.Assert(args, gc.DeepEquals, args)
			return nil
		},
		NArgs: 2,
		Flags: flag.Flags{
			{Short: "u", Long: "url", Type: argument.Bool},
		},
	}

	err := app.Exec(args)
	c.Assert(err, gc.NotNil)

	app.Flags[0].Type = argument.Int
	app.Handler = func(flags flag.Flags, args []string) error {
		s, l := flags.Int("u"), flags.Int("url")
		c.Assert(s, gc.Equals, 0)
		c.Assert(l, gc.Equals, 0)
		c.Assert(args, gc.DeepEquals, []string{"./test"})
		return nil
	}
	
	err = app.Exec([]string{"./test", "-u", "huifsdh1"})
	c.Assert(err, gc.NotNil)

}

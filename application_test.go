package skapt_test

import (
	"bytes"
	"errors"
	"io"

	"github.com/hoenirvili/skapt"
	"github.com/hoenirvili/skapt/argument"
	"github.com/hoenirvili/skapt/flag"
	gc "gopkg.in/check.v1"
)

type appSuite struct{}

var _ = gc.Suite(&appSuite{})

func handler(ctx *skapt.Context) error {
	return nil
}
func (a appSuite) defaultApplication(stdout, stderr io.Writer) *skapt.Application {
	return &skapt.Application{
		Name:        "test",
		Description: "test description",
		Handler: func(ctx *skapt.Context) error {
			return nil
		},
		Stdout: stdout,
		Stderr: stderr,
	}
}

func (a appSuite) TestExecValidateWithErrors(c *gc.C) {
	stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}

	apps := []skapt.Application{
		{},
		{Name: "test"},
		{
			Name:    "test",
			Handler: handler,
			Stdout:  stdout,
			Stderr:  stderr,
		},
		{
			Name:    "test",
			Handler: handler,
			Flags:   flag.Flags{{}},
		},
	}

	for _, app := range apps {
		err := app.Exec(nil)
		c.Assert(err, gc.NotNil)
	}

	so, se := stdout.String(), stderr.String()
	c.Assert(so, gc.DeepEquals, "")
	c.Assert(se, gc.DeepEquals, "No arguments to execute")

}

func (a appSuite) TestExec(c *gc.C) {
	args := []string{"./test"}
	stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}
	app := skapt.Application{
		Name: "test",
		Handler: func(ctx *skapt.Context) error {
			s, l := ctx.Bool("u"), ctx.Bool("url")
			c.Assert(s, gc.Equals, false)
			c.Assert(l, gc.Equals, false)
			c.Assert(ctx.Args, gc.DeepEquals, args)
			return nil
		},
		Flags: flag.Flags{
			{Short: "u", Long: "url", Type: argument.Bool},
		},
		Stdout: stdout,
		Stderr: stderr,
	}

	err := app.Exec(args)
	c.Assert(err, gc.IsNil)

	so, se := stdout.String(), stderr.String()
	c.Assert(so, gc.DeepEquals, "")
	c.Assert(se, gc.DeepEquals, "")

	app.Stdout, app.Stderr = nil, nil
	err = app.Exec(args)
	c.Assert(err, gc.IsNil)

}

func (a appSuite) TestExecHandler(c *gc.C) {
	args := []string{
		"./downloader", "-u", "https://someexternalink.com",
		"--times=3",
		"--debug", "--hint=10", "merge-link", "another-arg",
	}

	stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}
	app := skapt.Application{
		Name: "test",
		Flags: flag.Flags{
			{Short: "u", Type: argument.String},
			{Long: "times", Type: argument.Int},
			{Long: "debug", Type: argument.Bool},
			{Long: "hint", Type: argument.Int},
		},
		NArgs: 2,
		Handler: func(ctx *skapt.Context) error {
			u := ctx.String("u")
			t := ctx.Int("times")
			d := ctx.Bool("debug")
			h := ctx.Int("hint")
			c.Assert(u, gc.DeepEquals, "https://someexternalink.com")
			c.Assert(t, gc.DeepEquals, 3)
			c.Assert(d, gc.Equals, true)
			c.Assert(h, gc.DeepEquals, 10)
			c.Assert(ctx.Args, gc.DeepEquals, []string{"./downloader",
				"merge-link", "another-arg"})
			return nil
		},
		Stdout: stdout,
		Stderr: stderr,
	}

	err := app.Exec(args)
	c.Assert(err, gc.IsNil)
	so, se := stdout.String(), stdout.String()
	c.Assert(so, gc.Equals, "")
	c.Assert(se, gc.Equals, "")

}

func (a appSuite) TestExecWithErrors(c *gc.C) {
	args := []string{"./test"}
	stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}
	app := skapt.Application{
		Name: "test",
		Handler: func(ctx *skapt.Context) error {
			s, l := ctx.Bool("u"), ctx.Bool("url")
			c.Assert(s, gc.Equals, false)
			c.Assert(l, gc.Equals, false)
			c.Assert(ctx.Args, gc.DeepEquals, args)
			return nil
		},
		NArgs: 2,
		Flags: flag.Flags{
			{Short: "u", Long: "url", Type: argument.Bool},
		},
		Stdout: stdout,
		Stderr: stderr,
	}

	err := app.Exec(args)
	c.Assert(err, gc.NotNil)
	so, se := stdout.String(), stderr.String()
	c.Assert(so, gc.DeepEquals, "")
	c.Assert(se, gc.DeepEquals, "Expecting at least 2 additional arguments")

	stdout.Reset()
	stderr.Reset()

	app.Flags[0].Type = argument.Int
	app.Handler = func(ctx *skapt.Context) error {
		s, l := ctx.Int("u"), ctx.Int("url")
		c.Assert(s, gc.Equals, 0)
		c.Assert(l, gc.Equals, 0)
		c.Assert(ctx.Args, gc.DeepEquals, []string{"./test"})
		return nil
	}

	err = app.Exec([]string{"./test", "-u", "huifsdh1"})
	c.Assert(err, gc.NotNil)
	so, se = stdout.String(), stderr.String()
	c.Assert(so, gc.DeepEquals, "")
	c.Assert((len(se) > 0), gc.Equals, true)
}

var (
	expectedHelp = `
Usage: test [OPTIONS] [ARG...]
       test [ --help | -h | -v | --version ]

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent consectetur 
in dolor lobortis, non ultrices nibh condimentum. Fusce suscipit ultrices. 
Sed a malesuada urna. Lorem ipsum dolor sit consectetur adipiscing 
elit. Vivamus laoreet tellus vel sem euismod, accumsan odio pulvinar. 
Donec eget ante venenatis, gravida eros sagittis elit. Nam nec 
arcu augue. Sed mattis lobortis at malesuada leo bibendum at. 
Pellentesque et condimentum erat. feugiat id ex non iaculis. Donec 
efficitur ac lectus hendrerit. Sed feugiat augue nec nibh rutrum, 
in ornare viverra. 

Options:

-l --long-description  Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent consectetur 
                       in dolor lobortis, non ultrices nibh condimentum. Fusce suscipit ultrices. 
                       Sed a malesuada urna. Lorem ipsum dolor sit consectetur adipiscing 
                       elit. Vivamus laoreet tellus vel sem euismod, accumsan odio pulvinar. 
                       Donec eget ante venenatis, gravida eros sagittis elit. Nam nec 
                       arcu augue. Sed mattis lobortis at malesuada leo bibendum at. 
                       Pellentesque et condimentum erat. feugiat id ex non iaculis. Donec 
                       efficitur ac lectus hendrerit. Sed feugiat augue nec nibh rutrum, 
                       in ornare viverra. 
-h --help              Print out the help menu
-v --version           Print out the version of the program
`[1:]

	expectedVersion = `
Version 1.0.0
`[1:]
)

const description = `
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent consectetur 
in dolor lobortis, non ultrices nibh condimentum. Fusce suscipit 
ultrices. Sed a malesuada urna. Lorem ipsum dolor sit 
consectetur adipiscing elit. Vivamus laoreet tellus vel sem euismod, 
accumsan odio pulvinar. Donec eget ante venenatis, gravida eros 
sagittis elit. Nam nec arcu augue. Sed mattis lobortis 
at malesuada leo bibendum at. Pellentesque et condimentum erat. 
feugiat id ex non iaculis. Donec efficitur ac lectus 
hendrerit. Sed feugiat augue nec nibh rutrum, in ornare viverra. 
`

func (a appSuite) TestExecRender(c *gc.C) {
	stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}
	app := skapt.Application{
		Name:        "test",
		Description: description,
		Handler: func(ctx *skapt.Context) error {
			return nil
		},
		Flags: flag.Flags{
			{
				Short: "l", Long: "long-description",
				Description: description,
			},
		},
		Stdout: stdout,
		Stderr: stderr,
	}

	args := []string{"./test", "--help"}
	err := app.Exec(args)
	c.Assert(err, gc.IsNil)

	se := stderr.String()
	so := stdout.String()
	c.Assert(so, gc.DeepEquals, expectedHelp)
	c.Assert(se, gc.DeepEquals, "")

	stdout.Reset()
	stderr.Reset()

	args = []string{"./test", "--version"}
	app.Version = "1.0.0"
	err = app.Exec(args)
	c.Assert(err, gc.IsNil)

	so = stdout.String()
	se = stderr.String()
	c.Assert(so, gc.DeepEquals, expectedVersion)
	c.Assert(se, gc.DeepEquals, "")

}

type deferErr struct{}

func (d deferErr) Write([]byte) (int, error) {
	return 0, deferr
}

var deferr = errors.New("err")

func (a appSuite) TestExecDeferError(c *gc.C) {
	app := skapt.Application{
		Name:        "test",
		Description: "test description",
		Handler: func(ctx *skapt.Context) error {
			return nil
		},
		Stderr: &deferErr{},
	}

	err := app.Exec(nil)
	c.Assert(err, gc.DeepEquals, deferr)
}

func (a appSuite) TestExecRequired(c *gc.C) {
	stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}
	app := a.defaultApplication(stdout, stderr)
	app.Flags = flag.Flags{
		{Short: "e", Long: "expl", Required: true},
	}

	err := app.Exec([]string{"./test"})
	c.Assert(err, gc.NotNil)
	so, se := stdout.String(), stderr.String()

	c.Assert(so, gc.DeepEquals, "")
	c.Assert(se, gc.DeepEquals, "Option -e --expl is required")

	stdout.Reset()
	stderr.Reset()

	err = app.Exec([]string{"./test", "-e"})
	c.Assert(err, gc.IsNil)
	so, se = stdout.String(), stderr.String()

	c.Assert(so, gc.DeepEquals, "")
	c.Assert(se, gc.DeepEquals, "")
}

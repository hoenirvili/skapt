package skapt_test

import (
	gc "gopkg.in/check.v1"
)

type appSuite struct{}

var _ = gc.Suite(&appSuite{})

// func (a appSuite) TestExecWithErrors(c *gc.C) {
// 	apps := []skapt.Application{
// 		{},
// 		{
// 			Name: "test",
// 		},
// 		{
// 			Name:    "test",
// 			Handler: handler,
// 			Flags:   flag.Flags{{}},
// 		},
// 		{
// 			Name:     "test",
// 			Handler:  handler,
// 			Commands: command.Commands{{Name: ""}},
// 		},
// 		{
// 			Name:     "test",
// 			Handler:  handler,
// 			Commands: command.Commands{{Name: "test"}},
// 		},
// 		{
// 			Name:    "test",
// 			Handler: handler,
// 			Commands: command.Commands{{
// 				Name:    "test",
// 				Handler: handler,
// 				Flags:   flag.Flags{{}},
// 			}},
// 		},
// 	}
//
// 	errs := []error{
// 		errors.New("skapt: Empty application name"),
// 		errors.New("skapt: Empty application handler"),
// 		errors.New("skapt: Empty flag name"),
// 		errors.New("skapt: Empty command name"),
// 		errors.New("skapt: Command has no handler attached"),
// 		errors.New("skapt: Empty flag name"),
// 	}
//
// 	for key, app := range apps {
// 		got := app.Exec(nil)
// 		c.Assert(got, gc.DeepEquals, errs[key])
// 	}
//
// 	expected := errors.New("skapt: No arguments to execute")
// 	app := skapt.Application{
// 		Name:    "test",
// 		Handler: handler,
// 		Flags: flag.Flags{
// 			{
// 				Short: "test",
// 			},
// 		},
// 	}
// 	got := app.Exec(nil)
// 	c.Assert(got, gc.DeepEquals, expected)
//
// 	expected = errors.New("skapt: Cannot have type Bool and requried true")
// 	app.Required = true
// 	got = app.Exec([]string{"./test"})
// 	c.Assert(got, gc.DeepEquals, expected)
//
// 	expected = errors.New("skapt: Command ./test requires a value")
// 	app.Required = true
// 	app.Type = argument.String
// 	got = app.Exec([]string{"./test"})
// 	c.Assert(got, gc.DeepEquals, expected)
// }
//
// func (a appSuite) TestExecWithEmtyArgs(c *gc.C) {
// 	apps := []skapt.Application{
// 		{
// 			Name:    "test",
// 			Handler: handler,
// 			Flags:   flag.Flags{{Short: "test"}},
// 		}, {
// 			Name:    "test",
// 			Handler: handler,
// 			Flags:   flag.Flags{{Long: "test"}},
// 		}, {
// 			Name:    "test",
// 			Handler: handler,
// 			Flags: flag.Flags{{
// 				Short: "test",
// 				Long:  "test",
// 			}},
// 		}, {
// 			Name:    "test",
// 			Handler: handler,
// 			Commands: command.Commands{{
// 				Name:    "test",
// 				Handler: handler,
// 			}},
// 		}, {
// 			Name:    "test",
// 			Handler: handler,
// 			Flags:   flag.Flags{{Short: "test"}},
// 			Commands: command.Commands{{
// 				Name:    "test",
// 				Handler: handler,
// 			}},
// 		}, {
// 			Name:    "test",
// 			Handler: handler,
// 			Flags:   flag.Flags{{Long: "test"}},
// 			Commands: command.Commands{{
// 				Name:    "test",
// 				Handler: handler,
// 			}},
// 		}, {
// 			Name:    "test",
// 			Handler: handler,
// 			Flags: flag.Flags{{
// 				Short: "test",
// 				Long:  "test",
// 			}},
// 			Commands: command.Commands{{
// 				Name:    "test",
// 				Handler: handler,
// 			}},
// 		}, {
// 			Name:    "test",
// 			Handler: handler,
// 			Flags: flag.Flags{{
// 				Short: "test",
// 				Long:  "test",
// 			}},
// 			Commands: command.Commands{{
// 				Name:    "test",
// 				Handler: handler,
// 				Flags: flag.Flags{{
// 					Short: "test",
// 				}},
// 			}},
// 		}, {
// 			Name:    "test",
// 			Handler: handler,
// 			Flags: flag.Flags{{
// 				Short: "test",
// 				Long:  "test",
// 			}},
// 			Commands: command.Commands{{
// 				Name:    "test",
// 				Handler: handler,
// 				Flags:   flag.Flags{{Long: "test"}},
// 			}},
// 		}, {
// 			Name:    "test",
// 			Handler: handler,
// 			Flags: flag.Flags{{
// 				Short: "test",
// 				Long:  "test",
// 			}},
// 			Commands: command.Commands{{
// 				Name:    "test",
// 				Handler: handler,
// 				Flags: flag.Flags{
// 					{Short: "test", Long: "test"}},
// 			}},
// 		},
// 	}
//
// 	for _, app := range apps {
// 		got := app.Exec(nil)
// 		c.Assert(got, gc.NotNil)
// 		c.Assert(got, gc.DeepEquals,
// 			errors.New("skapt: No arguments to execute"))
// 	}
// }
//
// func (a appSuite) TestExecHandlerError(c *gc.C) {
// 	testError := fmt.Errorf("test error")
// 	app := skapt.Application{
// 		Name: "test",
// 		Handler: func(ctx context.Context) error {
// 			c.Assert(ctx, gc.DeepEquals,
// 				context.New(nil, nil))
// 			return testError
// 		},
// 	}
//
// 	args := []string{"./test"}
// 	err := app.Exec(args)
// 	c.Assert(err, gc.DeepEquals, testError)
// }
//
// func (a appSuite) TestExecWithFlags(c *gc.C) {
// 	app := skapt.Application{
// 		Name: "test",
// 		Handler: func(ctx context.Context) error {
// 			value := ctx.String("url")
// 			c.Assert(value, gc.DeepEquals, "test.com")
// 			return nil
// 		},
// 		Flags: flag.Flags{
// 			{Short: "u", Long: "url", Type: argument.String},
// 		},
// 		Type: argument.Bool,
// 	}
//
// 	err := app.Exec([]string{"test", "-u", "test.com"})
// 	c.Assert(err, gc.IsNil)
//
// 	err = app.Exec([]string{"test", "--url=test.com"})
// 	c.Assert(err, gc.IsNil)
//
// 	app.Type = argument.String
// 	err = app.Exec([]string{"test", "-u", "test.com", "somevalue"})
// 	c.Assert(err, gc.IsNil)
//
// 	app.Type = argument.Int
// 	err = app.Exec([]string{"test", "-u", "test.com", "1"})
// 	c.Assert(err, gc.IsNil)
// }
//
// func (a appSuite) TestExecWithCommands(c *gc.C) {
// 	app := skapt.Application{
// 		Name: "test",
// 		Handler: func(ctx context.Context) error {
// 			value := ctx.String("url")
// 			c.Assert(value, gc.DeepEquals, "test.com")
// 			return nil
// 		},
// 		Flags: flag.Flags{
// 			{Short: "u", Long: "url", Type: argument.String},
// 		},
// 		Type: argument.Bool,
// 		Commands: command.Commands{{
// 			Name: "subtest",
// 			Type: argument.String,
// 			Handler: func(ctx context.Context) error {
// 				value := ctx.String("nik")
// 				c.Assert(value, gc.DeepEquals, "testvalue")
// 				return nil
// 			},
// 			Flags: flag.Flags{{Long: "nik", Type: argument.String}},
// 		}},
// 	}
//
// 	args := []string{"./test", "-u", "test.com", "subtest", "--nik=testvalue"}
// 	err := app.Exec(args)
// 	c.Assert(err, gc.IsNil)
//
// 	args = []string{"./test", "-u", "test.com", "subtest", "--nik=testvalue", "somevalue"}
// 	err = app.Exec(args)
// 	c.Assert(err, gc.IsNil)
//
// 	args = []string{"./test", "-u", "test.com", "subtest", "--nik=testvalue", "1"}
// 	app.Commands[0].Type = argument.Int
// 	err = app.Exec(args)
// 	c.Assert(err, gc.IsNil)
// }
//
// func (a appSuite) TestExecDefaultFlags(c *gc.C) {
// 	app := skapt.Application{
// 		Name: "test",
// 		Handler: func(ctx context.Context) error {
// 			return nil
// 		},
// 		Type: argument.Bool,
// 	}
//
// 	args := []string{"./test -v"}
// 	err := app.Exec(args)
// 	c.Assert(err, gc.IsNil)
//
// 	args = []string{"./test -h"}
// 	err = app.Exec(args)
// 	c.Assert(err, gc.IsNil)
//
// 	args = []string{"./test --version"}
// 	err = app.Exec(args)
// 	c.Assert(err, gc.IsNil)
//
// 	args = []string{"./test --help"}
// 	err = app.Exec(args)
// 	c.Assert(err, gc.IsNil)
// }

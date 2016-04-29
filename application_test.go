package Skapt

import (
	"fmt"
	"os"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type SkaptSuite struct{}

var _ = Suite(&SkaptSuite{})

func (s *SkaptSuite) TestAppFlag(c *C) {
	os.Args = []string{"", "-f", "--move", "10", "-s", "this is cool"}

	app := NewApp()
	app.SetName("Skapt")
	app.SetUsage("App flag base")
	app.SetDescription("Example of flag pattern base app")
	app.SetVersion(false, "1.0.0")
	app.SetAuthors([]string{"Hoenir", "Vili", "Skapt"})
	app.AppendNewOption("-f", "--force", "Force message", BOOL, func() {
		fmt.Println("Force flag parsed!")
	})
	app.AppendNewOption("-m", "--move", "Move command instruction", INT, func() {
		fmt.Println("Move command instruction parsed ", app.Int("--move"))
	})
	app.AppendNewOption("-s", "--stringy", "Stringy this msg", STRING, func() {
		if len(app.String("-s")) > 0 {
			fmt.Println("stringy")
		} else {
			fmt.Println("not Stringy")
		}
	})

	c.Assert(app.Name(), Equals, "Skapt")
	c.Assert(app.Usage(), Equals, "App flag base")
	c.Assert(app.Description(), Equals, "Example of flag pattern base app")
	c.Assert(app.Version(), Equals, "1.0.0")
	c.Assert(app.Authors(), DeepEquals, []string{"Hoenir", "Vili", "Skapt"})
	c.Assert(app.Args(), DeepEquals, []string{"-f", "--move", "10", "-s", "this is cool"})
	c.Assert(app.Bool("--help"), Equals, false)
	c.Assert(app.Bool("-f"), Equals, true)
	c.Assert(app.String("-s"), Not(Equals), "")
	opts := app.Options()

	c.Assert(opts, Not(IsNil))
	c.Assert(opts[0].Name(), Equals, "-f")
	c.Assert(opts[0].Alias(), Equals, "--force")
	c.Assert(opts[0].Description(), Equals, "Force message")
	c.Assert(opts[0].TypeFlag(), Equals, BOOL)
	c.Assert(opts[0].action, Not(IsNil))
	c.Assert(opts[1].Name(), Equals, "-m")
	c.Assert(opts[1].Alias(), Equals, "--move")
	c.Assert(opts[1].Description(), Equals, "Move command instruction")
	c.Assert(opts[1].TypeFlag(), Equals, INT)
	c.Assert(opts[1].action, Not(IsNil))
	c.Assert(opts[2].Name(), Equals, "-s")
	c.Assert(opts[2].Alias(), Equals, "--stringy")
	c.Assert(opts[2].Description(), Equals, "Stringy this msg")
	c.Assert(opts[2].TypeFlag(), Equals, STRING)
	c.Assert(opts[2].action, Not(IsNil))
	app.Run()
}

func (s *SkaptSuite) TestAppCommands(c *C) {
	os.Args = []string{"", "init", "-f", "-m", "100"}
	app := NewApp()
	app.SetName("Skapt")
	app.SetUsage("App command base")
	app.SetDescription("Example of command pattern base app")
	app.SetVersion(false, "1.0.0")
	app.SetAuthors([]string{"Hoenir", "Vili", "Skapt"})
	app.AppendNewCommand(
		"init",
		"init the app with the .conf filed",
		"in order to start coding you need to init it first",
		[][]string{
			{
				"-f",
				"--force",
				"Force message",
				"BOOL",
			},
			{
				"-m",
				"--move",
				"Move Command instruction",
				"INT",
			},
			{
				"-s",
				"--stringy",
				"Stringy this msg",
				"STRING",
			},
		},
		[]Handler{
			func() {
				fmt.Println("Force message parsed!")
			},
			func() {
				fmt.Println("Move command instruction", app.Int("--move"))
			},
			func() {
				if len(app.String("-s")) > 0 {
					fmt.Println("stringy")
				} else {
					fmt.Println("not Stringy")
				}
			},
		})

	c.Assert(app.Bool("--help"), Equals, false)
	c.Assert(app.Bool("-f"), Equals, true)
	c.Assert(app.Name(), Equals, "Skapt")
	c.Assert(app.Usage(), Equals, "App command base")
	c.Assert(app.Description(), Equals, "Example of command pattern base app")
	c.Assert(app.Version(), Equals, "1.0.0")
	c.Assert(app.Authors(), DeepEquals, []string{"Hoenir", "Vili", "Skapt"})

	cmds := app.Commands()

	c.Assert(cmds, Not(IsNil))
	c.Assert(cmds[0].Name(), Equals, "init")
	c.Assert(cmds[0].Description(), Equals, "init the app with the .conf filed")
	c.Assert(cmds[0].Usage(), Equals, "in order to start coding you need to init it first")

	opts := cmds[0].Options()

	c.Assert(opts, Not(IsNil))
	c.Assert(opts[0].Name(), Equals, "-f")
	c.Assert(opts[0].Alias(), Equals, "--force")
	c.Assert(opts[0].Description(), Equals, "Force message")
	c.Assert(opts[0].TypeFlag(), Equals, BOOL)
	c.Assert(opts[0].action, Not(IsNil))
	c.Assert(opts[1].Name(), Equals, "-m")
	c.Assert(opts[1].Alias(), Equals, "--move")
	c.Assert(opts[1].Description(), Equals, "Move Command instruction")
	c.Assert(opts[1].TypeFlag(), Equals, INT)
	c.Assert(opts[1].action, Not(IsNil))
	c.Assert(opts[2].Name(), Equals, "-s")
	c.Assert(opts[2].Alias(), Equals, "--stringy")
	c.Assert(opts[2].Description(), Equals, "Stringy this msg")
	c.Assert(opts[2].TypeFlag(), Equals, STRING)
	c.Assert(opts[2].action, Not(IsNil))

	app.Run()
}

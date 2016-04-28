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
	os.Args = []string{"", "-f"}

	app := NewApp()
	app.SetName("Skapt")
	app.SetUsage("Command base flag")
	app.SetDescription("Example of command pattern base app")
	app.SetVersion(false, "1.0.0")
	app.SetAuthors([]string{"Hoenir", "Vili", "Skapt"})
	app.AppendNewOption("-f", "--force", "Force message", BOOL, func() {
		fmt.Println("Force flag parsed!")
	})
	app.AppendNewOption("-m", "--move", "Move command instruction", INT, func() {
		fmt.Println("Move command instruction parsed")
	})

	c.Assert(app.Name(), Equals, "Skapt")
	c.Assert(app.Usage(), Equals, "Command base flag")
	c.Assert(app.Description(), Equals, "Example of command pattern base app")
	c.Assert(app.Version(), Equals, "1.0.0")
	c.Assert(app.Authors(), DeepEquals, []string{"Hoenir", "Vili", "Skapt"})
	c.Assert(app.Args(), DeepEquals, []string{"-f"})
	c.Assert(app.Bool("--help"), Equals, false)
	c.Assert(app.Bool("-f"), Equals, true)
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
	app.Run()
}

package Skapt

import (
	"fmt"

	. "gopkg.in/check.v1"
)

func (s *SkaptSuite) TestCommand(c *C) {
	cmd := Command{
		name:        "init",
		description: "this does this",
		usage:       "usage of this",
		options: []Option{
			{
				name:        "-f",
				alias:       "--ff",
				description: "this does this",
				typeFlag:    BOOL,
				action:      func() { fmt.Println("test") },
			},
		},
	}

	c.Assert(cmd.Name(), Equals, "init")
	c.Assert(cmd.Description(), Equals, "this does this")
	c.Assert(cmd.Usage(), Equals, "usage of this")
	c.Assert(cmd.NameOptions(), DeepEquals, []string{"-f"})

	opts := cmd.Options()
	c.Assert(opts[0].Name(), Equals, "-f")
	c.Assert(opts[0].Alias(), Equals, "--ff")
	c.Assert(opts[0].Description(), Equals, "this does this")
	c.Assert(opts[0].TypeFlag(), Equals, BOOL)
	c.Assert(opts[0].action, Not(IsNil))
}

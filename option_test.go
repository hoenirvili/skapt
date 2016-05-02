package Skapt

import (
	"fmt"

	. "gopkg.in/check.v1"
)

func (s *SkaptSuite) TestOptions(c *C) {
	opts := []Option{
		{
			name:        "-f",
			alias:       "--ff",
			description: "this does this",
			typeFlag:    BOOL,
			action:      func() { fmt.Println("test1") },
		},
		{
			name:        "-s",
			alias:       "--save",
			description: "save this file",
			typeFlag:    STRING,
			action:      func() { fmt.Println("test2") },
		},
		{
			name:        "-k",
			alias:       "kkk",
			description: "kkk this file",
			typeFlag:    INT,
			action:      func() { fmt.Println("compl1") },
		},
	}

	c.Assert(opts[0].Name(), Equals, "-f")
	c.Assert(opts[0].Alias(), Equals, "--ff")
	c.Assert(opts[0].Description(), Equals, "this does this")
	c.Assert(opts[0].TypeFlag(), Equals, BOOL)
	c.Assert(opts[0].action, Not(IsNil))
	c.Assert(opts[1].Name(), Equals, "-s")
	c.Assert(opts[1].Alias(), Equals, "--save")
	c.Assert(opts[1].Description(), Equals, "save this file")
	c.Assert(opts[1].TypeFlag(), Equals, STRING)
	c.Assert(opts[1].action, Not(IsNil))
	c.Assert(opts[2].Name(), Equals, "-k")
	c.Assert(opts[2].Alias(), Equals, "kkk")
	c.Assert(opts[2].Description(), Equals, "kkk this file")
	c.Assert(opts[2].TypeFlag(), Equals, INT)
	c.Assert(opts[2].action, Not(IsNil))
}

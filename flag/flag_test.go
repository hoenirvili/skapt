package flag_test

import (
	"github.com/hoenirvili/skapt/flag"
	gc "gopkg.in/check.v1"
)

type flagSuite struct{}

var _ = gc.Suite(&flagSuite{})

func (f flagSuite) TestValidate(c *gc.C) {
	flags := []flag.Flag{
		{Short: "u"},
		{Long: "url"},
		{Short: "u", Long: "url"},
	}

	for _, flag := range flags {
		err := flag.Validate()
		c.Assert(err, gc.IsNil)
	}

}

func (f flagSuite) TestValidateWithErrors(c *gc.C) {
	flags := []flag.Flag{
		{},
		{Short: "u", Long: "u"},
	}

	for _, flag := range flags {
		err := flag.Validate()
		c.Assert(err, gc.NotNil)
	}
}

func (f flagSuite) TestString(c *gc.C) {
	flags := []flag.Flag{
		{},
		{Short: "u"},
		{Long: "kill"},
		{Short: "f", Long: "full"},
	}

	strs := []string{
		"",
		"-u",
		"--kill",
		"-f --full",
	}
	for key, flag := range flags {
		str := flag.String()
		c.Assert(str, gc.DeepEquals, strs[key])
	}
}

func (f flagSuite) TestIs(c *gc.C) {
	flags := []flag.Flag{
		{},
		{Short: "u"},
		{Long: "url"},
		{Short: "u", Long: "url"},
	}

	expected := []bool{false, true, false, true}
	for key, flag := range flags {
		got := flag.Is("u")
		c.Assert(got, gc.Equals, expected[key])
	}

	expected = []bool{false, false, false, false}
	for key, flag := range flags {
		got := flag.Is("")
		c.Assert(got, gc.Equals, expected[key])
	}
}

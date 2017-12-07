package flag_test

import (
	"github.com/hoenirvili/skapt/flag"
	gc "gopkg.in/check.v1"
)

type flagSuite struct{}

var _ = gc.Suite(&flagSuite{})

func (f flagSuite) TestValidateErrors(c *gc.C) {
	flag := flag.Flag{}
	err := flag.Validate()
	c.Assert(err, gc.NotNil)
}

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

func (f flagSuite) TestString(c *gc.C) {
	flags := []flag.Flag{
		{},
		{Short: "u"},
		{Long: "url"},
		{Short: "u", Long: "url"},
	}

	strs := []string{
		"",
		"-u",
		"--url",
		"-u --url",
	}
	for key, flag := range flags {
		str := flag.String()
		c.Assert(str, gc.DeepEquals, strs[key])
	}
}

func (f flagSuite) TestEq(c *gc.C) {
	flags := []flag.Flag{
		{},
		{Short: "u"},
		{Long: "url"},
		{Short: "u", Long: "url"},
	}
	expected := []bool{false, true, false, true}
	for key, flag := range flags {
		got := flag.Eq("u")
		c.Assert(got, gc.Equals, expected[key])
	}

	flag := flags[3]
	got := flag.Eq("")
	c.Assert(got, gc.Equals, false)
}

func (f flagSuite) TestIsFlag(c *gc.C) {
	args := []string{
		"", "jfj21", "  dd ",
		"u", "--url", "-u",
		"-u-d", "-kk@@@", "-k d",
	}
	expected := []bool{
		false, false, false,
		false, true, true,
		true, true, false,
	}

	for key, arg := range args {
		got := flag.Valid(arg)
		c.Assert(got, gc.DeepEquals, expected[key])
	}
}

package flag_test

import (
	"github.com/hoenirvili/skapt/argument"
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

func (f flagsSuite) TestBool(c *gc.C) {
	flags := flag.Flags{}
	got := flags.Bool("")
	c.Assert(got, gc.Equals, false)

	flags = flag.Flags{{Short: "u", Long: "l"}}
	got = flags.Bool("u")
	c.Assert(got, gc.Equals, false)
}

func (f flagsSuite) TestInt(c *gc.C) {
	flags := flag.Flags{}
	got := flags.Int("")
	c.Assert(got, gc.Equals, 0)

	flags = flag.Flags{{Short: "u", Long: "l"}}
	got = flags.Int("u")
	c.Assert(got, gc.Equals, 0)
}

func (f flagsSuite) TestString(c *gc.C) {
	flags := flag.Flags{}
	got := flags.String("")
	c.Assert(got, gc.Equals, "")

	flags = flag.Flags{{Short: "u", Long: "l"}}
	got = flags.String("u")
	c.Assert(got, gc.Equals, "")
}

func (f flagsSuite) TestValidate(c *gc.C) {
	flags := f.newFlags()
	err := flags.Validate()
	c.Assert(err, gc.NotNil)

	flags = flags[1:]
	err = flags.Validate()
	c.Assert(err, gc.IsNil)

	flags = flag.Flags{}
	err = flags.Validate()
	c.Assert(err, gc.IsNil)
}

func (f flagsSuite) TestValidateWithError(c *gc.C) {
	flags := flag.Flags{
		{Short: "u", Long: "ul"},
		{Short: "u", Long: "ul"},
		{Short: "ul", Long: "u"},
	}
	err := flags.Validate()
	c.Assert(err, gc.NotNil)
}

var empty flag.Flag

func (f flagsSuite) TestFlag(c *gc.C) {
	flags := f.newFlags()
	fl := flags.Flag("")
	c.Assert(fl, gc.IsNil)

	fl = flags.Flag("u")
	c.Assert(fl, gc.DeepEquals, &flag.Flag{
		Short: "u", Long: "url"})
}

func (f flagsSuite) TestParse(c *gc.C) {
	flags := f.newFlags()[1:]

	test := struct {
		argss [][]string
		ups   [][]string
	}{
		argss: [][]string{
			{"-u", "", "--full", "somevalue", "someothervalue"},
			{"noflag", "", ""},
			{"--unknown"},
			{},
		},
		ups: [][]string{
			{"somevalue", "someothervalue"},
			{"noflag"},
			{"unknown"},
			{},
		},
	}

	for key, args := range test.argss {
		unparsed, err := flags.Parse(args)
		c.Assert(unparsed, gc.DeepEquals, test.ups[key])
		c.Assert(err, gc.IsNil)
	}
}

func (f flagsSuite) TestParseRequired(c *gc.C) {
	flags := f.newFlags()[1:]
	flags[2].Required = true
	args := []string{"-u", "--full", "somevlaue"}
	unparsed, err := flags.Parse(args)
	c.Assert(err, gc.NotNil)
	c.Assert(unparsed, gc.IsNil)

	args = []string{"-l"}
	unparsed, err = flags.Parse(args)
	c.Assert(err, gc.IsNil)
	c.Assert(unparsed, gc.IsNil)
}

func (f flagsSuite) TestValueParse(c *gc.C) {
	flags := flag.Flags{
		{Short: "u", Long: "url", Type: argument.String},
		{Short: "d", Long: "debug", Type: argument.Bool},
		{Short: "t", Long: "times", Type: argument.Int},
	}

	args := []string{"-u", "www.google.com", "-t", "3", "--debug"}
	unparsed, err := flags.Parse(args)
	c.Assert(err, gc.IsNil)
	c.Assert(unparsed, gc.IsNil)
	link := flags.String("u")
	n := flags.Int("t")
	debug := flags.Bool("debug")
	c.Assert(link, gc.Equals, "www.google.com")
	c.Assert(n, gc.Equals, 3)
	c.Assert(debug, gc.Equals, true)
}

func (f flagsSuite) TestParseWithErrors(c *gc.C) {
	flags := f.newFlags()
	flags[0] = flag.Flag{
		Short: "t",
		Long:  "ticks",
	}

	tests := []struct {
		args []string
		t    argument.Type
	}{{[]string{"--ticks=llldsl1iudhaf", "-l"}, argument.Int},
		{[]string{"--ticks=llldsl1iudhaf", "-l"}, argument.Bool},
		{[]string{"--ticks", "other", "-l"}, argument.String},
		{[]string{"--ticks=", "-l"}, argument.Int},
		{[]string{"-t", "-l"}, argument.Int},
		{[]string{"--ticks=", "-l"}, argument.Type(3)},
	}

	for _, test := range tests {
		flags[0].Type = test.t
		unparsed, err := flags.Parse(test.args)
		c.Assert(unparsed, gc.IsNil)
		c.Assert(err, gc.NotNil)
	}
}

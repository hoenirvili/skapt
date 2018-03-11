package argument_test

import (
	"github.com/hoenirvili/skapt/argument"
	gc "gopkg.in/check.v1"
)

type argumentSuite struct{}

var _ = gc.Suite(&argumentSuite{})

// TODO(hoenir): refactor unit tests
func (a argumentSuite) TestShort(c *gc.C) {
	tests := struct {
		args     []string
		expected []bool
	}{
		args:     []string{"a", "hudshf", "-", "-uu", "-k", "k-", "-k-", "-hguhfuighaisudh"},
		expected: []bool{false, false, false, false, true, false, false, false},
	}

	for key, arg := range tests.args {
		got := argument.Short(arg)
		c.Assert(got, gc.Equals, tests.expected[key])
	}
}

func (a argumentSuite) TestLong(c *gc.C) {
	tests := struct {
		args     []string
		expected []bool
	}{
		args:     []string{"a", "--hudshf", "-", "-uu", "--k=", "k-", "-k-", "-hguhfuighaisudh"},
		expected: []bool{false, true, false, false, true, false, false, false},
	}

	for key, arg := range tests.args {
		got := argument.Long(arg)
		c.Assert(got, gc.Equals, tests.expected[key])
	}
}

func (argumentSuite) TestShortTrim(c *gc.C) {
	tests := struct {
		args     []string
		expected []string
	}{
		args:     []string{"a", "hudshf", "-", "uu", "-k", "k-", "-k-", "-hguhfuighaisudh"},
		expected: []string{"a", "hudshf", "-", "uu", "k", "k-", "-k-", "-hguhfuighaisudh"},
	}

	for key, arg := range tests.args {
		got := argument.ShortTrim(arg)
		c.Assert(got, gc.Equals, tests.expected[key])
	}
}

func (a argumentSuite) TestLongTrim(c *gc.C) {
	tests := struct {
		args     []string
		expected []string
	}{
		args:     []string{"a", "--hudshf", "-", "-uu", "--k=", "k-", "-k-", "-hguhfuighaisudh"},
		expected: []string{"a", "hudshf", "-", "-uu", "k", "k-", "-k-", "-hguhfuighaisudh"},
	}

	for key, arg := range tests.args {
		arg, value := argument.LongTrim(arg)
		c.Assert(value, gc.Equals, "")
		c.Assert(arg, gc.Equals, tests.expected[key])
	}

	arg, value := argument.LongTrim("--flag=value")
	c.Assert(arg, gc.Equals, "flag")
	c.Assert(value, gc.Equals, "value")
}

func (a argumentSuite) TestNewValue(c *gc.C) {
	types := []argument.Type{argument.Bool, argument.String, argument.Int}
	value := "value"

	for _, t := range types {
		v := argument.NewValue(value, t)
		c.Assert(v, gc.NotNil)
	}
}

func (a argumentSuite) TestValueParse(c *gc.C) {
	tests := []struct {
		t argument.Type
		v string
	}{
		{t: argument.Bool, v: ""},
		{t: argument.String, v: "stringvalue"},
		{t: argument.Int, v: "3"},
	}

	for _, test := range tests {
		v := argument.NewValue(test.v, test.t)
		c.Assert(v, gc.NotNil)
		err := v.Parse()
		c.Assert(err, gc.IsNil)
	}
}

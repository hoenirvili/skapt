package argument_test

import (
	"github.com/hoenirvili/skapt/argument"
	gc "gopkg.in/check.v1"
)

type argumentSuite struct{}

var _ = gc.Suite(&argumentSuite{})

func (a argumentSuite) TestShort(c *gc.C) {
	tests := []struct {
		arg   string
		value bool
	}{
		{"a", false},
		{"--hudshf", false},
		{"-", false},
		{"-uu", false},
		{"--k=", false},
		{"k-", false},
		{"-k-", false},
		{"-hguhfuighaisudh", false},
		{"-1-1-2-5", false},
		{"-1", true},
	}

	for _, test := range tests {
		value := argument.Short(test.arg)
		c.Assert(value, gc.Equals, test.value)
	}
}

func (a argumentSuite) TestLong(c *gc.C) {
	tests := []struct {
		arg   string
		value bool
	}{
		{"a", false},
		{"--hudshf", true},
		{"-", false},
		{"-uu", false},
		{"--k=", true},
		{"k-", false},
		{"-k-", false},
		{"-hguhfuighaisudh", false},
		{"--k=iosjdf ijasofij aosidfjoasifj", true},
		{"--k=--312-3--1321-", true},
		{"--k=213 88 2882", true},
	}

	for _, test := range tests {
		value := argument.Long(test.arg)
		c.Assert(value, gc.Equals, test.value)
	}
}

func (a argumentSuite) TestShortTrim(c *gc.C) {
	tests := []struct {
		arg   string
		value string
	}{
		{"a", "a"},
		{"hudshf", "hudshf"},
		{"-", "-"},
		{"uu", "uu"},
		{"-k", "k"},
		{"k-", "k-"},
		{"-k-", "-k-"},
		{"-hguhfuighaisudh", "-hguhfuighaisudh"},
	}

	for _, test := range tests {
		arg := argument.ShortTrim(test.arg)
		c.Assert(arg, gc.Equals, test.value)
	}
}

func (a argumentSuite) TestLongTrim(c *gc.C) {
	tests := []struct {
		arg   string
		value string
	}{
		{"a", "a"},
		{"--hudshf", "hudshf"},
		{"-", "-"},
		{"-uu", "-uu"},
		{"--k=", "k"},
		{"k-", "k-"},
		{"-k-", "-k-"},
		{"-hguhfuighaisudh", "-hguhfuighaisudh"},
	}

	for _, test := range tests {
		arg, value := argument.LongTrim(test.arg)
		c.Assert(value, gc.Equals, "")
		c.Assert(arg, gc.Equals, test.value)
	}

	arg, value := argument.LongTrim("--flag=value")
	c.Assert(arg, gc.Equals, "flag")
	c.Assert(value, gc.Equals, "value")
}

func (a argumentSuite) TestNewValue(c *gc.C) {
	tests := []struct {
		t argument.Type
		v string
	}{
		{argument.Bool, "value"},
		{argument.String, "value"},
		{argument.Int, "value"},
	}

	for _, test := range tests {
		v := argument.NewValue(test.v, test.t)
		c.Assert(v, gc.NotNil)
	}
}

func (a argumentSuite) TestValueParse(c *gc.C) {
	tests := []struct {
		t argument.Type
		v string
	}{
		{argument.Bool, ""},
		{argument.String, "stringvalue"},
		{argument.Int, "3"},
	}

	for _, test := range tests {
		v := argument.NewValue(test.v, test.t)
		c.Assert(v, gc.NotNil)
		err := v.Parse()
		c.Assert(err, gc.IsNil)
	}
}

func (a argumentSuite) TestValueParseWithErrors(c *gc.C) {
	tests := []struct {
		t argument.Type
		v string
	}{
		{argument.Int, "fdsauhfusdihfa"},
	}

	for _, test := range tests {
		v := argument.NewValue(test.v, test.t)
		c.Assert(v, gc.NotNil)
		err := v.Parse()
		c.Assert(err, gc.NotNil)
	}
}

func (a argumentSuite) TestValueBool(c *gc.C) {
	v := argument.NewValue("", argument.Bool)
	c.Assert(v, gc.NotNil)
	err := v.Parse()
	c.Assert(err, gc.IsNil)

	value := v.Bool()
	c.Assert(value, gc.Equals, true)
}

func (a argumentSuite) TestValueBoolWithError(c *gc.C) {
	v := argument.NewValue("", argument.Bool)
	c.Assert(v, gc.NotNil)

	value := v.Bool()
	c.Assert(value, gc.Equals, false)
}

func (a argumentSuite) TestValueInt(c *gc.C) {
	v := argument.NewValue("3", argument.Int)
	c.Assert(v, gc.NotNil)
	err := v.Parse()
	c.Assert(err, gc.IsNil)

	value := v.Int()
	c.Assert(value, gc.Equals, 3)
}

func (a argumentSuite) TestValueIntWithError(c *gc.C) {
	v := argument.NewValue("3", argument.Int)
	c.Assert(v, gc.NotNil)

	value := v.Int()
	c.Assert(value, gc.Equals, 0)

}

func (a argumentSuite) TestValueString(c *gc.C) {
	v := argument.NewValue("string", argument.String)
	c.Assert(v, gc.NotNil)
	err := v.Parse()
	c.Assert(err, gc.IsNil)

	value := v.String()
	c.Assert(value, gc.Equals, "string")
}

func (a argumentSuite) TestValueStringWithError(c *gc.C) {
	v := argument.NewValue("", argument.String)
	c.Assert(v, gc.NotNil)

	value := v.String()
	c.Assert(value, gc.Equals, "")
}

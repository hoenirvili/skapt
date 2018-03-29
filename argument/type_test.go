package argument_test

import (
	"fmt"

	"github.com/hoenirvili/skapt/argument"
	gc "gopkg.in/check.v1"
)

type typeSuite struct{}

var _ = gc.Suite(&typeSuite{})

func (t typeSuite) TestString(c *gc.C) {
	unknown := argument.Type(4)
	tests := []struct {
		arg argument.Type
		str string
	}{
		{argument.Bool, "bool"},
		{argument.Int, "int"},
		{argument.String, "string"},
		{argument.Float, "float"},
		{unknown, "unknown type"},
	}
	for _, test := range tests {
		str := fmt.Sprintf("%s", test.arg)
		c.Assert(str, gc.Equals, test.str)
	}
}

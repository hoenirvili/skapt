package flag_test

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type FlagSuite struct{}

var _ = Suite(&FlagSuite{})

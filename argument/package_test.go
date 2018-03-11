package argument_test

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type ArgumentTest struct{}

var _ = Suite(&ArgumentTest{})

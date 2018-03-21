package flag

import (
	"github.com/hoenirvili/skapt/argument"
)

var ParsedAndRequired = Flag{
	Short:    "s",
	Long:     "some",
	value:    &argument.Value{},
	Type:     argument.String,
	Required: true,
}

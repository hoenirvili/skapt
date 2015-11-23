package Skapt

import "errors"

// Error flags
var (
	errNFlagAlias = errors.New("Inappropriate number of aliases")
)

// C style exit flgs
const (
	EXIT_SUCCESS = iota
	EXIT_FAILURE
)

package Skapt

import (
	"errors"
	"fmt"
	"os"
)

// Error flags
var (
	errNFlagAlias  = errors.New("Inappropriate number of aliases")
	errTINT        = errors.New("Can't parse the int from this flag")
	errUnknownFlag = errors.New("Unknow type flag")
	errFlagDec     = errors.New("The app has no flag named this way")
	errNFlags      = errors.New("The numebr of flags dosen't correspond")
	errDeclFlag    = errors.New("The flag entered dosen't match any flag declared")
)

// C style exit flgs
const (
	SUCCESS = iota
	FAILURE
)

func errOnExit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(FAILURE)
}

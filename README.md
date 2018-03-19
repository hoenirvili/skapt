# skapt                                                                                                                                                                                                                                     
[![Build Status](https://travis-ci.org/hoenirvili/skapt.svg?branch=master)](https://travis-ci.org/hoenirvili/skapt) [![Go Report Card](https://goreportcard.com/badge/github.com/hoenirvili/skapt)](https://goreportcard.com/report/github.com/hoenirvili/skapt) [![GoDoc](https://godoc.org/github.com/hoenirvili/skapt?status.svg)](https://godoc.org/github.com/hoenirvili/skapt) [![Coverage Status](https://coveralls.io/repos/github/hoenirvili/skapt/badge.svg?branch=master)](https://coveralls.io/github/hoenirvili/skapt?branch=master)

### Package for building command line apps in Go

> I was inspired from other command line libraries to do my own implementation in Go.

![experimental](doc/ref.png)


# Api example
```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hoenirvili/skapt"
	"github.com/hoenirvili/skapt/argument"
	"github.com/hoenirvili/skapt/flag"
)

func main() {
	app := skapt.Application{
		Name:        "Example",
		Description: "Example is an example of command line app",
		Version:     "1.0.0",
		Handler: func(ctx *skapt.Context) error {
			w := ctx.Int("wait")
			wait := time.Duration(w)
			exp := ctx.String("e")
			time.Sleep(wait * time.Second)
			_, err := fmt.Fprintf(ctx.Stdout, exp)
			return err
		},
		Flags: flag.Flags{
			{
				Short: "e", Long: "exp",
				Description: "Print something funny",
				Type:        argument.String,
				Required:    true,
			},
			{
				Short: "w", Long: "wait",
				Description: "How many seconds to wait until you print",
				Type:        argument.Int,
			},
		},
	}

	app.Exec(os.Args)
}
```
## Help output
```bash
$ : ./main --help
Usage: Example [OPTIONS] [ARG...]
       Example [ --help | -h | -v | --version ]

Example is an example of command line app

Options:

-e --exp      Print something funny
-w --wait   How many seconds to wait until you print
-h --help     Print out the help menu
-v --version  Print out the version of the program
```

## Version output
```bash
$ : ./main -v
Version 1.0.0
```

## Run output
``` bash
$ : ./main -e "Example text" --wait=3
Example text
```

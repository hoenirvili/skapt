# skapt                                                                                                                                                                                                                                     
[![Build Status](https://travis-ci.org/hoenirvili/skapt.svg?branch=master)](https://travis-ci.org/hoenirvili/skapt) [![Go Report Card](https://goreportcard.com/badge/github.com/hoenirvili/skapt)](https://goreportcard.com/report/github.com/hoenirvili/skapt) [![GoDoc](https://godoc.org/github.com/hoenirvili/skapt?status.svg)](https://godoc.org/github.com/hoenirvili/skapt) [![Coverage Status](https://coveralls.io/repos/github/hoenirvili/skapt/badge.svg?branch=master)](https://coveralls.io/github/hoenirvili/skapt?branch=master)
 [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

### Lightweight package for building command line apps in Go

> I was inspired from other command line libraries to do my own package in Go.

![experimental](doc/ref.png)


# API example

Example of command line with two arguments. One flag is required to be explicitly passed.  

All the code in this example is written in one main.go file.

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
		Flags: flag.Flags{{
			Short: "e", Long: "exp",
			Description: "Print something funny",
			Type:        argument.String,
			Required:	 true,
		}, {
			Short: "w", Long: "wait",
			Description: "How many seconds to wait until you print",
			Type:        argument.Int,
		}},
	}
	app.Exec(os.Args)
}
```

If we try to pass a long argument `--help` than we get this auto-generated output.

By default we append help and version flags if the ```Flags``` slice does not contain any of them.

## Help output
```bash
$ : main --help
Usage: Example [OPTIONS] [ARG...]
       Example [ --help | -h | -v | --version ]

Example is an example of command line app

Options:

-e --exp      Print something funny
-w --wait     How many seconds to wait until you print
-h --help     Print out the help menu
-v --version  Print out the version of the program
```

For the version flag.

## Version output
```bash
$ : main -v
Version 1.0.0
$ : main --version
Version 1.0.0
```

In our ```Application struct``` we declared 2 flags, -e/--exp and -w/--wait, 
if we pass valid values, the program sleeps 3 seconds and outputs the message passed in -e.

## Run output
``` bash
$ : main -e "Example text" --wait=3
Example text
```

## Error output

We treat all basic errors by default.

```bash
$ : main
Option -e --exp is required
```

This also checks if the value passed is valid.

```bash
$ : main -w fjsuiajfsd
Cannot parse value "fjsiadufj" as int
```


# FAQ

1. Why should I use this package instead of a more popular one?

	Really, nothing is stopping you to use one that is more battle tested, a more matured package. 				\
	The reason for developing another command line parsing package was that I often want to write 				\
	a simple lightweight command line app and I don't need more features than some basic flag value checking 	\
	and a simpler flag value retrieval api.


2. Why not standard lib one?

	I didn't like the interface they provided and their flag value retrieval.

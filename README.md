# Skapt                                                                                                                                                                                                                                     
[![Build Status](https://travis-ci.org/hoenirvili/Skapt.svg)](https://travis-ci.org/hoenirvili/Skapt)

### Package for building command line apps in Go

![experimental](doc/ref.png)

> I was inspired from other cli frameworks in Go and for the fun/learning purpose I'm trying to do my own little framework.

**Note** : This package will support the two main command line patterns.


### Documentation 
https://godoc.org/github.com/hoenirvili/Skapt

### Flag
**Flag** pattern is the pattern that executable has flag options for changing its behavior. For example, grep command inherits this pattern.


```go
package main

import (
	"fmt"

	"github.com/hoenirvili/Skapt"
)

func main() {
	app := Skapt.NewApp()
	app.SetName("Skapt")
	app.SetUsage("Flag pattern base app")
	app.SetDescription("Example of flag pattern base app")
	app.SetVersion(false, "1.0.0")
	app.SetAuthors([]string{"Hoenir"})

	app.AppendNewOption(Skapt.OptionParams{
		Name:"-f", 
		Alias:"--full", 
		Description:"fill up the file with spaces", 
		Type: Skapt.BOOL, 
		Action: func() {
			fmt.Println("Fill up with spaces")
		},
	})

	app.Run()
}
```

### Sub-Command
**Sub-Command** pattern is the pattern that executable takes sub-command for change its behavior. Git command is one example for this pattern or node package manager(npm).Git takes push, pull subcommands and as for npm init, start, stop, update, upgrade etc.

```go

package main

import (
	"fmt"

	"github.com/hoenirvili/Skapt"
)

func main() {
	app := Skapt.NewApp()
	app.SetName("Skapt")
	app.SetUsage("Command pattern base app")
	app.SetDescription("Example of command pattern base app")
	app.SetVersion(false, "1.0.0")
	app.SetAuthors([]string{"Hoenir"})

	app.AppendNewCommand(Skapt.CommandParams{
		Name:"Init", 
		Description: "Init the project with a working dir", 
		Usage: "Full usage description",
		Flags:[][]string{
			{
				"-c",
				"--check",
				"Lorem ipsum modicus",
				"BOOL", "",
			},
		},
		Actions: []Skapt.Handler{
			func() {
				fmt.Println("Init")
			},
		}})

	app.Run()
}

```

### Template suppport

#### Flag base

```bash
NAME:	Skapt

USAGE:
	Flag pattern base app

DESCRIPTION:
	Example of flag pattern base app

OPTIONS:

	-f, --full
		fill up the file with spaces


	--help, -h  print out the help message

AUTHORS :
	 Hoenir 
VERSION:
	1.0.0

```

###  Command base

```bash
NAME :  Skapt

USAGE:
        Command pattern base app

DESCRIPTION:
        Example of command pattern base app


COMMANDS:

        Init - Full usage description
                Init the project with a working dir

                -c, --check
                Lorem ipsum modicus



        --help, -h  print out the help message

VERSION:
        1.0.0

AUTHORS:
         Hoenir 
```

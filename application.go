package Skapt

import (
	"os"
	"strconv"
)

// App struct is the block of dataype that will store
// all of the semantic and accessories in order to
// handle the application in a better way incapsulating
// filds of interest
type App struct {
	// App name
	name string
	// Describes the app
	usage string
	// Descibes the app usage
	description string
	// Slice of prefefined options aka flags for the command to parse
	options []Option
	// A slice of predefined commands for the app to exec
	commands []Command
	// Authors of the app
	authors []string
	// Version number of the app
	version Version
	// Application command line arguments
	args []string
}

// Cache all flags in the args attribute of App
func (a *App) initArgs() {
	a.args = os.Args[1:]
}

func (a App) Bool(name string) bool {
	// flag based app
	if a.commands == nil {
		for _, opt := range a.options {
			if (opt.name == name || opt.alias == name) && opt.typeFlag == BOOL {
				return true
			}
		}
	} else {
		//command base app
		if a.options == nil {
			for _, cmd := range a.commands {
				for _, opt := range cmd.options {
					if (opt.name == name || opt.alias == name) && opt.typeFlag == BOOL {
						return true
					}
				}
			}
		}
	}
	return false
}

// wrapper for getTarget
func valueOption(opt Option, args []string, optionName string) string {
	var (
		i       int
		target  string
		argsLen = len(args)
	)

	for i = 0; i < argsLen; i++ {
		if args[i] == opt.name && opt.typeFlag == STRING {
			_, target = getTarget(opt.name, args, i, opt.typeFlag)
			break
		} else {
			if args[i] == opt.alias && opt.typeFlag == STRING {
				_, target = getTarget(opt.alias, args, i, opt.typeFlag)
				break
			}
		}
	} //for
	return target
}

//TODO make string ,int flags
//work with aliases and primary names

// Checks if the flag/command-flag exists
// and returns the value of that target
func getTarget(name string, args []string, i int, typeFlag uint8) (int, string) {
	if args[i] == name {
		switch typeFlag {
		case INT:
			if v, err := atoiWrapper(args[i+1]); err == nil {
				return v, ""
			} else {
				errOnExit(err)
			}
		case STRING:
			return 0, args[i+1]
		}
	}
	return 0, ""
}

// Basic simple wrapper for strConv
// providing custom error output
func atoiWrapper(value string) (int, error) {
	val, err := strconv.Atoi(value)
	//if error
	if err != nil {
		return val, errTINT
	} else {
		return val, nil
	}
}

// New returns a new App instance
// true => sub-command type
// false => flag type
func NewApp() *App {
	var app App
	// init
	app.initArgs()
	//return
	return &app
}

// AppendNewCommand appends a new command to our cli App
func (a *App) AppendNewCommand(name, desc, usg string, flags [][]string, actions []Handler) {
	// flag pattern not intended
	if a.options == nil {
		//new object
		var cmd Command
		// set the  content of obj
		cmd.SetName(name)
		cmd.SetDescription(desc)
		cmd.SetUsage(usg)
		cmd.SetOptionsOfACommand(flags, actions)
		// addthe new command to the slice of commands
		a.commands = append(a.commands, cmd)
	}
}

func (a *App) AppenNewOption(name, alias string, reqflg []string, typeFlag uint8, action Handler) {
	// sub command pattern not intended
	if a.commands == nil {
		var opt Option
		// set the conent of the obj
		opt.SetName(name)
		opt.SetAlias(alias)
		opt.SetRequireFlags(reqflg)
		if action != nil {
			opt.SetAction(action)
		}
		opt.SetTypeFlag(typeFlag)
		a.options = append(a.options, opt)
	}
}

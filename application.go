package Skapt

import "os"

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

func (a App) String(name string) string {
	var (
		//	doubleQuotes byte = 0x22
		//	singleQuote  byte = 0x27
		i int
	)
	// flag based app
	if a.commands == nil {
		for _, val := range a.options {
			if val.name == name || val.alias == name {
				if val.typeFlag == STRING {
					for i = 0; i < len(a.args); i++ {
						if a.args[i] == name {
							// TODO:
							// We need a better way to find if the string
							// is a flag/command-subflag of our app and return
							// the target
							// Note that we need to sanitize our flag
							// 1) [--flag=="/to/smth"]
							// 2) [--flag=='to/smth']
							// 3) [--flag==to/smth]
							// 4) [--flag==] [to/smth]
							// etc.
						}
					}
				}
			}
		}
	} else {
		//command base app
		if a.options == nil {

		}
	}
	return ""
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

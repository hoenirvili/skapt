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

// Bool returns true if the flag is present
// on os.Args/app.Args
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

func (a App) String(name string) string {
	var target string // empty defaul value
	// flag based app
	if a.commands == nil {
		// for every option in app
		for _, opt := range a.options {
			// if we find the flag that means that
			// is declared in our app
			// standard-name
			if (opt.name == name || opt.alias == name) && opt.typeFlag == STRING {
				target, _ = getTarget(opt, a.args)
				break
			}
		}
	} else {
		// command base app
		// TODO make it also for
		// command line oriented mode app
		if a.options == nil {
			for _, cmd := range a.commands {
				for _, opt := range cmd.options {
					if (opt.name == name || opt.alias == name) && opt.typeFlag == STRING {
						target, _ = getTarget(opt, a.args)
						goto end
					}
				}
			}
		} //if
	}
end:
	return target
}

// Int returns the target INT type flag
func (a App) Int(name string) int {
	var (
		target int // 0 default value
	)

	// flag based app
	if a.commands == nil {
		// for every option in app
		for _, opt := range a.options {
			// if we find the flag that means that
			// is declared in our app
			// standard-name
			if (opt.name == name || opt.alias == name) && opt.typeFlag == INT {
				_, target = getTarget(opt, a.args)
				break
			}
		}
	} else {
		// command base app
		// TODO make it also for
		// command line oriented mode app
		if a.options == nil {
			for _, cmd := range a.commands {
				for _, opt := range cmd.options {
					if (opt.name == name || opt.alias == name) && opt.typeFlag == INT {
						_, target = getTarget(opt, a.args)
						goto end
					}
				}
			}
		} //if
	}
end:
	return target
}

// NewApp returns a new App instance
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
		// add new command to the slice of commands
		a.commands = append(a.commands, cmd)
	}
}

// AppendNewOption appends a new option to our cli App
func (a *App) AppendNewOption(name, alias, desc string, typeFlag uint8, action Handler) {
	// sub command pattern not intended
	if a.commands == nil {
		var opt Option
		// set the conent of the obj
		opt.SetName(name)
		opt.SetAlias(alias)
		opt.SetDescription(desc)
		opt.SetTypeFlag(typeFlag)
		if action != nil {
			opt.SetAction(action)
		}
		a.options = append(a.options, opt)
	}
}

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

func getTarget(opt Option, name string, args []string) string {
	var (
		lenArgs = len(args)
		i       int
		target  string
	)

	for i = 0; i < lenArgs; i++ {
		// if it's an alias
		if args[i] == opt.alias {
			target = args[i+1]
			break
		} else {
			// if it's the std-name
			if args[i] == opt.name {
				target = args[i+1]
				break
			}
		}
	}

	return target
}

func (a App) String(name string) string {
	var target string
	// flag based app
	if a.commands == nil {
		// for every option in app
		for _, opt := range a.options {
			// if we find the flag that means that
			// is declared in our app
			// standard-name
			if (opt.name == name || opt.alias == name) && opt.typeFlag == STRING {
				target = getTarget(opt, name, a.args)
				break
			}
		}
	} else {
		// command base app
		// TODO make it also for
		// command line oriented mode app
		if a.options == nil {

		}
	}

	return target
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

package Skapt

import (
	"os"
	"strings"
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

// SetName sets the name of the app
func (a *App) SetName(appName string) {
	a.name = appName
}

// SetUsage func sets the usage description of the app
func (a *App) SetUsage(usgDesc string) {
	a.usage = usgDesc
}

// SetDescription func sets the description of the app
func (a *App) SetDescription(desc string) {
	a.description = desc
}

// SetAuthors sets the authors of the app
func (a *App) SetAuthors(auth []string) {
	a.authors = auth
}

// SetVersion func sets the current version
// from the main VERSION file or hardcoded one
func (a *App) SetVersion(fromFile bool, versNum string) {
	// Set version automated from VERSION file
	if fromFile {
		a.version.loadVersion()
	} else {
		// Or wrie it manually
		s := strings.Split(versNum, ".")
		for i, val := range s {
			switch i {
			case 0:
				a.version.version = val
			case 1:
				a.version.majorRevision = val
			case 2:
				a.version.minorRevision = val
			case 3:
				a.version.fixRevisionDet = val
			}
		}
	}
}

// Name returns the name of the app
func (a App) Name() string {
	return a.name
}

// Usage return the text usage of the app
func (a App) Usage() string {
	return a.usage
}

// Description return the description of the app
func (a App) Description() string {
	return a.description
}

// Version return the versioning number
func (a App) Version() string {
	return a.version.Full()
}

// Authors returns a  slice of authors
func (a App) Authors() []string {
	return a.authors
}

// Options returns all flag options
func (a App) Options() []Option {
	return a.options
}

// Commands returns the slice of declared commands
func (a App) Commands() []Command {
	return a.commands
}

// Args returns the arguments passed on the command line
// This uses os.args but without the first element of the slice[0]
func (a App) Args() []string {
	return a.args
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

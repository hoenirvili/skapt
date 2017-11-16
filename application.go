package Skapt

import (
	"os"
	"strings"
)

// App struct is the block of dataype that will store
// all of the semantic and accessories in order to
// handle the application
type App struct {
	// App name
	name string
	// Describes the app
	usage string
	// Descibes the app usage
	description string
	// Slice of predefined options aka flags for the command to parse
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

// SetDescription function sets the description of the app
func (a *App) SetDescription(desc string) {
	a.description = desc
}

// SetAuthors sets the authors of the app
func (a *App) SetAuthors(auth []string) {
	a.authors = auth
}

// SetVersion func sets the current version
// from the main VERSION file or hard-coded one
func (a *App) SetVersion(fromFile bool, versNum string) {
	// Set version automated from VERSION file
	if fromFile {
		a.version.loadVersion()
	} else {
		// Or write it manually
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

// Version return the version number
func (a App) Version() string {
	return a.version.String()
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

// Bool returns true or false if that flag with that name
// was passed on os.Args
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

// String returns the target string of the name flag
// example String("-f") will return "file.txt" if you
// passed a correct argument for the flag you declared
func (a App) String(name string) string {
	var target string // empty default value
	// flag based app
	if a.commands == nil {
		// for every option in app
		for _, opt := range a.options {
			// if we find the flag that means that
			// is declared in our app standard-name
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
		}
	}
end:
	return target
}

// Int returns the target int of the name flag
// example INT("-nr") will return "6"(as type int) if you
// passed a correct argument for the flag you declared
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
		}
	}
end:
	return target
}

// NewApp returns a new App instance
func NewApp() *App {
	var app App
	app.initArgs()
	return &app
}

type CommandParams struct {
	Name        string
	Description string
	Usage       string
	Flags       [][]string
	Actions     []Handler
}

// AppendNewCommand appends a new command to our cli App
func (a *App) AppendNewCommand(params CommandParams) {
	// flag pattern not intended
	if a.options == nil {
		var cmd Command
		cmd.SetName(params.Name)
		cmd.SetDescription(params.Description)
		cmd.SetUsage(params.Usage)
		cmd.SetOptionsOfACommand(params.Flags, params.Actions)
		a.commands = append(a.commands, cmd)
	}
}

type OptionParams struct {
	Name        string
	Alias       string
	Description string
	Type        uint8
	Action      Handler
}

// AppendNewOption appends a new option to our cli App
func (a *App) AppendNewOption(params OptionParams) {
	// sub command pattern not intended
	if a.commands == nil {
		var opt Option
		opt.SetName(params.Name)
		opt.SetAlias(params.Alias)
		opt.SetDescription(params.Description)
		opt.SetTypeFlag(params.Type)
		if params.Action != nil {
			opt.SetAction(params.Action)
		}
		a.options = append(a.options, opt)
	}
}

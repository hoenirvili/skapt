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
	// Slice of prefefined options aka flags for the command to parse
	description string
	// Descibes the app usage
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

// function that parses Options
func parseOptions(args []string, opts []Option) {
	// cache all option that was executed
	cacheLen := len(args)
	var cacheOpt = make([]uint8, cacheLen)
	var flagCount uint8
	// TODO: major refactoring of the code
	// for every argument in our cli
	for i, arg := range args {
		// for every option in our flag cli
		for _, opt := range opts {
			// if we found a valid option given as arg
			if opt.name == arg || opt.alias == arg {
				// try to find it's dependencys
				if opt.requireFlags == nil {
					// it dosen't have any sort of dependency
					// and just execute the handler
					opt.Exec()
					// cache the process
					cacheOpt[i] = 1
					flagCount++
				} else {
					// we have dependecyflags that the flag handler of the flag
					// TODO:find a way to implement the target flag like --path="to/path/file"
				}
			}
		}
	}
}

// Function that parses subcommands
//TODO: make the func to parse all the commands
func parseCommands() {
}

//TODO: we must make the parssing function
// to execute every command flag / flags
// Run the App
func (a App) Run() {
	// we have filled the args buffer
	if len(a.args) > 0 {
		// we have defined our app tobe flag based
		if a.commands == nil {
			// parse all our args and execute the handlers
			parseOptions(a.args, a.options)
		} else {
			// we have define our app to be sub-command based
			if a.options == nil {
				// parse SubCommand and execute the hadlers of the flags
				parseCommands()
			}
		}
	} else {
		//TODO: make the template sistem to generate all the echo content
		//help_tempalte()
	}
}

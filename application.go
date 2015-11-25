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
func (a *App) initFlags() {
	a.args = os.Args[1:]
}

// New returns a new App instance
// true => sub-command type
// false => flag type
func NewApp() *App {
	var app App
	// init
	app.initFlags()
	//return
	return &app
}

//TODO: we must make the parssing function
// to execute every command flag / flags
// Run the App
func (a *App) Run() {
	if len(a.args) > 0 {
		//		for i, val := range a.args {
		//
		//		}
	} else {
		//help_tempalte()
	}
}

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
	description string
	// Descibes the app usage
	usage string
	// A slice of predefined commands for the app to exec
	commands []Command
	// Version number of the app
	version Version
	// set the type of cli app
	mode bool
	// the application command line arguments
	args []string
}

// Cache all flags in the args attribute of App
func (a *App) initFlags() {
	a.args = os.Args[1:]
}

// New returns a new App instance
// true => sub-command type
// false => flag type
func New(mode bool) *App {
	var app App
	// init
	app.SetMode(mode)
	app.initFlags()
	//return
	return &app
}

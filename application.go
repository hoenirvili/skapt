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
	// A slice of predefined options/flags for the app to interpret
	options []Options
	// Version number of the app
	version Version
	// the application command line arguments
	args []string
}

// Cache all flags in the args attribute of App
func (a *App) initFlags() {
	a.args = os.Args[1:]
}

// New returns a new App instance
func New() *App {
	var app App
	app.initFlags()
	return &app
}

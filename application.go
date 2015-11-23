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

// Check if all flags have one alis or none
func flagNCheck(container [][]string) error {
	for _, val := range container {
		if len(val) < 3 {
			return errNFlagAlias
		}
	}
	return nil
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

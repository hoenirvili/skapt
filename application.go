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

func (a *App) SetCommands(names, descriptions, usages []string, opt [][]Options) {
	var commands = make([]Command, len(names))
	lenComm := len(commands)
	lenOpt := len(opt)

	for i := 0; i < lenComm; i++ {
		commands[i].SetName(names[i])
		commands[i].SetUsage(usages[i])
		commands[i].SetDescription(descriptions[i])
		//options := make([]Options, lenOpt)
		for j := 0; j < lenOpt; j++ {

		}
		// TODO: set options for every command to parse
		//commands[i].SetOptions(optNames, handlers)
	}

	a.commands = commands
}

package Skapt

import (
	"os"
	"strings"
)

///////////////////////////////////////////////////////
//
//				GLOBAL TYPE
//
///////////////////////////////////////////////////////

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

///////////////////////////////////////////////////////
//
//				METHODS
//				SET
///////////////////////////////////////////////////////

// SetName func set's the name of the app
func (a *App) SetName(appName string) {
	a.name = appName
}

// SetUsage func set's the usage description of the app
func (a *App) SetUsage(usgDesc string) {
	a.usage = usgDesc
}

// SetDescription func set's the description of the app
func (a *App) SetDescription(desc string) {
	a.description = desc
}

// SetVersion func set's the current version
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
				break
			case 1:
				a.version.majorRevision = val
				break
			case 2:
				a.version.minorRevision = val
				break
			case 3:
				a.version.fixRevisionDet = val
				break
			}
		}
	}
}

// SetNameOptions func set's all the flags that our app will parse
func (a *App) SetNameOptions(flags []string) {
	var opt = make([]Options, len(flags))
	for i := 0; i < len(flags); i++ {
		opt[i].SetName(flags[i])
	}
	a.options = opt
}

// SetOptionHandlers func set's all the handlers
// for every flag that we have in our app
func (a *App) SetOptionHandlers(handler []FlagFunc) {
	for i := 0; i < len(a.options); i++ {
		a.options[i].handler = handler[i]
	}
}

///////////////////////////////////////////////////////
//
//				METHODS
//				GET
///////////////////////////////////////////////////////

// GetName returns the name of the app
func (a App) GetName() string {
	return a.name
}

//GetUsage return the text usage of the app
func (a App) GetUsage() string {
	return a.usage
}

//GetDescription return the description of the app
func (a App) GetDescription() string {
	return a.description
}

//GetVersion return the versioning number
func (a App) GetVersion() string {
	return a.version.GetVersion()
}

// GetNameOptions func returns all flags
// or if the is not a single flag set it will
// return nil
func (a App) GetNameOptions() []string {
	if len(a.options) > 0 {
		var rOpt = make([]string, len(a.options))
		for i, val := range a.options {
			rOpt[i] = val.name
		}
		// Return the exact options that are set
		return rOpt
	}
	// Return empty string
	return nil
}

// GetArgs returns the arguments passed on the command line
// This uses os.args but without the first element of the slice[0]
func (a App) GetArgs() []string {
	return a.args
}

///////////////////////////////////////////////////////
//
//				New App Instance
//
///////////////////////////////////////////////////////

// New returns a new App instance
func New() *App {
	var app App
	app.initFlags()
	return &app
}

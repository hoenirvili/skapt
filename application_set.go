package Skapt

// SetName func set's the name of the app
import "strings"

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

// Set app mode to be flag or sub-command type
func (a *App) SetMode(mode bool) {
	a.mode = mode
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

// Set command-flags or flags of the applications
// infoContainer stores the name, desciption usage of the command
func (a *App) SetCommandOption(infoContainer [][]string, opt [][]Options) {
	switch a.Mode() {
	case true:
		// TODO: set command and for every command the options
	case false:
		// TODO: set just options for every command to parse
	default:
		panic("Mode is of the app is not set")
	}
}

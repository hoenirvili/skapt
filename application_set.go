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

// SetOptionf cun set's the all app option/flags and their handlers
func (a *App) SetOptions(flags [][]string, actions []Handler) {
	// return the number of lines
	nFlags := len(flags)
	// create a slice of options
	a.options = make([]Option, nFlags)
	// fil the slice
	for i := 0; i < nFlags; i++ {
		a.options[i].SetName(flags[i][0])
		a.options[i].SetAlias(flags[i][1])
		a.options[i].SetRequireFlags(flags[i][2:])
		a.options[i].SetAction(actions[i])
	}
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

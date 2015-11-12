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

// SetOptions func set's all the flags name and handler
func (a *App) SetOptions(name []string, handler []FlagFunc) {
	var opt = make([]Options, len(name))

	if lenOpt := len(name); len(name) == len(handler) {
		for i := 0; i < lenOpt; i++ {
			opt[i].SetName(name[i])
			opt[i].SetHandler(handler[i])
		}
	} else {
		panic("Can't set name and handler options of the app")
	}

	a.options = opt
}

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

// SetOptionf func set's the all app option/flags and their handlers
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

// AppendNewCommand appends a new command to our cli App
func (a *App) AppendNewCommand(name, desc, usg string, flags [][]string, actions []Handler) {
	// flag pattern not intended
	if a.options == nil {
		var cmd Command
		cmd.SetName(name)
		cmd.SetDescription(desc)
		cmd.SetUsage(usg)
		// return the number of lines
		nFlags := len(flags)
		// create a slice of options
		cmd.options = make([]Option, nFlags)
		// fil the slice
		for i := 0; i < nFlags; i++ {
			cmd.options[i].SetName(flags[i][0])
			cmd.options[i].SetAlias(flags[i][1])
			cmd.options[i].SetRequireFlags(flags[i][2:])
			cmd.options[i].SetAction(actions[i])
		}
		a.commands = append(a.commands, cmd)
	}
}

// Set's the authors of the app
func (a *App) SetAuthors(auth []string) {
	a.authors = auth
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

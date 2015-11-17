package Skapt

// SetName func set's the name of the app
import (
	"fmt"
	"os"
	"strings"
)

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

// TODO: Thinking of a better way to implement this clutter

// Set command-flags or flags of the applications
// infoContainer stores the name, desciption usage of the command
func (a *App) SetCommandOption(infoCommands [][]string, optionName []string, handlers []FlagFunc) {
	switch a.Mode() {
	case true:
		// how many command we have
		ncomm := len(infoCommands)
		// each attribute  of the command
		nattr := len(infoCommands[0])

		// test if we have entered a valid number
		// of attributes
		if nattr != 3 {
			fmt.Println("Invalid numver of attributes ! ")
			os.Exit(1)
		}

		// slice of command of how many command we have
		comm := make([]Command, ncomm)

		// for every command we have a namme,usgae,description
		// and options for the command
		// i => number of command
		// j for every description
		for i := 0; i < ncomm; i++ {
			for j := 0; j < nattr; j++ {
				switch j {
				//name of the command
				case 0:
					comm[i].SetName(infoCommands[i][j])
				//the description of the command
				case 1:
					comm[i].SetDescription(infoCommands[i][j])
				// the usage of the command
				case 2:
					comm[i].SetUsage(infoCommands[i][j])
				}
			}
			// TODO: setOption for every command now
		}

	case false:
		// TODO: set just options for every command to parse
	default:
		panic("Mode is of the app is not set")
	}
}

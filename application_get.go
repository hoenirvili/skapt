package Skapt

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

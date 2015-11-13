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

// GetArgs returns the arguments passed on the command line
// This uses os.args but without the first element of the slice[0]
func (a App) GetArgs() []string {
	return a.args
}

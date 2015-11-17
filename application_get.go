package Skapt

// Name returns the name of the app
func (a App) Name() string {
	return a.name
}

// Usage return the text usage of the app
func (a App) Usage() string {
	return a.usage
}

// Description return the description of the app
func (a App) Description() string {
	return a.description
}

// Mode return if the mode is actived or not
func (a App) Mode() bool {
	return a.mode
}

// Version return the versioning number
func (a App) Version() string {
	return a.version.Version()
}

// Args returns the arguments passed on the command line
// This uses os.args but without the first element of the slice[0]
func (a App) Args() []string {
	return a.args
}

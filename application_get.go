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

// Version return the versioning number
func (a App) Version() string {
	return a.version.Full()
}

// Authors returns a  slice of authors
func (a App) Authors() []string {
	return a.authors
}

// Options returns all flag options
func (a App) Options() []Option {
	return a.options
}

func (a App) Commands() []Command {
	return a.commands
}

// Args returns the arguments passed on the command line
// This uses os.args but without the first element of the slice[0]
func (a App) Args() []string {
	return a.args
}

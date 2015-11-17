package Skapt

// Command struct is the struct that wil store
// a subcommand of the application.
type Command struct {
	// Command name
	name string
	// description of the command
	description string
	// usage
	usage string
	// Slice of prefefined options aka flags for the command to parse
	options []Option
}

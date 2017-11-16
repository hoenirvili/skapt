package Skapt

// Command is the struct that wil store
// a subcommand of the application.
type Command struct {
	// Command name
	name string
	// description of the command
	description string
	// usage
	usage string
	// Slice of predefined options aka flags for the command to parse
	options []Option
}

// Name method gets the name of the command
func (c Command) Name() string {
	return c.name
}

// Usage method gets the usage of the command
func (c Command) Usage() string {
	return c.usage
}

// Description method gets the description of the command
func (c Command) Description() string {
	return c.description
}

// Options returns slice of declared options/flags
func (c Command) Options() []Option {
	return c.options
}

// NameOptions func returns all flags
// or if is not a single flag set it will
// return nil
func (c Command) NameOptions() []string {
	if len(c.options) > 0 {
		var rOpt = make([]string, len(c.options))
		for i, val := range c.options {
			rOpt[i] = val.name
		}
		// Return the exact options that are set
		return rOpt
	}
	// Return empty string
	return nil
}

// SetName sets the name of the command
func (c *Command) SetName(commName string) {
	c.name = commName
}

// SetUsage sets the usage of the command
func (c *Command) SetUsage(usg string) {
	c.usage = usg
}

// SetDescription sets description of the command
func (c *Command) SetDescription(desc string) {
	c.description = desc
}

// SetOptionsOfACommand sets the options of a command
func (c *Command) SetOptionsOfACommand(flags [][]string, actions []Handler) {
	// return the number of lines
	nFlags := len(flags)
	// create a slice of options
	c.options = make([]Option, nFlags)
	// fill the slice
	for i := 0; i < nFlags; i++ {
		if len(flags[i][:]) > 3 {
			c.options[i].SetName(flags[i][0])
			c.options[i].SetAlias(flags[i][1])
			c.options[i].SetDescription(flags[i][2])
			c.options[i].SetTypeFlag(stringToFlag(flags[i][3]))
			c.options[i].SetAction(actions[i])
		} else {
			errOnExit(errNFlags)
		}
	}
}

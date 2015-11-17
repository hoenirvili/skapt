package Skapt

// Get the name of the command
func (c Command) Name() string {
	return c.name
}

// Get the usage of the command
func (c Command) Usage() string {
	return c.usage
}

// Get the description of the command
func (c Command) Description() string {
	return c.description
}

// NameOptions func returns all flags
// or if the is not a single flag set it will
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

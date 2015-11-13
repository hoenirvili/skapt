package Skapt

// Get the name of the command
func (c Command) GetName() string {
	return c.name
}

// Get the usage of the command
func (c Command) GetUsage() string {
	return c.usage
}

// Ge the description of the command
func (c Command) GetDescription() string {
	return c.description
}

// GetNameOptions func returns all flags
// or if the is not a single flag set it will
// return nil
func (c Command) GetNameOptions() []string {
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

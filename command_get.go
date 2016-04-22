package Skapt

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

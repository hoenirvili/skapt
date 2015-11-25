package Skapt

// Set the name of the command
func (c *Command) SetName(commName string) {
	c.name = commName
}

// Set the usage of the command
func (c *Command) SetUsage(usg string) {
	c.usage = usg
}

// Set description of the command
func (c *Command) SetDescription(desc string) {
	c.description = desc
}

// Set the options of a command
func (c *Command) SetOptionsOfACommand(flags [][]string, actions []Handler) {
	// return the number of lines
	nFlags := len(flags)
	// create a slice of options
	c.options = make([]Option, nFlags)
	// fil the slice
	for i := 0; i < nFlags; i++ {
		c.options[i].SetName(flags[i][0])
		c.options[i].SetAlias(flags[i][1])
		c.options[i].SetRequireFlags(flags[i][2:])
		c.options[i].SetAction(actions[i])
	}
}

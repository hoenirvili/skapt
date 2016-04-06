package Skapt

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
	//TODO verify length of options
	// return the number of lines
	nFlags := len(flags)
	// create a slice of options
	c.options = make([]Option, nFlags)
	// fil the slice
	for i := 0; i < nFlags; i++ {
		if len(flags[i][:]) > 3 {
			c.options[i].SetName(flags[i][0])
			c.options[i].SetAlias(flags[i][1])
			c.options[i].SetDescription(flags[i][2])
			// TODO transfrom type form string to uint8
			c.options[i].SetTypeFlag(flags[i][3])
			c.options[i].SetAction(actions[i])
		} else {
			errOnExit(errNFlags)
		}
	}
}

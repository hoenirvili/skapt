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

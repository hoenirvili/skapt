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

// SetOptions func set's all the flags name and handler
func (c *Command) SetOptions(name []string, handler []FlagFunc) {
	var opt = make([]Option, len(name))

	if lenOpt := len(name); len(name) == len(handler) {
		for i := 0; i < lenOpt; i++ {
			opt[i].SetName(name[i])
			opt[i].SetHandler(handler[i])
		}
	} else {
		panic("Can't set name and handler options of the app")
	}
	c.options = opt
}

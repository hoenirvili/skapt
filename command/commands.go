package command

type Commands []Command

func (c Commands) Validate() error {
	for _, command := range c {
		if err := command.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (c Commands) Has(name string) bool {
	for _, command := range c {
		if command.Name == name {
			return true
		}
	}

	return false
}

func (c Commands) Names() []string {
	n := len(c)
	names := make([]string, n, n)

	for key, command := range c {
		names[key] = command.Name
	}

	return names
}

func (c Commands) AppendDefault() {
	for _, command := range c {
		command.Flags.AppendDefault()
	}
}

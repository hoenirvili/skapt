package Skapt

// SetName func set's the name of the flag
func (o *Option) SetName(flag string) {
	o.name = flag
}

// Set alias name for the option
func (o *Option) SetAlias(als string) {
	o.alias = als
}

// Set other flag that the option requires
func (o *Option) SetRequireFlags(reqFl []string) {
	o.requireFlags = reqFl
}

// SetHandler func set's the function that will
// execute when the flag is lunched
func (o *Option) SetAction(handlr Handler) {
	o.action = handlr
}

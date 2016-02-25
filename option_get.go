package Skapt

// Name gets option name
func (o Option) Name() string {
	return o.name
}

// Alias gets alias of option
func (o Option) Alias() string {
	return o.alias
}

// RequireFlags gets all dependency flags of options
func (o Option) RequireFlags() []string {
	return o.requireFlags
}

// TypeFlag gets type of option
func (o Option) TypeFlag() uint8 {
	return o.typeFlag
}

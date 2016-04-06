package Skapt

// Name gets option name
func (o Option) Name() string {
	return o.name
}

// Alias gets alias of option
func (o Option) Alias() string {
	return o.alias
}

// TypeFlag gets type of option
func (o Option) TypeFlag() uint8 {
	return o.typeFlag
}

//Description gets description of option
func (o Option) Description() string {
	return o.description
}

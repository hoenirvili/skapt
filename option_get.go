package Skapt

func (o Option) Name() string {
	return o.name
}

func (o Option) Alias() string {
	return o.alias
}

func (o Option) RequireFlags() []string {
	return o.requireFlags
}

func (o Option) TypeFlag() uint8 {
	return o.typeFlag
}

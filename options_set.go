package Skapt

// SetName func set's the name of the flag
func (o *Options) SetName(flag string) {
	o.name = flag
}

// SetHandler func set's the function that will
// execute when the flag is lunched
func (o *Options) SetHandler(handler FlagFunc) {
	o.handler = handler
}

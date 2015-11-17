package Skapt

// FlagFunc main type of funtion
type FlagFunc func()

// Options is the struct that will hold a flag
// and the handler that will execute when
// our app will parse that flag.
type Option struct {
	name    string
	handler FlagFunc
}

// Run run the handler
func (o Option) Run() {
	o.handler()
}

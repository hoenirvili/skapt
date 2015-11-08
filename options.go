package Skapt

///////////////////////////////////////////////////////
//
//				GLOBAL TYPE
//
///////////////////////////////////////////////////////

// FlagFunc main type of funtion
type FlagFunc func()

// Options is the struct that will hold a flag
// and the handler that will execute when
// our app will parse that flag.
type Options struct {
	name    string
	handler FlagFunc
}

///////////////////////////////////////////////////////
//
//				METHODS
//				SET
///////////////////////////////////////////////////////

// SetName func set's the name of the flag
func (o *Options) SetName(flag string) {
	o.name = flag
}

// SetHandler func set's the function that will
// execute when the flag is lunched
func (o *Options) SetHandler(handler FlagFunc) {
	o.handler = handler
}

///////////////////////////////////////////////////////
//
//				METHODS
//
///////////////////////////////////////////////////////

// Run run the handler
func (o Options) Run() {
	o.handler()
}

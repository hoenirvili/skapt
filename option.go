package Skapt

// Handler is the type of action
type Handler func()

// Options is the struct that will hold a flag
// and the handler that will execute when
// our app will parse that flag.
type Option struct {
	name         string
	alias        string
	requireFlags []string
	action       Handler
}

// Run run the handler
func (o Option) Exec() {
	o.action()
}

func (o Option) Name() string {
	return o.name
}

func (o Option) Alias() string {
	return o.alias
}

func (o Option) RequireFlags() []string {
	return o.requireFlags
}

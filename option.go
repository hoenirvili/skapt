package Skapt

// Handler is the type of action
type Handler func()

const (
	BOOL = iota
	STRING
	UNKNOWN
)

// Options is the struct that will hold a flag
// and the handler that will execute when
// our app will parse that flag.
type Option struct {
	name         string
	alias        string
	requireFlags []string
	typeFlag     uint8
	action       Handler
}

// Run run the handler
func (o Option) Exec() {
	o.action()
}

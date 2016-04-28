package Skapt

// Handler is the type of action
type Handler func()

// All flag types of App
const (
	BOOL    uint8 = iota // BOOL flag type
	STRING               // STRING flag type
	INT                  // INT flag type
	UNKNOWN              // UNKNOWN flag type
)

// Option is the struct that will hold a flag
// and the handler that will execute when
// our app will parse that flag.
type Option struct {
	// primary name of flag
	name string
	// alias name of the flag
	alias string
	// description of the flag
	description string
	// type of the flag
	typeFlag uint8
	// the handler when that will be executed
	// when the flag is parsed
	action Handler
}

// Exec run the handler
func (o Option) Exec() {
	o.action()
}

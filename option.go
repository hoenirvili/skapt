package Skapt

// Handler is the type of action
type Handler func()

// All flag types of App
const (
	BOOL    = iota // BOOL flag type
	STRING         // STRING flag type
	INT            // INT flag type
	UNKNOWN        // UNKNOWN flag type
)

// Option is the struct that will hold a flag
// and the handler that will execute when
// our app will parse that flag.
type Option struct {
	name        string
	alias       string
	description string
	typeFlag    uint8
	action      Handler
}

// Exec run the handler
func (o Option) Exec() {
	o.action()
}

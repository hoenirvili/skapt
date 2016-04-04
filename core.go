package Skapt

import "fmt"

// Run the App
func (a App) Run() {

	// we have filled the args buffer
	if len(a.args) > 0 {
		// we have defined our app to be flag based
		if a.commands == nil {
			// parse all our args and execute the handlers
			flagBaseApp(a.args, a.options)
		} else {
			// we have define our app to be sub-command based
			if a.options == nil {
				// parse SubCommand and execute the hadlers of the flags
				commandBaseApp()
			}
		}
	} else {
		fmt.Println("temaplte system")
		//TODO: make the template system to generate all the echo content
		//help_tempalte()
	}
}

// TODO: BIG REFACTORING
func flagBaseApp(args []string, opts []Option) {
	fmt.Println("flag base app")
}

/// Function that parses subcommands
//TODO: make the func to parse all the commands
func commandBaseApp() {
	fmt.Println("command base app")
}

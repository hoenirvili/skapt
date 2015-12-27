package Skapt

import "fmt"

// function that parses Options
func optionBaseApp(args []string, opts []Option) {
	// cache all option that was executed
	cacheLen := len(args)
	var cacheOpt = make([]uint8, cacheLen)
	var nFlagsParsed uint8
	// for every argument in our cli
	for i, arg := range args {
		// for every option in our flag cli
		for _, opt := range opts {
			parseOptions(arg, opt, i, cacheOpt, &nFlagsParsed)
		}
	}

	fmt.Println(args)
	fmt.Println(cacheOpt, nFlagsParsed)
}

// check option and if it's valid execute handler
func parseOptions(arg string, opt Option, i int, cacheOpt []uint8, flagCount *uint8) {
	var flag bool
	// if we found a valid option given as arg
	if (opt.name == arg || opt.alias == arg) && arg != "" {
		// try to find it's dependencys
		if opt.requireFlags == nil {
			// it dosen't have any sort of dependency
			// execute the handler
			flag = exec(opt)
			// cache the process
			if flag {
				cacheOpt[i] = 1
				*flagCount++
			}
		} else {
			// we have dependecyflags that the flag handler of the flag
			// TODO:find a way to implement the target flag like --path="to/path/file"
		}
	}
}

// execute action of type handler
func exec(opt Option) bool {

	if opt.action != nil {
		opt.Exec()
		// flag succesfully parsed and exec
		return true
	}
	// flag that depends on other flag/flags
	return false
}

//TODO: we must make the parssing function
// to execute every command flag / flags
// Run the App
func (a App) Run() {

	// we have filled the args buffer
	if len(a.args) > 0 {
		// we have defined our app to be flag based
		if a.commands == nil {
			// parse all our args and execute the handlers
			optionBaseApp(a.args, a.options)
		} else {
			// we have define our app to be sub-command based
			if a.options == nil {
				// parse SubCommand and execute the hadlers of the flags
				commandBaseApp()
			}
		}
	} else {
		//TODO: make the template system to generate all the echo content
		//help_tempalte()
	}

}

// Function that parses subcommands
//TODO: make the func to parse all the commands
func commandBaseApp() {
	fmt.Println("nothing happening")

}

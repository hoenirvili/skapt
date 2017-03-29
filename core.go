package Skapt

import "fmt"

// Run the App
func (a App) Run() {
	// init parser on stack
	p := parser{}
	// we have args
	if len(a.args) > 0 {
		// we have defined our app to be flag based
		if a.commands == nil {
			// default flag
			a.AppendNewOption(OptionParams{
				Name:        "-h",
				Alias:       "--help",
				Description: "Print out the help message",
				Type:        BOOL,
				Action:      a.echoHelp,
			})
			a.AppendNewOption(OptionParams{
				Name:        "-v",
				Alias:       "--version",
				Description: "Print out the version of the app",
				Type:        BOOL,
				Action:      func() { fmt.Println(a.version) },
			})
			// parse all our args and execute the handlers
			p.flagBaseApp(&a)
		} else {
			// we have define our app to be sub-command based
			if a.options == nil {
				// default flag
				a.AppendNewCommand(CommandParams{
					Name:        "help",
					Description: "",
					Usage:       "Print out the help message",
					Flags:       nil,
					Actions:     []Handler{a.echoHelp},
				})
				a.AppendNewCommand(CommandParams{
					Name:        "version",
					Description: "",
					Usage:       "Print out the version of the app",
					Flags:       nil,
					Actions:     []Handler{func() { fmt.Println(a.version) }},
				})
				// parse SubCommand and execute the hadlers of the flags
				p.commandBaseApp(&a)
			}
		}
	} else {
		a.echoHelp()
	}
}

// parser type used for parsing the command line arguments
// we can deal with a command base app or a flag simple one
type parser struct {
	// checkoOpts slice of checked flags
	checkedOpts []Option
	// indexListUnparsed the specific
	// index for every flag that was been parsed
	indexListUnparsed []int
	// ignoreList holder for target
	// for flag type (int, string)
	ignoreList []int
	// reqList request flags
	reqList []int
}

// commandBaseApp starts parsing the arguments as
// if we are dealing with a command base app rathar
// than a simple flag one
func (p *parser) commandBaseApp(a *App) {
	var cmd *Command
	check := false
	args := a.args
	lenArgs := len(args)
	lenCmds := len(a.commands)
	// parse every flag of any type
	// we need to know what command was passed
	for i := 0; i < lenArgs; i++ {
		for j := 0; j < lenCmds; j++ {
			if args[i] == a.commands[j].name {
				cmd = &a.commands[j]
				break
			}
		}
	}

	// we found our command
	if cmd != nil {
		lenOpts := len(cmd.options)
		opts := cmd.options
		// parse every flag of any type
		for i := 1; i < lenArgs; i++ {
			// reset check to ensure that every flag is transparent
			check = false
			// for every option in our app
			for j := 0; j < lenOpts; j++ {
				// if the flag is declared on our slice of option
				// if the flag was not parssed yet
				// if the flag is not just a target for another flag
				// and the flag has an action
				if p.isStateFullFlag(args, opts, i, j) {
					// put the flag in the list to parse
					p.checkedOpts = append(p.checkedOpts, opts[j])
					// if we have a special flag that requires a target
					switch opts[j].typeFlag {
					case INT, STRING:
						if i+1 < lenArgs {
							if !isOption(opts, args[i+1]) { // test the bound array and if the next args is not an option
								// add the target into ignoreList
								p.ignoreList = append(p.ignoreList, i+1)
							}
						} else {
							// else error message, bound check failed and that means we require a target and target was not passed
							fmt.Printf("The arg %s requires a specific target\n Please check the user manual with -h, --help ", args[i])
							a.echoHelp()
							goto exit_grace
						}
					}
					// after all the process just auth the flag checking it and exit the loop
					check = true
					break
					// if the flag is stateless that means it's just a dependency flag
				} else if p.isStatelessFlag(args, opts, i, j) {
					// we append the flag into our dependency list
					p.reqList = append(p.reqList, i)
					// test if the flag is INT,STRING
					// we must make sure it's passed a valid coresponding target
					// if we have a special flag that requires a target
					switch opts[j].typeFlag {
					case INT, STRING:
						if i+1 < lenArgs {
							if !isOption(opts, args[i+1]) { // test the bound array and if the next args is not an option
								// add the target into ignoreList
								p.ignoreList = append(p.ignoreList, i+1)
							}
						} else {
							// else error message, bound check failed and that means we require a target and target was not passed
							fmt.Printf("The arg %s requires a specific target\n Please check the user manual with -h, --help ", args[i])
							a.echoHelp()
							goto exit_grace
						}
					}
					// after all the proces just auth the flag checking it and exit the loop
					check = true
					break
				}
			} //for
			// if the flag was not checked and was not in our ignoreList been aded yet.
			if !check && !p.existInIgnoreList(i) {
				// add the unparsed/unknow flag into list
				p.indexListUnparsed = append(p.indexListUnparsed, i)
			}
		} //for

		if len(p.indexListUnparsed) > 0 {
			n := len(p.indexListUnparsed)
			for i := 0; i < n; i++ {
				fmt.Printf("Unknown flag %s\n", args[p.indexListUnparsed[i]])
			}
			fmt.Println("Check the user manual with -h, --help flags")
		} else {
			for _, action := range p.checkedOpts {
				action.Exec()
			}
		}
	} else {
		a.echoHelp()
	}

exit_grace:
}

// flagBaseApp start the parses of the flag base
// app for normal flgas
func (p *parser) flagBaseApp(a *App) {
	check := false
	args := a.args
	opts := a.options

	lenOpts := len(opts) // number of options declared
	lenArgs := len(args) // number of args

	// parse every flag of any type
	for i := 0; i < lenArgs; i++ {
		// reset check to ensure that every flag is transparent
		check = false
		// for every option in our app
		for j := 0; j < lenOpts; j++ {
			// if the flag is declared on our slice of option
			// if the flag was not parssed yet
			// if the flag is not just a target for another flag
			// and the flag has an action
			if p.isStateFullFlag(args, opts, i, j) {
				// put the flag in the list to parse
				p.checkedOpts = append(p.checkedOpts, opts[j])
				// if we have a special flag that requires a target
				switch opts[j].typeFlag {
				case INT, STRING:
					if i+1 < lenArgs {
						if !isOption(opts, args[i+1]) { // test the bound array and if the next args is not an option
							// add the target into ignoreList
							p.ignoreList = append(p.ignoreList, i+1)
						}
					} else {
						// else error message, bound check failed and that means we require a target and target was not passed
						fmt.Printf("The arg %s requires a specific target\n Please check the user manual with -h, --help ", args[i])
						a.echoHelp()
						goto exit_grace
					}
				}
				// after all the process just auth the flag checking it and exit the loop
				check = true
				break
				// if the flag is stateless that means it's just a dependency flag
			} else if p.isStatelessFlag(args, opts, i, j) {
				// we append the flag into our dependency list
				p.reqList = append(p.reqList, i)
				// test if the flag is INT,STRING
				// we must make sure it's passed a valid coresponding target
				// if we have a special flag that requires a target
				switch opts[j].typeFlag {
				case INT, STRING:
					if i+1 < lenArgs {
						if !isOption(opts, args[i+1]) { // test the bound array and if the next args is not an option
							// add the target into ignoreList
							p.ignoreList = append(p.ignoreList, i+1)
						}
					} else {
						// else error message, bound check failed and that means we require a target and target was not passed
						fmt.Printf("The arg %s requires a specific target\n Please check the user manual with -h, --help ", args[i])
						a.echoHelp()
						goto exit_grace
					}
				}
				// after all the proces just auth the flag checking it and exit the loop
				check = true
				break
			}
		} //for
		// if the flag was not checked and was not in our ignoreList been aded yet.
		if !check && !p.existInIgnoreList(i) {
			// add the unparsed/unknow flag into list
			p.indexListUnparsed = append(p.indexListUnparsed, i)
		}
	} //for

	if len(p.indexListUnparsed) > 0 {
		n := len(p.indexListUnparsed)
		for i := 0; i < n; i++ {
			fmt.Printf("Unknown flag %s\n", args[p.indexListUnparsed[i]])
		}
		fmt.Println("Check the user manual with -h, --help flags")
	} else {
		for _, action := range p.checkedOpts {
			action.Exec()
		}
	}

exit_grace:
}

// isStateFullFlag tests :
// if the flag is declared on our slice of option
// if the flag was not parssed yet
// if the flag is not just a target for another flag
// and the flag has an action
func (p parser) isStateFullFlag(args []string, opts []Option, i, j int) bool {
	if (args[i] == opts[j].name || args[i] == opts[j].alias) && !p.argsWasParsed(opts[j]) && !p.existInIgnoreList(i) && opts[j].action != nil {
		return true
	}
	return false
}

// isStatelessFlag tests :
// if the flag is declared on our slice of option
// if the flag was not parssed yet
// if the flag is not just a target for another flag
// and the flag has NO action
func (p parser) isStatelessFlag(args []string, opts []Option, i, j int) bool {
	if (args[i] == opts[j].name || args[i] == opts[j].alias) && !p.argsWasParsed(opts[j]) && !p.existInIgnoreList(i) && opts[j].action == nil {
		return true
	}
	return false
}

// argsWasParsed test if the flag was parsed
func (p parser) argsWasParsed(opt Option) bool {
	lenParsed := len(p.checkedOpts)
	for i := 0; i < lenParsed; i++ {
		if opt.name == p.checkedOpts[i].name {
			return true
		}
	}
	return false
}

// isOption test if a arg is matching  a predeclared option
func isOption(opts []Option, s string) bool {
	lopt := len(opts)
	for i := 0; i < lopt; i++ {
		if opts[i].name == s || opts[i].alias == s {
			return true
		}
	}

	return false
}

// existInIngnoreList test if the flag ,
// targets are on the ignoreList
func (p parser) existInIgnoreList(index int) bool {
	lenList := len(p.ignoreList)
	for i := 0; i < lenList; i++ {
		if p.ignoreList[i] == index {
			return true
		}
	}
	return false
}

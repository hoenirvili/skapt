package Skapt

import "fmt"

type parser struct {
	// slice of checked flags
	checkedOpts []Option
	//the specific index for every flag that was been parsed
	indexListUnparsed []int
	// holder for target for flag type (int, string)
	ignoreList []int
	// request flags
	reqList []int
}

func (p *parser) flagBaseApp(args []string, opts []Option) {
	check := false
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
					if i < lenArgs && !isOption(opts, args[i+1]) { // test the bound array and if the next args is not an option
						// add the target into ignoreList
						p.ignoreList = append(p.ignoreList, i+1)
					} else {
						// else error message, bound check failed and that means we require a target and target was not passed
						// TODO template
						fmt.Println("bound check failed or the next arg is option we require a target and the target was not correctly")
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
					if i < lenArgs && !isOption(opts, args[i+1]) { // test the bound array and if the next args is not an option
						// add the target into ignoreList
						p.ignoreList = append(p.ignoreList, i+1)
					} else {
						// else error message, bound check failed and that means we require a target and target was not passed
						// TODO template
						fmt.Println("bound check failed or the next arg is option we require a target and the target was not correctly")
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
		fmt.Println("check manual")
	} else {
		for _, action := range p.checkedOpts {
			action.Exec()
		}
	}
}

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

// test if the flag was parsed
func (p parser) argsWasParsed(opt Option) bool {
	lenParsed := len(p.checkedOpts)
	for i := 0; i < lenParsed; i++ {
		if opt.name == p.checkedOpts[i].name {
			return true
		}
	}
	return false
}

// test if a arg is matching a predeclared option
func isOption(opts []Option, s string) bool {
	lopt := len(opts)
	for i := 0; i < lopt; i++ {
		if opts[i].name == s || opts[i].alias == s {
			return true
		}
	}

	return false
}

// test if the flag , targets are on the ignoreList
func (p parser) existInIgnoreList(index int) bool {
	lenList := len(p.ignoreList)
	for i := 0; i < lenList; i++ {
		if p.ignoreList[i] == index {
			return true
		}
	}
	return false
}

// Run the App
func (a App) Run() {
	// init parser on stack
	p := parser{}
	// we have args
	if len(a.args) > 0 {
		// we have defined our app to be flag based
		if a.commands == nil {
			// parse all our args and execute the handlers
			p.flagBaseApp(a.args, a.options)
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

/// Function that parses subcommands
//TODO: make the func to parse all the commands
func commandBaseApp() {
	fmt.Println("command base app")
}

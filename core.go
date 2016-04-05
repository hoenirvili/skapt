package Skapt

import "fmt"

// Run the App
func (a App) Run() {

	// we have args
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

type parser struct {
	// slice of checked flags
	checkedOpts []Option
	//the specific index for every flag that was been parsed
	indexListUnparsed []int
	// holder for target for flag type (int, string)
	ignoreList []int
}

func flagBaseApp(args []string, opts []Option) {

	var (
		parser parser
		check  = false
	)

	lenOpts := len(opts) // number of options declared
	lenArgs := len(args) // number of args

	// parse every argument BOOL, STRING, INT without dependency
	_ = "breakpoint"
	for i := 0; i < lenArgs; i++ {
		// reset check
		check = false
		for j := 0; j < lenOpts; j++ {
			if (args[i] == opts[j].name || args[i] == opts[j].alias) && !argsWasParsed(opts[j], parser.checkedOpts) && !existInIgnoreList(i, parser.ignoreList) {
				parser.checkedOpts = append(parser.checkedOpts, opts[j])
				switch opts[j].typeFlag {
				case INT, STRING:
					if i < lenArgs { // test the bound array
						parser.ignoreList = append(parser.ignoreList, i+1)
					} else {
						// else error message, bound check failed and that means we require a target and target was not passed
						fmt.Println("error")
					}
				}
				check = true
				break
			} //if
		} //for
		if !check && !existInIgnoreList(i, parser.ignoreList) {
			parser.indexListUnparsed = append(parser.indexListUnparsed, i)
		}
	} //for

	if len(parser.indexListUnparsed) > 0 {
		fmt.Println("check manual")
	} else {
		for _, action := range parser.checkedOpts {
			action.Exec()
		}
	}
}

func argsWasParsed(opt Option, parsed []Option) bool {
	lenParsed := len(parsed)
	for i := 0; i < lenParsed; i++ {
		if opt.name == parsed[i].name {
			return true
		}
	}
	return false
}

func existInIgnoreList(index int, ignoreList []int) bool {
	lenList := len(ignoreList)
	for i := 0; i < lenList; i++ {
		if ignoreList[i] == index {
			return true
		}
	}
	return false
}

/// Function that parses subcommands
//TODO: make the func to parse all the commands
func commandBaseApp() {
	fmt.Println("command base app")
}

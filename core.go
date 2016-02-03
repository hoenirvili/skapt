package Skapt

import (
	"fmt"
	"os"
)

//TODO: we must make the parssing function
// to execute every command flag / flags
// Run the App
func (a App) Run() {

	// we have filled the args buffer
	if len(a.args) > 0 {
		// we have defined our app to be flag based
		if a.commands == nil {
			// parse all our args and execute the handlers
			parser(a.args, a.options)
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

func parser(args []string, opts []Option) {
	var (
		i, j        int
		execHandler []Handler
		targets     []string
		k, l        int
		dependency  []string
		ack         int
		tFlag       bool
	)

	// for every arg in app
	for i < len(args) {
		//for every opt in app
		for j < len(opts) {
			// if arg is in the opt
			if opts[j].name == args[i] || opts[j].alias == args[i] {
				//if opt dosen't have req
				if opts[j].requireFlags == nil {
					tFlag = false
					if opts[j].typeFlag == BOOL {
						execHandler = append(execHandler, opts[j].action)
					} else { //it req a target
						targets = append(targets, args[i+1])
						tFlag = true
						execHandler = append(execHandler, opts[j].action)
					}
				} else { // we have dependency
					dependency = opts[j].requireFlags
					for k < len(args) {
						for l < len(dependency) {
							if dependency[l] == args[k] {
								//test know if depencency is BOOL or need target
								trgts := dependParse(args, dependency, opts)
								targets = append(targets, trgts)
								execHandler = append(execHandler, opts[j].action)
								ack++
							}
							l++
						}
						if ack != len(dependency) {
							fmt.Fprintf(os.Stdout, "\n %s depends on %\n", opts[j].name, dependency)
							goto grace_exit
						}
						k++
					}
				}
			}
			j++
		}
		if tFlag {
			i = i + 2
		} else {
			i++
		}
	}
grace_exit:
}

func dependParse(args []string, dependency []string, opts []Option) string {
	var (
		target string
	)
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(dependency); j++ {
			//test
		}
	}
	return target
}

/// Function that parses subcommands
//TODO: make the func to parse all the commands
func commandBaseApp() {
	fmt.Println("nothing happening")
}

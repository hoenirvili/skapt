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
		i, j, nrTargets, nrHandler int
		execHandler                []Handler
		targets                    []string
		k, l                       int
		dependency                 []string
		ack                        int
		tFlag                      bool
		checkedFlags               []string
	)
	_ = "breakpoint"

	// for every arg in app
	for i < len(args) {
		//for every opt in app
		for j = 0; j < len(opts); j++ {
			// if arg is in the opt
			if opts[j].name == args[i] || opts[j].alias == args[i] {
				//if opt dosen't have req
				if opts[j].requireFlags == nil {
					tFlag = false
					if opts[j].typeFlag == BOOL {
						execHandler = append(execHandler, opts[j].action)
						checkedFlags = append(checkedFlags, args[i])
					} else { //it req a target
						targets = append(targets, args[i+1])
						tFlag = true
						execHandler = append(execHandler, opts[j].action)
						checkedFlags = append(checkedFlags, args[i])
					}
				} else { // we have dependency
					dependency = opts[j].requireFlags
					for k = 0; k < len(args); k++ {
						ack = 0
						for l = 0; l < len(dependency); l++ {
							if dependency[l] == args[k] {
								//test know if depencency is BOOL or need target
								trgts, exe := dependParse(args, dependency[l], opts[j], opts)
								if trgts != "" {
									targets = append(targets, trgts)
								}
								if exe == nil {
									execHandler = append(execHandler, exe)
								}
								ack++
							}
						}
						if ack != len(dependency) {
							fmt.Fprintf(os.Stdout, "\n %s depends on %\n", opts[j].name, dependency)
							goto grace_exit
						}
						checkedFlags = append(checkedFlags, args[k])
					}
				}
			}
		}
		if tFlag {
			i = i + 2
		} else {
			i++
		}
	}

	nrTargets = len(targets)
	nrHandler = len(execHandler)

	fmt.Println(nrTargets)
	fmt.Println(nrHandler)

	//everythings it's fine just exec all
	if nrTargets+nrHandler == len(args) {
		for _, do := range execHandler {
			do()
		}
	}

grace_exit:
}

func dependParse(args []string, dependency string, opt Option, opts []Option) (string, Handler) {
	var (
		target string
		exec   Handler
	)

	for i := 0; i < len(opts); i++ {
		if dependency == opts[i].name || opts[i].alias == dependency {
			for j := 0; j < len(args); i++ {
				if args[j] == opts[i].name || opts[i].alias == args[j] {
					// test if dependency is valid
					// a dependency flag does not require an action
					// just target/presence
					if opts[i].action == nil {
						if opts[i].typeFlag == BOOL {
							exec = opt.action
							target = ""
							goto return_it
						} else {
							exec = nil
							target = args[j+1]
							goto return_it
						}
					} //if
				} //if
			} //for
		}
	}
return_it:
	return target, exec
}

/// Function that parses subcommands
//TODO: make the func to parse all the commands
func commandBaseApp() {
	fmt.Println("nothing happening")
}

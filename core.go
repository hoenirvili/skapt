package Skapt

import (
	"fmt"
	"os"
)

// Run TODO: we must make the parssing function
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

// TODO: BIG REFACTORING
// TODO: ADD alias flag dependency
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

	// for every arg in app
	for i < len(args) {
		//for every opt in app
		for j = 0; j < len(opts); j++ {
			if isAlreadyCheckedTargets(targets, args[i]) {
				break
			}
			// if arg is in the opt
			if opts[j].name == args[i] || opts[j].alias == args[i] {
				//if opt dosen't have req
				tFlag = false
				if opts[j].requireFlags == nil {
					// we have targets
					if opts[j].typeFlag != BOOL {
						targets = append(targets, args[i+1])
						tFlag = true
					}
					// if is not target do the usual append
					execHandler = append(execHandler, opts[j].action)
					checkedFlags = append(checkedFlags, args[i])
					break
					// the flag in order to run requires other flags
				} else {
					// checkif the the depend flag was already checked/parsed
					if isAlreadyCheckedDepend(checkedFlags, args[i]) {
						break
					}
					// copy all depend
					dependency = opts[j].requireFlags
					for l = 0; l < len(dependency); l++ {
						ack = 0
						for k = 0; k < len(args); k++ {
							if dependency[l] == args[k] {
								//test know if depencency is BOOL or need target
								trgts, exe := dependParse(args, dependency[l], opts[j], opts)
								// if the flag is not BOOL and we have a target
								if trgts != "" {
									targets = append(targets, trgts)
								}
								if exe == nil {
									execHandler = append(execHandler, exe)
								}
								ack++
								break
							}
						} // for
						// after we checked for all depend flags and we
						// didn't find it just exit and interupt it(and don't execute anything)
						if ack != len(dependency) {
							// TODO: better error prone message
							fmt.Fprintf(os.Stdout, "\n %s depends on %\n", opts[j].name, dependency)
							goto grace_exit
						}
						// if the depend flag was cought just append it to the checked list
						checkedFlags = append(checkedFlags, dependency[l])
					} //for

					// after all the dependency flag parsed
					// we need to also parse the flag itself
					if opts[j].typeFlag != BOOL {
						targets = append(targets, args[i+1])
						tFlag = true
					}
					execHandler = append(execHandler, opts[j].action)
					checkedFlags = append(checkedFlags, args[i])
					break
				} // else
			} //if
		} //for

		if tFlag {
			i = i + 2
		} else {
			i++
		}
	} //for

	//DEBUG INFO
	nrTargets = len(targets)
	nrHandler = len(execHandler)
	fmt.Print("Nr of targets \t: ")
	fmt.Println(nrTargets)
	fmt.Print("Targets \t: ")
	fmt.Println(targets)
	fmt.Print("Nr of handlers \t: ")
	fmt.Println(nrHandler)
	fmt.Print("execHandlers \t: ")
	fmt.Println(execHandler)
	fmt.Println("\t\tRUNS\t\t")

	//everythings it's fine just exec all
	if nrTargets+nrHandler == len(args) {
		for _, do := range execHandler {
			if do != nil {
				do()
			}
		}
	} else {
		// TODO: better error prone message
		// TODO: add the unknow flags to the message
		fmt.Fprintf(os.Stdout, "\n Unknown flag, for help please -h/--help \n")
	}

grace_exit:
}

func dependParse(args []string, dependency string, opt Option, opts []Option) (string, Handler) {
	var (
		target string
		exec   Handler
	)

	for i := 0; i < len(opts); i++ {
		// find the option that has the dependency
		if dependency == opts[i].name || opts[i].alias == dependency {
			//check every argument if the flag is present in args
			for j := 0; j < len(args); j++ {
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

func isAlreadyCheckedDepend(container []string, depend string) bool {
	for i := 0; i < len(container); i++ {
		if container[i] == depend {
			return true
		}
	}

	return false
}

func isAlreadyCheckedTargets(container []string, target string) bool {
	for i := 0; i < len(container); i++ {
		if container[i] == target {
			return true
		}
	}
	return false
}

/// Function that parses subcommands
//TODO: make the func to parse all the commands
func commandBaseApp() {
	fmt.Println("nothing happening")
}

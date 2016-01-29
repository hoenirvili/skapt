package Skapt

/*
// wrapper for getTarget
func valueOption(opt Option, args []string, optionName string) string {
	var (
		i       int
		target  string
		argsLen = len(args)
	)

	for i = 0; i < argsLen; i++ {
		if args[i] == opt.name && opt.typeFlag == STRING {
			_, target = getTarget(opt.name, args, i, opt.typeFlag)
			break
		} else {
			if args[i] == opt.alias && opt.typeFlag == STRING {
				_, target = getTarget(opt.alias, args, i, opt.typeFlag)
				break
			}
		}
	} //for
	return target
}

//TODO make string ,int flags
//work with aliases and primary names

// Checks if the flag/command-flag exists
// and returns the value of that target
func getTarget(name string, args []string, i int, typeFlag uint8) (int, string) {
	if args[i] == name {
		switch typeFlag {
		case INT:
			if v, err := atoiWrapper(args[i+1]); err == nil {
				return v, ""
			} else {
				errOnExit(err)
			}
		case STRING:
			return 0, args[i+1]
		}
	}
	return 0, ""
}

// Basic simple wrapper for strConv
// providing custom error output
func atoiWrapper(value string) (int, error) {
	val, err := strconv.Atoi(value)
	//if error
	if err != nil {
		return val, errTINT
	} else {
		return val, nil
	}
}

*/

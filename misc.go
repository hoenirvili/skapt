package Skapt

import "strconv"

func getTarget(opt Option, args []string) (string, int) {

	var (
		lenArgs = len(args)
		i       int
		targetS string
		targetI int
		err     error
	)

	// for every argument passed
	for i = 0; i < lenArgs; i++ {
		// if the arg is equal with the name or
		// if the arg is equal with the alas
		if args[i] == opt.alias || args[i] == opt.name {
			if opt.typeFlag == STRING {
				targetS = args[i+1]
			} else {
				if opt.typeFlag == INT {
					targetI, err = atoiWrapper(args[i+1])
				}
			}
			break
		} //if
	}

	//if err had occured
	// handle it
	if err != nil {
		errOnExit(err)
	}

	return targetS, targetI
}

// Basic simple wrapper for strConv
// providing custom error output
func atoiWrapper(value string) (int, error) {
	val, err := strconv.Atoi(value)

	//if error
	if err != nil {
		return val, errTINT
	}

	return val, nil
}

package argument

import "errors"

type Arguments struct {
	// idx is the last identified index in args
	idx int
	// name is the last idtified name in args
	name string
	// args is the underlying original args
	args []string
}

func NewArguments(args []string) *Arguments {
	return &Arguments{
		idx:  -1,
		name: "",
		args: args,
	}
}

func (a *Arguments) ContainsOne(names []string) (bool, error) {
	idx := -1
	argname := ""
	for _, name := range names {
		found := false
		for i, arg := range a.args {
			if arg == name && found {
				return false, errors.New("arguments contains more than one")
			}
			idx, argname = i, name
			found = true
		}
	}

	if idx == -1 {
		return false, nil
	}

	a.idx, a.name = idx, argname

	return true, nil
}

func (a Arguments) FoundAt() (int, error) {
	if a.idx == -1 {
		return -1, errors.New("No name found in arg")
	}
	return a.idx, nil
}

func (a Arguments) FoundName() (string, error) {
	if a.name == "" {
		return "", errors.New("No name found in arg")
	}

	return a.name, nil
}

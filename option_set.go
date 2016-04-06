package Skapt

// SetName func set's the name of the flag
func (o *Option) SetName(flag string) {
	o.name = flag
}

// SetAlias sets alias name for the option
func (o *Option) SetAlias(als string) {
	o.alias = als
}

// SetAction func set's the function that will
// execute when the flag is lunched
func (o *Option) SetAction(handlr Handler) {
	o.action = handlr
}

// SetDescription set's the descript of the flag
// what the flag does
func (o *Option) SetDescription(desc string) {
	o.description = desc
}

// SetTypeFlag sets the type of flag
// INT
// STRING
// INT
func (o *Option) SetTypeFlag(typeOfFlag interface{}) {
	//type assertion
	switch v := typeOfFlag.(type) {

	case uint8:
		switch v {
		case BOOL:
			o.typeFlag = BOOL
			break
		case STRING:
			o.typeFlag = STRING
			break
		case INT:
			o.typeFlag = INT
			break
		default:
			o.typeFlag = UNKNOWN
		}

		break
	case string:
		switch v {

		case "INT":
			o.typeFlag = INT
			break
		case "STRING":
			o.typeFlag = STRING
			break
		case "BOOL":
			o.typeFlag = BOOL
			break
		default:
			o.typeFlag = UNKNOWN
		}

	default:
		o.typeFlag = UNKNOWN

	}

	if o.typeFlag == UNKNOWN {
		errOnExit(errUnknownFlag)
	}
}

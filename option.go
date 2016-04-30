package Skapt

// Handler is the type of action
type Handler func()

// All flag types of App
const (
	BOOL    uint8 = iota // BOOL flag type
	STRING               // STRING flag type
	INT                  // INT flag type
	UNKNOWN              // UNKNOWN flag type
)

// Option is the struct that will hold a flag
// and the handler that will execute when
// our app will parse that flag.
type Option struct {
	// primary name of flag
	name string
	// alias name of the flag
	alias string
	// description of the flag
	description string
	// type of the flag
	typeFlag uint8
	// the handler when that will be executed
	// when the flag is parsed
	action Handler
}

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

// Name gets option name
func (o Option) Name() string {
	return o.name
}

// Alias gets alias of option
func (o Option) Alias() string {
	return o.alias
}

// TypeFlag gets type of option
func (o Option) TypeFlag() uint8 {
	return o.typeFlag
}

//Description gets description of option
func (o Option) Description() string {
	return o.description
}

// Exec run the handler
func (o Option) Exec() {
	o.action()
}

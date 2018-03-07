package parser

import (
	"fmt"

	"github.com/hoenirvili/skapt/argument"
	"github.com/hoenirvili/skapt/command"
	"github.com/hoenirvili/skapt/context"
	"github.com/hoenirvili/skapt/flag"
)

// Parser type that holds the implementation
// of the parsing command line arguments
type Parser struct {
	root command.Command
	subs command.Commands
}

func parse(args []string, cmd *command.Command) error {
	if len(args) == 0 {
		return nil
	}

	for i, arg := range args {
		for j := range cmd.Flags {

			if !cmd.Flags[j].Eq(arg) {
				continue
			}

			// extract the value of the current flag
			value := ""
			if i+1 < len(args) {
				value = args[i+1]
			}

			if err := cmd.Flags[j].ParseValue(value); err != nil {
				return err
			}

			switch cmd.Flags[j].Type {
			case argument.Bool:
				// remove the flag found
				args = append(args[:i], args[i+1:]...)
			case argument.String, argument.Int:
				// remove the flag found and his value
				args = append(args[:i], args[i+1:]...)
				args = append(args[:i], args[i+1:]...)
			default:
				// invalid type
				return fmt.Errorf("skapt: cannot parse flag of type %v", cmd.Flags[j].Type)
			}
		}
	}

	if err := cmd.Flags.AllRequiredParsed(); err != nil {
		return err
	}

	// if we still have args left we should parse those also
	if len(args) != 0 {
		switch cmd.Type != argument.Bool {
		case true:
			if err := cmd.ParseValue(args[0]); err != nil {
				return err
			}
			args = append(args[:0], args[1:]...)
		case false:
			if err := cmd.ParseValue(""); err != nil {
				return err
			}
		}
	}

	if len(args) != 0 {
		return fmt.Errorf("skapt: unknown flags %+v", args)
	}

	return nil
}

// Parse parses the command line arguments and returns all handlers and contexts in order
func (p Parser) Parse(args []string) ([]context.Handler, []context.Context, error) {
	args = p.Strip(args[1:])
	sub, i, ok := p.FindSubCommandIn(args)

	share := make(map[string]interface{})
	contexts := make([]context.Context, 0, 2)
	handlers := make([]context.Handler, 0, 2)

	switch ok {
	case true:
		if sub.Type != argument.Bool && p.root.Required {
			p.root.Required = false
		}
	case false:
		i = len(args)
	}

	if err := parse(args[:i], &p.root); err != nil {
		return nil, nil, err
	}

	contexts = append(contexts, context.New(p.root.Flags, share))
	handlers = append(handlers, p.root.Handler)

	if ok {
		if err := parse(args[i+1:], sub); err != nil {
			return nil, nil, err
		}
		contexts = append(contexts, context.New(sub.Flags, share))
		handlers = append(handlers, sub.Handler)
	}

	return handlers, contexts, nil
}

// Strip stripes all -, and -- from flags and splits their values
func (p Parser) Strip(args []string) []string {
	if len(args) == 0 {
		return args
	}
	strip := make([]string, 0, len(args))
	for _, arg := range args {
		if !flag.Valid(arg) {
			strip = append(strip, arg)
		}

		if argument.Short(arg) {
			strip = append(strip, argument.ShortTrim(arg))
			continue
		}
		if argument.Long(arg) {
			flag, value := argument.LongTrim(arg)
			strip = append(strip, flag, value)
		}
	}

	return strip
}

// New returns a new parsed based on the root command and subcommands provided
func New(root command.Command, subs command.Commands) *Parser {
	return &Parser{
		root: root,
		subs: subs,
	}
}

// FindSubCommandIn searches trough args and returns the first command found
// alongside with the position that is in args and true
func (p Parser) FindSubCommandIn(args []string) (*command.Command, int, bool) {
	n := len(p.subs)
	for key, arg := range args {
		for i := 0; i < n; i++ {
			if p.subs[i].Name == arg {
				return &p.subs[i], key, true
			}
		}
	}

	return nil, -1, false
}

// Package skapt provides a tiny interface
// to create and manage your command line applications
package skapt

import (
	"fmt"
	"os"
	"regexp"
	"text/template"

	"github.com/hoenirvili/skapt/flag"
)

// Application will hold all the information for creating
// and parsing the command line
type Application struct {
	// Name of the command line application
	Name string
	// Usage is the usage of the command line
	Usage string
	// Description holds the description of the command line
	Description string
	// Version is the version of the application
	Version string
	// Flags holds list of the root command
	Flags flag.Flags
	// NArgs minim required value args
	NArgs int
	// Handler is the root main handler
	Handler func(flags flag.Flags, args []string) error
}

// validate if the holds valid information
// in order to be executed by Exec
func (a Application) validate() error {
	if a.Name == "" {
		return fmt.Errorf("skapt: Empty application name")
	}
	if a.Handler == nil {
		return fmt.Errorf("skapt: Empty application handler")
	}

	if a.Flags != nil {
		if err := a.Flags.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// Exec executes the command line based on the args provided
func (a Application) Exec(args []string) error {
	if err := a.validate(); err != nil {
		return err
	}

	if len(args) == 0 {
		return fmt.Errorf("skapt: No arguments to execute")
	}

	a.Flags.AppendHelpIfNotPresent()
	a.Flags.AppendVersionIfNotPreset()

	args, err := a.Flags.Parse(args)
	if err != nil {
		return err
	}

	switch {
	case a.Flags.Bool("help"):
		return a.render(help)
	case a.Flags.Bool("version"):
		return a.render(version)
	}

	if len(args) < a.NArgs {
		return fmt.Errorf("skapt: Need at least %d value args", a.NArgs)
	}

	return a.Handler(a.Flags, args)
}

var help = `
{{if .Usage}}{{.Usage}}{{else}}Usage: {{.Name}} [OPTIONS] [ARG...]{{end}}
       {{.Name}} [ --help | -h | -v | --version ]

{{wrap .Description false}}

Options:
{{range .Flags}}
  {{if.Short}}-{{.Short}}{{end}}{{if .Long}} --{{.Long}}{{end}} 	{{wrap .Description true}}{{end}}
`[1:]

var version = `Version {{.}}`

func (a *Application) render(templ string) error {
	funcMap := template.FuncMap{
		"wrap": func(description string, tab bool) string {
			re := regexp.MustCompile(`(?mi)\S+`)
			if len(description) < 90 {
				return description
			}
			str := ""
			words := re.FindAllString(description, -1)
			for key, word := range words {
				if key != 0 && key%12 == 0 {
					if !tab {
						str += "\n"
						continue
					}
					str += "\n\t\t"
				}
				str += word + " "
			}
			return str
		},
	}

	t, err := template.New("t").Funcs(funcMap).Parse(templ)
	if err != nil {
		return err
	}

	return t.Execute(os.Stdout, a)
}

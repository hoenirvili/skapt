// Package skapt provides a tiny interface
// to create and manage your command line applications
package skapt

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"text/tabwriter"
	"text/template"

	"github.com/hoenirvili/skapt/flag"
)

// Context holds context specific information
// for the current running handler
type Context struct {
	// Flags contains the parsed flags
	flag.Flags
	// Args additional command line arguments
	Args []string
	// Stdout writer to stdout
	Stdout io.Writer
	// Stdout writer to stderr
	Stderr io.Writer
}

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
	Handler func(ctx *Context) error
	// Stdout place to write information that the user
	// needs to know
	Stdout io.Writer
	// Stderr place where all error messages should be written
	Stderr io.Writer
}

// validate if the holds valid information
// in order to be executed by Exec
func (a Application) validate() error {
	if a.Name == "" {
		return fmt.Errorf("empty application name")
	}
	if a.Handler == nil {
		return fmt.Errorf("empty application handler")
	}

	if a.Flags != nil {
		if err := a.Flags.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// Exec executes the command line based on the args provided
func (a Application) Exec(args []string) (err error) {
	if err := a.validate(); err != nil {
		return err
	}

	if a.Stdout == nil {
		a.Stdout = os.Stdout
	}
	if a.Stderr == nil {
		a.Stderr = os.Stderr
	}

	// if there is any error, wrap it into a byte slice
	// and output it to stderr
	defer func() {
		if err != nil {
			_, errFpr := fmt.Fprintf(a.Stderr, err.Error())
			if errFpr != nil {
				err = errFpr
			}
		}
	}()

	if len(args) == 0 {
		return fmt.Errorf("no arguments given")
	}

	a.Flags.AppendHelpIfNotPresent()
	a.Flags.AppendVersionIfNotPreset()

	args, err = a.Flags.Parse(args)
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
		return fmt.Errorf("need at least %d additional arguments", a.NArgs)
	}

	if err := a.Flags.RequiredAreParsed(); err != nil {
		return err
	}

	return a.Handler(&Context{
		Flags:  a.Flags,
		Args:   args,
		Stdout: a.Stdout,
		Stderr: a.Stderr,
	})
}

var help = `
{{if .Usage}}{{.Usage}}\t{{else}}Usage:\t{{.Name}} [OPTIONS] [ARG...]{{end}}
\t{{.Name}} [ --help | -h | -v | --version ]

{{wrap .Description false}}

Options:
{{range .Flags}}
{{if.Short}}-{{.Short}}{{end}} {{if .Long}}--{{.Long}}{{end}}\t\t{{wrap .Description true}}{{end}}
`[1:]

var version = `
Version {{.Version}}
`[1:]

var re = regexp.MustCompile(`(?mi)\S+`)

// render renders the specified template to stdout
func (a *Application) render(templ string) error {
	funcMap := template.FuncMap{
		"wrap": func(description string, tab bool) string {
			if len(description) < 90 {
				return description
			}
			str := ""
			words := re.FindAllString(description, -1)
			for key, word := range words {
				if key != 0 && key%10 == 0 {
					if tab {
						str += "\n\\t\\t"
					} else {
						str += "\n"
					}
				}
				str += word + " "
			}
			return str
		},
	}

	t := template.Must(template.New("t").Funcs(funcMap).Parse(templ))
	buffer := &bytes.Buffer{}
	if err := t.Execute(buffer, a); err != nil {
		return err
	}
	str := buffer.String()
	str = strings.Replace(str, "\\t", "\t", -1)
	w := tabwriter.NewWriter(a.Stdout, 0, 0, 1, ' ', 0)
	if _, err := fmt.Fprintf(w, str); err != nil {
		return err
	}
	return w.Flush()
}

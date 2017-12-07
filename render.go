package skapt

import (
	"fmt"
	"io"
	"text/template"
)

const (
	simpleTemplate = `
Usage: {{ .Name }} [OPTION]... {{ .Argument }}

{{ rpad "" 5 }} {{ .Description }}

Options:
{{range $opt := .Options }}
-{{ $opt.Name }}, --{{ $opt.Alias }} [{{ $opt.PlaceHolder }}] {{ rpad "" 20 }} {{ $opt.Description }}
{{ end }}
-h, --help  print out the help message
`

	commandTemplate = `
NAME :	{{ .Name }}
USAGE:
	{{ .Usage }} {{ .Name }} [OPTION]... COMMAND {{ .Argument }}
DESCRIPTION:
	{{ .Description }}
COMMANDS:
{{ range $cmd := .Commands }}
	{{ $cmd.Name }} - {{ $cmd.Usage }}
		{{ $cmd.Description }}
		{{ range $opt := $cmd.Options }}
		{{ $opt.Name }}, {{ $opt.Alias }} {{ $opt.PlaceHolder }}
		{{ $opt.Description }}
		{{ end }}
{{ end }}
	--help, -h  print out the help message
VERSION:
	{{ .Version }}
AUTHORS:
	{{ range $auth := .Authors }} {{ $auth }} {{ end }}
`

	usageTemplate = `Usage: {{ .Name }} [OPTION]... {{ .Argument }}
`
)

func tmpl(w io.Writer, text string, data interface{}) error {
	t := template.New("t")
	t.Funcs(template.FuncMap{
		"rpad": rpad,
	})
	template.Must(t.Parse(text))
	return t.Execute(w, data)
}

func rpad(s string, padding int) string {
	template := fmt.Sprintf("%ds", padding)
	template = "%-" + template
	return fmt.Sprintf(template, s)
}

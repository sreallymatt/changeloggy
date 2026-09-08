package config

import (
	"fmt"
	"text/template"
)

const defaultDateFormat = "January 2, 2006"

type Format struct {
	DateFormat *string `hcl:"date,optional"`
	Template   *string `hcl:"template,optional"`
}

var defaultTemplate = `## {{ .Version }} ({{ .Date }})
{{ if .Notes }}
{{ range .Notes }}* {{ .Body }} [GH-{{ .PR }}]
{{ end }}{{ end }}{{ range .Headings }}
{{ .Heading }}:

{{ range .Entries }}* {{ .Body }} [GH-{{ .PR }}]
{{ end }}{{ end }}`

func NewDefaultFormat() Format {
	return Format{
		DateFormat: new(defaultDateFormat),
		Template:   new(defaultTemplate),
	}
}

func (f Format) Validate() (errors []error) {
	if _, err := template.New("").Parse(*f.Template); err != nil {
		errors = append(errors, fmt.Errorf("invalid format template: %w", err))
	}
	return
}

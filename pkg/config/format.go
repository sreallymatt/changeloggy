package config

import (
	"fmt"
	"text/template"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
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
		DateFormat: pointer.To(defaultDateFormat),
		Template:   pointer.To(defaultTemplate),
	}
}

func (f Format) Validate() (errors []error) {
	if _, err := template.New("").Parse(*f.Template); err != nil {
		errors = append(errors, fmt.Errorf("invalid format template: %w", err))
	}
	return
}

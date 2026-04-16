package templatehelper

import (
	"bytes"
	"text/template"
)

type EntryData struct {
	PR   int64
	Body string
	Type string
	Kind string
}

type KindData struct {
	Heading string
	Entries []EntryData
}

type ReleaseData struct {
	Version string
	Date    string
	Kinds   []KindData
	Notes   []EntryData
}

func Render(tmpl string, data any) (string, error) {
	t, err := template.New("changelog").Parse(tmpl)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

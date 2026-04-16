package config

import (
	"fmt"
	"regexp"
)

type EntryType struct {
	Name     string  `hcl:"name,label"`
	Example  *string `hcl:"example,optional"`
	Regex    *string `hcl:"regex,optional"`
	Priority *int    `hcl:"priority,optional"`
}

func (t EntryType) Validate(kindName string) (errors []error) {
	if t.Regex != nil {
		if t.Example == nil {
			errors = append(errors, fmt.Errorf("`example` must be defined when `regex` is set (kind: %s, type: %s)", kindName, t.Name))
		}

		if _, err := regexp.Compile(*t.Regex); err != nil {
			errors = append(errors, fmt.Errorf("%w (kind: %s, type: %s)", err, kindName, t.Name))
		}
	}
	return
}

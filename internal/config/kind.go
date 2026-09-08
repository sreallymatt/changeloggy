package config

import (
	"fmt"
	"math"
)

type Kind struct {
	Heading  *string `hcl:"heading,optional"`
	Priority *int    `hcl:"priority,optional"`
	TypesMap map[string]EntryType
	Name     string      `hcl:"name,label"`
	Types    []EntryType `hcl:"type,block"`
}

func (k *Kind) ValidateTypes(names map[string][]string) (e []error) {
	priorities := make(map[int]string)

	for _, t := range k.Types {
		names[t.Name] = append(names[t.Name], k.Name)

		if t.Priority != nil {
			if _, ok := priorities[*t.Priority]; ok {
				e = append(e, fmt.Errorf("duplicate priority (`%d`) defined in kind `%s`", *t.Priority, k.Name))
			}
			priorities[*t.Priority] = k.Name
		}
		e = append(e, t.Validate(k.Name)...)
	}

	return
}

func (k *Kind) BuildTypesMap() {
	m := make(map[string]EntryType)
	for _, t := range k.Types {
		m[t.Name] = t
	}
	k.TypesMap = m
}

func (k *Kind) TypePriority(name string) int {
	t, ok := k.TypesMap[name]
	if !ok {
		return math.MaxInt32
	}
	if t.Priority != nil {
		return *t.Priority
	}
	return math.MaxInt32
}

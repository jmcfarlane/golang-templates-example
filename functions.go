package main

import (
	"strings"
	"text/template"
)

var (
	templateMap = template.FuncMap{
		"Upper": func(s string) string {
			return strings.ToUpper(s)
		},
	}
)

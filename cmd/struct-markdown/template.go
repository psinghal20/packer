package main

import (
	"text/template"
)

type Field struct {
	Name string
	Type string
	Docs string
}

type Struct struct {
	SourcePath string
	Name       string
	Filename   string
	Fields     []Field
}

var structDocsTemplate = template.Must(template.New("structDocsTemplate").
	Parse(`<!-- Code generated from the comments of the {{ .Name }} struct in {{ .SourcePath }}; DO NOT EDIT MANUALLY -->
{{range .Fields}}
-   ` + "`" + `{{ .Name}}` + "`" + ` ({{ .Type }}) - {{ .Docs }}
{{- end -}}`))

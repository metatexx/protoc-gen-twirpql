package genenums

import "text/template"

var enumTemplate = template.Must(template.New("").Parse(`package twirpql

import (
	"context"
	"errors"

	{{ range .Imports }}
	"{{.}}"{{ end }}
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/ast"
)
{{ range .Enums }}
func (ec *executionContext) _{{ .Name }}(ctx context.Context, sel ast.SelectionSet, v *{{.Pkg}}.{{.GoName}}) graphql.Marshaler {
	return graphql.MarshalString((*v).String())
}

func (ec *executionContext) unmarshalInput{{.Name}}(ctx context.Context, v interface{}) ({{.Pkg}}.{{.GoName}}, error) {
	switch v := v.(type) {
	case string:
		intValue, ok := {{.Pkg}}.{{.GoName}}_value[v]
		if !ok {
			return 0, errors.New("unknown value: " + v)
		}
		return {{.Pkg}}.{{.GoName}}(intValue), nil
	}
	return 0, errors.New("wrong type")
}
{{ end }}
`))

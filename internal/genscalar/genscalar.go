package genscalar

import (
	"bytes"
	"go/format"
	"io"
	"text/template"
)

var tmpl = template.Must(template.New("").Parse(templateText))

type data struct {
	Imports []string
	Types   map[string]string
}

// Render renders a scalar
// implementation.
func Render(mp map[string]string, imports map[string]struct{}, out io.Writer) error {
	d := &data{Types: mp}
	for i := range imports {
		d.Imports = append(d.Imports, i)
	}
	var b bytes.Buffer
	err := tmpl.Execute(&b, d)
	if err != nil {
		return err
	}
	bts, err := format.Source(b.Bytes())
	if err != nil {
		return err
	}
	_, err = io.Copy(out, bytes.NewReader(bts))
	return err
}

const templateText = `package twirpql

import (
	"encoding/json"
	"io"
	{{ range .Imports }}
	"{{ . }}"
	{{ end }}
)

{{range $key, $val := .Types}}
type {{$key}} {{$val}}

func (scalar *{{$key}}) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return nil
	}
	return json.Unmarshal([]byte(str), scalar)
}

func (scalar {{$key}}) MarshalGQL(w io.Writer) {
	json.NewEncoder(w).Encode(scalar)
}
{{end}}`

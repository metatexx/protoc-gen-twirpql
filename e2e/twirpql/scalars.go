package twirpql

import (
	"encoding/json"
	"io"

	"marwan.io/protoc-gen-twirpql/e2e"
)

type Translations map[string]*e2e.Word

func (scalar *Translations) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return nil
	}
	return json.Unmarshal([]byte(str), scalar)
}

func (scalar Translations) MarshalGQL(w io.Writer) {
	json.NewEncoder(w).Encode(scalar)
}

type Words map[string]*e2e.Word

func (scalar *Words) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return nil
	}
	return json.Unmarshal([]byte(str), scalar)
}

func (scalar Words) MarshalGQL(w io.Writer) {
	json.NewEncoder(w).Encode(scalar)
}

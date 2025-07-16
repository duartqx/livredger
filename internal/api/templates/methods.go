package templates

import (
	"encoding/json"
	"html/template"
	"slices"
)

func jsonify(v any) template.JS {
	b, err := json.Marshal(v)
	if err != nil {
		return template.JS("{}")
	}
	return template.JS(b)
}

func orEq(key string, values ...string) bool {
	return slices.Contains(values, key)
}

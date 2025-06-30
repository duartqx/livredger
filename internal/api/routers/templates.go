package routers

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

func parseTemplates(templates ...string) *template.Template {
	if len(templates) == 0 {
		panic("É necessário passar ao menos um template para compose")
	}

	return template.Must(template.New("").Funcs(
		template.FuncMap{
			"jsonify": jsonify,
			"orEq":    orEq,
			"sub":     func(a, b int) int { return a - b },
		},
	).ParseFS(templatesFS, templates...))
}

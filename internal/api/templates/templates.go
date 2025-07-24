package templates

import (
	"html/template"
	"io/fs"
)

func Compose(fs fs.FS, templates ...string) *template.Template {
	if len(templates) == 0 {
		panic("É necessário passar ao menos um template para compose")
	}

	return template.Must(template.New("").Funcs(
		template.FuncMap{
			"jsonify": jsonify,
			"orEq":    orEq,
			"sub":     func(a, b int) int { return a - b },
		},
	).ParseFS(fs, templates...))
}

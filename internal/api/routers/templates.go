package routers

import (
	"encoding/json"
	"html/template"
	"slices"

	h "github.com/duartqx/livredger/internal/common/http"
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

func compose(templates ...string) *template.Template {
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

type TemplatesLancamento struct {
	Consulta   *h.Templates
	Comando    *h.Templates
	Detalhes   *h.Templates
	Resultados *template.Template
}

type TemplateRegistry struct {
	Index       *h.Templates
	Lancamentos *TemplatesLancamento
}

var templateRegistry *TemplateRegistry

func ObterTemplateRegistry() *TemplateRegistry {
	if templateRegistry == nil {
		templateRegistry = &TemplateRegistry{
			Index: &h.Templates{
				ComBase: compose("index.html", "nav.html"),
				Partial: nil,
				Error:   nil,
			},
			Lancamentos: &TemplatesLancamento{
				Consulta: &h.Templates{
					ComBase: compose(
						"index.html",
						"nav.html",
						"lancamentos/consulta/lancamento.html",
						"lancamentos/consulta/form.html",
						"lancamentos/consulta/consulta.html",
					),
					Partial: compose(
						"lancamentos/consulta/lancamento.html",
						"lancamentos/consulta/form.html",
						"lancamentos/consulta/consulta.html",
					),
					Error: compose(
						"lancamentos/consulta/form.html",
						"lancamentos/consulta/consulta.html",
					),
				},
				Comando: &h.Templates{
					ComBase: compose(
						"index.html",
						"nav.html",
						"lancamentos/comando/form.html",
						"lancamentos/comando/criar.html",
					),
					Partial: compose(
						"lancamentos/comando/form.html",
						"lancamentos/comando/criar.html",
					),
					Error: nil,
				},
				Detalhes: &h.Templates{
					ComBase: compose(
						"index.html",
						"nav.html",
						"lancamentos/comando/form.html",
						"lancamentos/detalhes/detalhes.html",
					),
					Partial: compose(
						"lancamentos/comando/form.html",
						"lancamentos/detalhes/detalhes.html",
					),
					Error: nil,
				},
			},
		}
	}

	return templateRegistry
}

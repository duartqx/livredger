package routers

import (
	"encoding/json"
	"html/template"

	h "github.com/duartqx/livredger/internal/common/http"
)

func jsonify(v any) template.JS {
	b, err := json.Marshal(v)
	if err != nil {
		return template.JS("{}")
	}
	return template.JS(b)
}

func compose(templates ...string) *template.Template {
	if len(templates) == 0 {
		panic("É necessário passar ao menos um template para compose")
	}

	return template.Must(template.New("").Funcs(
		template.FuncMap{
			"jsonify": jsonify,
		},
	).ParseFS(templatesFS, templates...))
}

type TemplatesLancamento struct {
	Consulta   *h.Templates
	Comando    *h.Templates
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
			},
			Lancamentos: &TemplatesLancamento{
				Consulta: &h.Templates{
					ComBase: compose(
						"index.html",
						"nav.html",
						"lancamentos/consulta/consulta.html",
					),
					Partial: compose(
						"lancamentos/consulta/consulta.html",
					),
				},
				Comando: &h.Templates{
					ComBase: compose(
						"index.html",
						"nav.html",
						"lancamentos/criar.html",
					),
					Partial: compose(
						"lancamentos/criar.html",
					),
				},
				Resultados: compose(
					"lancamentos/consulta/lancamento.html",
					"lancamentos/consulta/resultados.html",
				),
			},
		}
	}

	return templateRegistry
}

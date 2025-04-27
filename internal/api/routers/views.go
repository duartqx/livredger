package routers

import (
	"log"
	"net/http"

	"html/template"
)

type Templates struct {
	ComBase *template.Template
	Partial *template.Template
}

func view(templates *Templates) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if auth != "" {
			log.Println("Authorization", auth)
		}

		if templates.Partial != nil && r.Header.Get("HX-Request") == "true" {
			if err := templates.Partial.ExecuteTemplate(w, "partial", nil); err != nil {
				panic(err)
			}
			return
		}
		if err := templates.ComBase.Execute(w, nil); err != nil {
			panic(err)
		}
	}
}

func viewsRouter() *RouterMap {
	return &RouterMap{
		"GET /{$}": view(
			&Templates{
				ComBase: template.Must(
					template.ParseFS(
						templatesFS,
						"index.html",
						"nav.html",
					),
				),
				Partial: nil,
			},
		),
		"GET /lancamentos": view(
			&Templates{
				ComBase: template.Must(
					template.ParseFS(
						templatesFS,
						"index.html",
						"nav.html",
						"lancamentos/consulta/listar.html",
						"lancamentos/consulta/lancamento.html",
					),
				),
				Partial: template.Must(
					template.ParseFS(
						templatesFS,
						"lancamentos/consulta/listar.html",
						"lancamentos/consulta/lancamento.html",
					),
				),
			},
		),
	}
}

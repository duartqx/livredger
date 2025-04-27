package routers

import (
	"log"
	"net/http"

	h "github.com/duartqx/livredger/internal/common/http"
)

type ViewContext struct {
	Templates *h.Templates
	Data      map[string]any
}

func view(ctx *ViewContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if auth != "" {
			log.Println("Authorization", auth)
		}

		if ctx.Templates.Partial != nil && r.Header.Get("HX-Request") == "true" {
			if err := ctx.Templates.Partial.ExecuteTemplate(w, "partial", ctx.Data); err != nil {
				panic(err)
			}
			return
		}
		if err := ctx.Templates.ComBase.ExecuteTemplate(w, "base", ctx.Data); err != nil {
			panic(err)
		}
	}
}

func viewsRouter() *RouterMap {
	registry := ObterTemplateRegistry()
	return &RouterMap{
		"GET /{$}": view(&ViewContext{
			Templates: registry.Index,
			Data:      map[string]any{"Active": "lancamentos"},
		}),
		"GET /lancamentos": view(&ViewContext{
			Templates: registry.Lancamentos.Consulta,
			Data:      map[string]any{"Active": "lancamentos"},
		}),
		"GET /lancamentos/criar": view(&ViewContext{
			Templates: registry.Lancamentos.Comando,
		}),
	}
}

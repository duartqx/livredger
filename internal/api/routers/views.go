package routers

import (
	"fmt"
	"log"
	"net/http"

	h "github.com/duartqx/livredger/internal/common/http"
	"github.com/duartqx/livredger/internal/common/types"
)

type DataFunc func(usuario *types.Usuario, r *http.Request) (map[string]any, error)

type ViewContext struct {
	ViewName  string
	Templates *h.Templates
	Data      map[string]any
	DataFunc  DataFunc
}

func view(ctx *ViewContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("View", ctx.ViewName)

		var usuario *types.Usuario

		auth := r.Header.Get("Authorization")

		if auth != "" {
			log.Println("Authorization", auth)
		}

		if ctx.Data != nil && ctx.DataFunc != nil {
			panic(fmt.Errorf("Você deve passar somente ViewContext.Data ou ViewContext.DataFunc, mas nunca ambos"))
		}

		var (
			data map[string]any = ctx.Data
			err  error
		)

		if ctx.DataFunc != nil {
			data, err = ctx.DataFunc(usuario, r)
			if err != nil {
				h.JsonErrorReponse(w, err)
				return
			}

		}

		if ctx.Templates.Partial != nil && r.Header.Get("HX-Request") == "true" {
			if err := ctx.Templates.Partial.ExecuteTemplate(w, "partial", data); err != nil {
				panic(err)
			}
			return
		}
		if err := ctx.Templates.ComBase.ExecuteTemplate(w, "base", data); err != nil {
			panic(err)
		}
	}
}

func viewsRouter() *RouterMap {
	registry := ObterTemplateRegistry()
	return &RouterMap{
		"GET /{$}": view(&ViewContext{
			ViewName:  "Index",
			Templates: registry.Index,
			Data:      map[string]any{"Active": "lancamentos"},
		}),
		"GET /lancamentos": view(&ViewContext{
			ViewName:  "ConsultarLancamentos",
			Templates: registry.Lancamentos.Consulta,
			Data:      map[string]any{"Active": "lancamentos"},
		}),
		"GET /lancamentos/criar": view(&ViewContext{
			ViewName:  "CriarLancamento",
			Templates: registry.Lancamentos.Comando,
		}),
	}
}

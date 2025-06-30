package routers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/duartqx/livredger/internal/api/decoders"
	"github.com/duartqx/livredger/internal/api/response"
)

type DataFunc func(r *http.Request) (map[string]any, error)

type ViewContext struct {
	ViewName  string
	Templates *response.Templates
	Data      map[string]any
	DataFunc  DataFunc
}

func View(ctx *ViewContext) http.HandlerFunc {
	if ctx.Data != nil && ctx.DataFunc != nil {
		panic(fmt.Errorf("Você deve passar somente ViewContext.Data ou ViewContext.DataFunc, mas nunca ambos"))
	}

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("View", ctx.ViewName)

		auth := r.Header.Get("Authorization")

		if auth != "" {
			log.Println("Authorization", auth)
		}

		var (
			data map[string]any = ctx.Data
			err  error
		)

		if ctx.DataFunc != nil {
			data, err = ctx.DataFunc(r)
		}

		switch {
		case err == nil && ctx.Templates.Partial != nil && r.Header.Get("HX-Request") == "true":
			if err := ctx.Templates.Partial.ExecuteTemplate(w, "partial", data); err != nil {
				panic(err)
			}
		case err == nil:
			if err := ctx.Templates.ComBase.ExecuteTemplate(w, "base", data); err != nil {
				panic(err)
			}
		case errors.Is(err, decoders.DecoderError) && ctx.Templates.Error != nil:
			w.WriteHeader(400)
			if err := ctx.Templates.Error.ExecuteTemplate(
				w, "error", map[string]any{"Errors": decoders.ParseDecodeError(err)},
			); err != nil {
				panic(err)
			}
		default:
			response.ErrorResponse(w, err)
		}
	}
}

func viewsRouter() *RouterMap {
	templateRegistry := ObterTemplateRegistry()
	return &RouterMap{
		"GET /{$}": View(&ViewContext{
			ViewName:  "Index",
			Templates: templateRegistry.Index,
			Data:      map[string]any{"Active": "Index"},
		}),
	}
}

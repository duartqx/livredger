package routers

import (
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"

	"github.com/duartqx/livredger/internal/api/decoders"
	"github.com/duartqx/livredger/internal/api/response"
	"github.com/duartqx/livredger/internal/api/templates"
)

type DataFunc func(r *http.Request) (map[string]any, error)

type ViewContext struct {
	ViewName string
	Template *template.Template
	Data     map[string]any
	DataFunc DataFunc
}

func ViewsRouter(templFS fs.FS) *RouterMap {
	return &RouterMap{
		"GET /{$}": View(&ViewContext{
			ViewName: "Index",
			Template: templates.Templates(templFS, "index.html", "nav.html"),
			Data:     map[string]any{"Active": "Index"},
		}),
	}
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
		case err == nil && r.Header.Get("HX-Request") == "true":
			if err := ctx.Template.ExecuteTemplate(w, "partial", data); err != nil {
				panic(err)
			}
		case err == nil:
			if err := ctx.Template.ExecuteTemplate(w, "base", data); err != nil {
				panic(err)
			}
		case errors.Is(err, decoders.FormDecoderError) && ctx.Template.Lookup("error") != nil:

			w.WriteHeader(400)

			data = map[string]any{"Errors": decoders.ParseFormDecodeError(err)}

			if err := ctx.Template.ExecuteTemplate(w, "error", data); err != nil {
				panic(err)
			}
		default:
			response.ErrorResponse(w, err)
		}
	}
}

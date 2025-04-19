package routers

import (
	"context"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/duartqx/livredger/internal/presentation/views"
)

func view(c templ.Component) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if auth != "" {
			log.Println("Authorization", auth)
		}

		c.Render(context.Background(), w)
	}
}

func viewsRouter() *RouterMap {
	return &RouterMap{
		"GET /{$}":         view(views.Index()),
		"GET /lancamentos": view(views.Lancamentos()),
	}
}

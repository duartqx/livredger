package routers

import (
	"io/fs"

	"github.com/duartqx/livredger/internal/api/templates"
)

func ViewsRouter(templFS fs.FS) *RouterMap {
	return &RouterMap{
		"GET /{$}": View(&ViewContext{
			ViewName: "Index",
			Template: templates.Templates(templFS, "index.html", "nav.html"),
			Data:     map[string]any{"Active": "Index"},
		}),
	}
}

package api

import (
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	cm "github.com/duartqx/ddgomiddlewares/cors"
	i "github.com/duartqx/ddgomiddlewares/interfaces"
	lm "github.com/duartqx/ddgomiddlewares/logger"
	rm "github.com/duartqx/ddgomiddlewares/recovery"
	tr "github.com/duartqx/ddgomiddlewares/trailling"

	"github.com/duartqx/livredger/internal/api/routers"
)

type Static struct {
	Fs   fs.FS
	Path string
}

type Dependencies struct {
	Static    *[]Static
	Templates fs.FS
}

type RouterMap map[string]http.HandlerFunc

type Mux struct {
	*http.ServeMux
}

func (m *Mux) Group(pattern string, handler http.Handler) error {
	if !strings.HasPrefix(pattern, "/") && !strings.HasSuffix(pattern, "/") {
		return fmt.Errorf("Invalid Router Pattern")
	}

	prefix := strings.TrimSuffix(pattern, "/")

	m.Handle(pattern, http.StripPrefix(prefix, handler))

	return nil
}

func (m *Mux) AddRoutes(rms ...RouterMap) {
	for _, rm := range rms {
		for pattern, router := range rm {
			m.HandleFunc(pattern, router)
		}
	}
}

func (m Mux) Use(mux http.Handler, middlewares ...i.Middleware) http.Handler {
	wrapped := mux
	for _, middleware := range middlewares {
		wrapped = middleware(wrapped)
	}
	return wrapped
}

func Router(dependencies *Dependencies) http.Handler {
	mux := Mux{ServeMux: http.NewServeMux()}

	for _, s := range *dependencies.Static {
		mux.Group(s.Path, http.StripPrefix("/", http.FileServer(http.FS(s.Fs))))
	}

	mux.AddRoutes(
		RouterMap(*routers.LancamentosRouter(dependencies.Templates)),
		RouterMap(*routers.ContasRouter(dependencies.Templates)),
		RouterMap(*routers.ViewsRouter(dependencies.Templates)),
	)

	return mux.Use(
		mux,
		tr.TrailingSlashMiddleware,
		lm.LoggerMiddleware,
		rm.RecoveryMiddleware,
		cm.CorsMiddleware,
	)
}

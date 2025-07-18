package api

import (
	"cmp"
	"context"
	"fmt"
	"io/fs"
	"net/http"
	"strings"
	"time"

	"github.com/duartqx/ddgomiddlewares/cors"
	"github.com/duartqx/ddgomiddlewares/interfaces"
	"github.com/duartqx/ddgomiddlewares/logger"
	"github.com/duartqx/ddgomiddlewares/recovery"
	"github.com/duartqx/ddgomiddlewares/trailling"

	"github.com/duartqx/livredger/internal/api/common"
	"github.com/duartqx/livredger/internal/api/routers"
)

type Static struct {
	Fs   fs.FS
	Path string
}

type Dependencies struct {
	Static         *[]Static
	Templates      fs.FS
	RequestTimeout int
}

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

func (m *Mux) AddRoutes(rms ...*common.RouterMap) {
	for _, rm := range rms {
		for pattern, router := range *rm {
			m.HandleFunc(pattern, router)
		}
	}
}

func (m Mux) Use(mux http.Handler, middlewares ...interfaces.Middleware) http.Handler {
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
		routers.LancamentosRouter(dependencies.Templates),
		routers.ContasRouter(dependencies.Templates),
		routers.ViewsRouter(dependencies.Templates),
	)

	return mux.Use(
		mux,
		func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				ctx, cancel := context.WithTimeout(
					r.Context(),
					time.Second*time.Duration(cmp.Or(dependencies.RequestTimeout, 5)),
				)
				defer cancel()

				next.ServeHTTP(w, r.WithContext(ctx))
			})
		},
		trailling.TrailingSlashMiddleware,
		logger.LoggerMiddleware,
		recovery.RecoveryMiddleware,
		cors.CorsMiddleware,
	)
}

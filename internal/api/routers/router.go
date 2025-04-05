package routers

import (
	"fmt"
	"net/http"
	"strings"

	cm "github.com/duartqx/ddgomiddlewares/cors"
	i "github.com/duartqx/ddgomiddlewares/interfaces"
	lm "github.com/duartqx/ddgomiddlewares/logger"
	rm "github.com/duartqx/ddgomiddlewares/recovery"
	tr "github.com/duartqx/ddgomiddlewares/trailling"
)

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

func (m Mux) Use(mux http.Handler, middlewares ...i.Middleware) http.Handler {
	wrapped := mux
	for _, middleware := range middlewares {
		wrapped = middleware(wrapped)
	}
	return wrapped
}

func Router() http.Handler {
	mux := Mux{ServeMux: http.NewServeMux()}

	for pattern, router := range *LancamentosRouter() {
		mux.HandleFunc(pattern, router)
	}

	return mux.Use(
		mux,
		tr.TrailingSlashMiddleware,
		lm.LoggerMiddleware,
		rm.RecoveryMiddleware,
		cm.CorsMiddleware,
	)
}

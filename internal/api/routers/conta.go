package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/duartqx/livredger/internal/api/response"

	"github.com/duartqx/livredger/internal/application/services/executores"
	"github.com/duartqx/livredger/internal/application/services/visualizadores"

	"github.com/duartqx/livredger/internal/common/mimetypes"
	"github.com/duartqx/livredger/internal/common/types"

	"github.com/duartqx/livredger/internal/domain/comandos"
	"github.com/duartqx/livredger/internal/domain/consultas"
	"github.com/duartqx/livredger/internal/domain/entidade"

	"github.com/duartqx/livredger/internal/infra"
)

func contasRouter() *RouterMap {
	return &RouterMap{
		"GET /api/contas": func(w http.ResponseWriter, r *http.Request) {
			var usuario *types.Usuario

			uow := infra.Bootstrap(usuario)
			defer uow.Close()

			consulta := consultas.ConsultaContasPadrao()

			resultado, err := visualizadores.BuscarContas(uow, consulta)

			if err != nil {
				response.JsonErrorResponse(w, err)
				return
			}

			w.Header().Set("Content-Type", mimetypes.JSON)
			if err := json.NewEncoder(w).Encode(resultado); err != nil {
				response.JsonErrorResponse(w, fmt.Errorf("%w: %w", types.InternalError, err))
				return
			}
		},
		"POST /api/contas": func(w http.ResponseWriter, r *http.Request) {

			handler := ApiPostHandler[comandos.AbrirConta, entidade.Conta]{
				Executor: executores.AbrirConta,
			}

			handler.Handle(w, r)
		},
	}
}

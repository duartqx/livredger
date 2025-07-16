package routers

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/duartqx/livredger/internal/api/comandos"
	"github.com/duartqx/livredger/internal/api/decoders"
	"github.com/duartqx/livredger/internal/api/response"

	"github.com/duartqx/livredger/internal/application/services/executores"
	"github.com/duartqx/livredger/internal/application/services/visualizadores"

	"github.com/duartqx/livredger/internal/common/mimetypes"
	"github.com/duartqx/livredger/internal/common/types"

	"github.com/duartqx/livredger/internal/domain/consultas"

	"github.com/duartqx/livredger/internal/infra"
)

func ContasRouter(fs fs.FS) *map[string]http.HandlerFunc {
	return &map[string]http.HandlerFunc{
		"GET /api/contas": func(w http.ResponseWriter, r *http.Request) {
			var usuario *types.Usuario

			uow := infra.Bootstrap(usuario)
			defer uow.Close()

			consulta := consultas.ConsultaContasPadrao()

			if err := r.ParseForm(); err != nil {
				response.JsonErrorResponse(w, fmt.Errorf("%w: %w", decoders.DecoderError, err))
				return
			}

			if err := decoders.Decoder().Decode(consulta, r.Form); err != nil {
				response.JsonErrorResponse(w, fmt.Errorf("%w: %w", decoders.DecoderError, err))
				return
			}

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
		"POST /api/contas": comandos.GenericComandoHandlerFunc(executores.AbrirConta),
	}
}

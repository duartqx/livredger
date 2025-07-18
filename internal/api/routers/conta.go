package routers

import (
	"fmt"
	"io/fs"
	"net/http"

	"github.com/duartqx/livredger/internal/api/command"
	"github.com/duartqx/livredger/internal/api/common"
	"github.com/duartqx/livredger/internal/api/decoders"
	"github.com/duartqx/livredger/internal/api/response"

	"github.com/duartqx/livredger/internal/application/services/executores"
	"github.com/duartqx/livredger/internal/application/services/visualizadores"

	"github.com/duartqx/livredger/internal/common/types"

	"github.com/duartqx/livredger/internal/domain/consultas"

	"github.com/duartqx/livredger/internal/infra"
)

func ContasRouter(fs fs.FS) *common.RouterMap {
	return &common.RouterMap{
		"GET /api/contas": func(w http.ResponseWriter, r *http.Request) {
			var usuario *types.Usuario

			uow, err := infra.Bootstrap(r.Context(), usuario)
			if err != nil {
				response.JsonErrorResponse(w, err)

				return
			}
			defer uow.Close()

			if err := r.ParseForm(); err != nil {
				response.JsonErrorResponse(w, err)

				return
			}

			consulta := consultas.ConsultaContasPadrao()

			if err := decoders.Decoder().Decode(consulta, r.Form); err != nil {

				response.JsonResponse(w, &visualizadores.Resposta[consultas.ConsultaContas]{
					Error: fmt.Errorf("%w: %w", decoders.DecoderError, err),
					Consulta: &visualizadores.Consulta[consultas.ConsultaContas]{
						Parsed: consulta,
						Raw:    r.Form,
					},
				})

				return
			}

			resultado, err := visualizadores.BuscarContas(uow, consulta)

			response.JsonResponse(w, &visualizadores.Resposta[consultas.ConsultaContas]{
				Resultado: resultado,
				Error:     err,
				Consulta: &visualizadores.Consulta[consultas.ConsultaContas]{
					Parsed: consulta,
					Raw:    r.Form,
				},
			})

			return
		},
		"POST /api/contas": command.GenericCommandHandlerFunc(executores.AbrirConta),
	}
}

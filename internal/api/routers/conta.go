package routers

import (
	"cmp"
	"errors"
	"io/fs"
	"net/http"

	"github.com/duartqx/livredger/internal/api/command"
	"github.com/duartqx/livredger/internal/api/common"
	"github.com/duartqx/livredger/internal/api/decoders"
	"github.com/duartqx/livredger/internal/api/response"

	"github.com/duartqx/livredger/internal/application/services/executores"
	"github.com/duartqx/livredger/internal/application/services/visualizadores"

	ce "github.com/duartqx/livredger/internal/common/errors"
	"github.com/duartqx/livredger/internal/common/types"

	"github.com/duartqx/livredger/internal/domain/consultas"

	"github.com/duartqx/livredger/internal/infra"
)

func ContasRouter(fs fs.FS) *common.RouterMap {
	return &common.RouterMap{
		"GET /api/contas": func(w http.ResponseWriter, r *http.Request) {
			var usuario *types.Usuario

			consulta := consultas.ConsultaContasPadrao()

			res := &visualizadores.Response[consultas.ConsultaContas]{
				Query: &visualizadores.Query[consultas.ConsultaContas]{
					Parsed: consulta,
					Raw:    &r.Form,
				},
			}

			uow, err := infra.Bootstrap(r.Context(), usuario)
			if err != nil {
				response.QueryJsonResponse(r.Context(), w, res.WithError(err))

				return
			}
			defer uow.Close()

			if err := cmp.Or(r.ParseForm(), decoders.NewFormDecoder().Decode(consulta, r.Form)); err != nil {
				response.QueryJsonResponse(
					r.Context(), w, res.WithError(errors.Join(ce.RequestError, err)),
				)
				return
			}

			resultado, err := visualizadores.BuscarContas(uow, consulta)

			response.QueryJsonResponse(r.Context(), w, res.WithResult(resultado).WithError(err))

			return
		},
		"POST /api/contas": command.GenericCommandHandlerFunc(executores.AbrirConta),
	}
}

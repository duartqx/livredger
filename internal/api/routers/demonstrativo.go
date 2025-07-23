package routers

import (
	"cmp"
	"errors"
	"io/fs"
	"net/http"

	"github.com/duartqx/livredger/internal/api/decoders"
	"github.com/duartqx/livredger/internal/api/response"
	"github.com/duartqx/livredger/internal/application/services/visualizadores"
	ce "github.com/duartqx/livredger/internal/common/errors"
	"github.com/duartqx/livredger/internal/common/types"
	"github.com/duartqx/livredger/internal/domain/consultas"
	"github.com/duartqx/livredger/internal/domain/entidade"
	"github.com/duartqx/livredger/internal/infra"
)

func DemonstrativosRouter(templFS fs.FS) *RouterMap {
	return &RouterMap{
		"GET /api/demonstrativos/mensal": queryDemonstrativosHandler(
			visualizadores.ConsultarDemonstrativoMensal,
		),
		"GET /api/demonstrativos/ultimos": queryDemonstrativosHandler(
			visualizadores.ConsultarDemonstrativoDosUltimosSeisMeses,
		),
	}
}

func queryDemonstrativosHandler[Q interface {
	consultas.ConsultaDemonstrativoMensal | consultas.ConsultaDemonstrativoUltimosSeisMeses
	Validate() error
}](
	visualizador func(infra.UnidadeDeTrabalho, *Q) (*visualizadores.Result[entidade.DemonstrativoMensal], error),
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var usuario *types.Usuario

		var query Q

		res := &visualizadores.Response[Q]{Query: &visualizadores.Query[Q]{Parsed: &query, Raw: &r.Form}}

		uow, err := infra.Bootstrap(r.Context(), usuario)
		if err != nil {
			response.QueryJsonResponse(r.Context(), w, res.WithError(err))
			return
		}
		defer uow.Close()

		if err := cmp.Or(r.ParseForm(), decoders.NewFormDecoder().Decode(&query, r.Form), query.Validate()); err != nil {
			response.QueryJsonResponse(r.Context(), w, res.WithError(errors.Join(ce.RequestError, err)))
			return
		}

		result, err := visualizador(uow, &query)

		response.QueryJsonResponse(r.Context(), w, res.WithResult(result).WithError(err))

		return
	}
}

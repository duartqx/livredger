package routers

import (
	"cmp"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/duartqx/livredger/internal/api/common"
	"github.com/duartqx/livredger/internal/api/decoders"
	"github.com/duartqx/livredger/internal/api/response"
	"github.com/duartqx/livredger/internal/application/services/visualizadores"
	"github.com/duartqx/livredger/internal/common/types"
	"github.com/duartqx/livredger/internal/domain/consultas"
	"github.com/duartqx/livredger/internal/domain/entidade"
	"github.com/duartqx/livredger/internal/infra"
)

func DemonstrativosRouter(fs fs.FS) *common.RouterMap {
	return &common.RouterMap{
		"GET /api/demonstrativos/mensal": getDemonstrativosHandler(
			visualizadores.ConsultarDemonstrativoMensal,
		),
		"GET /api/demonstrativos/ultimos": getDemonstrativosHandler(
			visualizadores.ConsultarDemonstrativoDosUltimosSeisMeses,
		),
	}
}

func getDemonstrativosHandler[C interface {
	consultas.ConsultaDemonstrativoMensal | consultas.ConsultaDemonstrativoUltimosSeisMeses
	Validar() error
}](
	visualizador func(infra.UnidadeDeTrabalho, *C) (*visualizadores.Resultado[entidade.DemonstrativoMensal], error),
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var usuario *types.Usuario

		var consulta C

		res := &visualizadores.Resposta[C]{Consulta: &visualizadores.Consulta[C]{Parsed: &consulta, Raw: &r.Form}}

		uow, err := infra.Bootstrap(r.Context(), usuario)
		if err != nil {
			response.JsonResponse(r.Context(), w, res.WithError(err))
			return
		}
		defer uow.Close()

		if err := cmp.Or(r.ParseForm(), decoders.Decoder().Decode(&consulta, r.Form), consulta.Validar()); err != nil {
			response.JsonResponse(r.Context(), w, res.WithError(fmt.Errorf("%w: %w", types.RequestError, err)))
			return
		}

		resultado, err := visualizador(uow, &consulta)

		response.JsonResponse(r.Context(), w, res.WithResultado(resultado).WithError(err))

		return
	}
}

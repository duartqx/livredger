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
			visualizadores.ConsultarDemonstrativoDosUltimosTresMeses,
		),
	}
}

type consultaDemonstrativo interface {
	consultas.ConsultaDemonstrativoMensal | consultas.ConsultaDemonstrativoUltimosSeisMeses
	Validar() error
}

func getDemonstrativosHandler[C consultaDemonstrativo](
	visualizador func(infra.UnidadeDeTrabalho, *C) (*visualizadores.Resultado[entidade.DemonstrativoMensal], error),
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		consulta := new(C)

		if err := cmp.Or(decoders.Decoder().Decode(consulta, r.Form), (*consulta).Validar()); err != nil {

			response.JsonResponse(r.Context(), w, &visualizadores.Resposta[C]{
				Error: fmt.Errorf("%w: %w", decoders.DecoderError, err),
				Consulta: &visualizadores.Consulta[C]{
					Parsed: consulta,
					Raw:    r.Form,
				},
			})

			return
		}

		resultado, err := visualizador(uow, consulta)

		response.JsonResponse(r.Context(), w, &visualizadores.Resposta[C]{
			Resultado: resultado,
			Error:     err,
			Consulta: &visualizadores.Consulta[C]{
				Parsed: consulta,
				Raw:    r.Form,
			},
		})

		return
	}
}

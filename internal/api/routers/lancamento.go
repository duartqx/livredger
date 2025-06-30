package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/duartqx/livredger/internal/api/decoders"
	"github.com/duartqx/livredger/internal/api/response"

	"github.com/duartqx/livredger/internal/application/services/executores"
	"github.com/duartqx/livredger/internal/application/services/visualizadores"

	"github.com/duartqx/livredger/internal/common/mimetypes"
	"github.com/duartqx/livredger/internal/common/types"

	"github.com/duartqx/livredger/internal/domain/consultas"

	"github.com/duartqx/livredger/internal/infra"
)

func parseConsultaDoForm(r *http.Request) (*consultas.ConsultaLancamentos, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}

	consulta := consultas.ConsultaLancamentosPadrao()

	if err := decoders.Decoder().Decode(consulta, r.Form); err != nil {
		return nil, fmt.Errorf("%w: %w", decoders.DecoderError, err)
	}

	return consulta, nil
}

func lancamentosRouter() *RouterMap {
	return &RouterMap{
		"GET /api/lancamentos": func(w http.ResponseWriter, r *http.Request) {
			var usuario *types.Usuario

			uow := infra.Bootstrap(usuario)
			defer uow.Close()

			consulta, err := parseConsultaDoForm(r)

			if err != nil {
				response.JsonErrorResponse(w, err)
				return
			}

			resultado, err := visualizadores.BuscarLancamentos(uow, consulta)

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
		"POST /api/lancamentos": ApiPostHandlerFunc(executores.CriarLancamento),
		"GET /lancamentos": View(&ViewContext{
			ViewName: "ConsultarLancamentos",
			Template: parseTemplates(
				"index.html",
				"nav.html",
				"lancamentos/consulta/lancamento.html",
				"lancamentos/consulta/form.html",
				"lancamentos/consulta/consulta.html",
			),
			DataFunc: func(r *http.Request) (map[string]any, error) {
				var usuario *types.Usuario

				uow := infra.Bootstrap(usuario)
				defer uow.Close()

				consulta, err := parseConsultaDoForm(r)

				if err != nil {
					return nil, err
				}

				resultado, err := visualizadores.BuscarLancamentos(uow, consulta)

				if err != nil {
					return nil, err
				}

				return map[string]any{
					"Active":    "Lancamentos",
					"Resultado": resultado,
				}, nil
			},
		}),
		"GET /lancamentos/{chave}": View(&ViewContext{
			ViewName: "DetalhesLancamentos",
			Template: parseTemplates(
				"index.html",
				"nav.html",
				"lancamentos/comando/form.html",
				"lancamentos/detalhes/detalhes.html",
			),
			DataFunc: func(r *http.Request) (map[string]any, error) {

				var usuario *types.Usuario

				uow := infra.Bootstrap(usuario)
				defer uow.Close()

				chave, err := uuid.Parse(r.PathValue("chave"))

				if err != nil {
					return nil, fmt.Errorf("%w: UUID Inválido: %w", types.BusinessLogicError, err)
				}

				resultado, err := visualizadores.BuscarLancamentos(uow, &consultas.ConsultaLancamentos{
					SomenteVersaoMaisRecente: false,
					Chave:                    chave,
					Paginacao: types.Paginacao{
						Pagina:    0,
						Ordenacao: types.Ordenacao{Campo: "timestamp", Direcao: "ASC"},
					},
				})

				if err != nil {
					return nil, err
				}

				return map[string]any{"Active": "Lancamentos", "Resultado": resultado}, nil
			},
		}),
		"GET /lancamentos/criar": View(&ViewContext{
			ViewName: "CriarLancamento",
			Template: parseTemplates(
				"index.html",
				"nav.html",
				"lancamentos/comando/form.html",
				"lancamentos/comando/criar.html",
			),
			Data: map[string]any{"Active": "Lancamentos"},
		}),
	}
}

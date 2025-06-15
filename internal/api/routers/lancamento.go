package routers

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/duartqx/livredger/internal/api/decoders"
	"github.com/duartqx/livredger/internal/application/services/executores"
	"github.com/duartqx/livredger/internal/application/services/visualizadores"
	h "github.com/duartqx/livredger/internal/common/http"
	"github.com/duartqx/livredger/internal/common/types"
	"github.com/duartqx/livredger/internal/domain/comandos"
	"github.com/duartqx/livredger/internal/domain/consultas"
	"github.com/duartqx/livredger/internal/domain/entidade"
	"github.com/duartqx/livredger/internal/infra"
	"github.com/google/uuid"
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
	templateRegistry := ObterTemplateRegistry()

	return &RouterMap{
		"GET /api/lancamentos": func(w http.ResponseWriter, r *http.Request) {
			var usuario *types.Usuario

			uow := infra.Bootstrap(usuario)
			defer uow.Close()

			consulta, err := parseConsultaDoForm(r)

			if err != nil {
				h.JsonErrorResponse(w, err)
				return
			}

			resultado, err := visualizadores.BuscarLancamentos(uow, consulta)

			if err != nil {
				h.JsonErrorResponse(w, err)
				return
			}

			h.HandleResponse(
				&h.Response[consultas.ConsultaLancamentos, entidade.Lancamento]{
					Writer:    w,
					Request:   r,
					Resultado: resultado,
				},
			)
		},
		"POST /api/lancamentos": func(w http.ResponseWriter, r *http.Request) {
			var comando comandos.CriarLancamento

			if err := json.NewDecoder(r.Body).Decode(&comando); err != nil {
				h.JsonErrorResponse(w, fmt.Errorf("%w: %w", types.BusinessLogicError, err))
				return
			}
			defer r.Body.Close()

			var usuario *types.Usuario

			uow := infra.Bootstrap(usuario)
			defer uow.Close()

			resultado, err := executores.CriarLancamento(uow, &comando)

			if err != nil {
				h.JsonErrorResponse(w, err)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)

			if err := json.NewEncoder(w).Encode(map[string]any{"resultado": resultado}); err != nil {
				h.JsonErrorResponse(w, fmt.Errorf("%w: %w", types.InternalError, err))
				return
			}
		},
		"GET /lancamentos": View(&ViewContext{
			ViewName:  "ConsultarLancamentos",
			Templates: templateRegistry.Lancamentos.Consulta,
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
			ViewName:  "DetalhesLancamentos",
			Templates: templateRegistry.Lancamentos.Detalhes,
			DataFunc: func(r *http.Request) (map[string]any, error) {

				var usuario *types.Usuario

				uow := infra.Bootstrap(usuario)
				defer uow.Close()

				chave, err := uuid.Parse(r.PathValue("chave"))

				if err != nil {
					return nil, err
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

				return map[string]any{
					"Active":    "Lancamentos",
					"Resultado": resultado,
				}, nil
			},
		}),
		"GET /lancamentos/criar": View(&ViewContext{
			ViewName:  "CriarLancamento",
			Templates: templateRegistry.Lancamentos.Comando,
			Data:      map[string]any{"Active": "Lancamentos"},
		}),
	}
}

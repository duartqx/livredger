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
)

func criarLancamento(w http.ResponseWriter, r *http.Request) {
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
}

func parseConsultaDoForm(r *http.Request) (*consultas.ConsultaLancamentos, error) {
	if err := r.ParseForm(); err != nil {
		return nil, fmt.Errorf("%w: %w", types.BusinessLogicError, err)
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
		"POST /api/lancamentos": criarLancamento,
	}
}

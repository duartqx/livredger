package routers

import (
	"encoding/json"
	"fmt"

	"net/http"

	e "github.com/duartqx/livredger/internal/application/services/executores"
	v "github.com/duartqx/livredger/internal/application/services/visualizadores"
	h "github.com/duartqx/livredger/internal/common/http"
	t "github.com/duartqx/livredger/internal/common/types"
	c "github.com/duartqx/livredger/internal/domain/comandos"
	"github.com/duartqx/livredger/internal/domain/consultas"
	"github.com/duartqx/livredger/internal/domain/entidade"
	i "github.com/duartqx/livredger/internal/infra"
)

func post(w http.ResponseWriter, r *http.Request) {
	var comando c.CriarLancamento

	if err := json.NewDecoder(r.Body).Decode(&comando); err != nil {
		h.JsonErrorReponse(w, fmt.Errorf("%w: %w", t.BusinessLogicError, err))
		return
	}
	defer r.Body.Close()

	var usuario *t.Usuario

	uow := i.Bootstrap(usuario)
	defer uow.Close()

	resultado, err := e.CriarLancamento(uow, &comando)

	if err != nil {
		h.JsonErrorReponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(map[string]any{"resultado": resultado}); err != nil {
		h.JsonErrorReponse(w, fmt.Errorf("%w: %w", t.InternalError, err))
		return
	}
}

// Query Params
//
//	q: consultas.ConsultaLancamentos
func get(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.JsonErrorReponse(w, fmt.Errorf("%w: %w", t.BusinessLogicError, err))
		return
	}

	var usuario *t.Usuario

	uow := i.Bootstrap(usuario)
	defer uow.Close()

	consulta := consultas.ConsultaLancamentosPadrao()

	if q := r.FormValue("q"); q != "" {
		if err := json.Unmarshal([]byte(q), &consulta); err != nil {
			h.JsonErrorReponse(w, fmt.Errorf("%w: %w", t.BusinessLogicError, err))
			return
		}
	}

	lancamentos, err := v.BuscarLancamentos(uow, consulta)

	if err != nil {
		h.JsonErrorReponse(w, err)
	}

	resultado := h.Resultado[consultas.ConsultaLancamentos, entidade.Lancamento]{
		Total:    len(*lancamentos),
		Consulta: consulta,
		Itens:    lancamentos,
	}

	h.HandleResponse(
		&h.Response[consultas.ConsultaLancamentos, entidade.Lancamento]{
			Writer:    w,
			Request:   r,
			Resultado: &resultado,
			Template:  ObterTemplateRegistry().Lancamentos.Resultados,
		},
	)
}

func lancamentosRouter() *RouterMap {
	return &RouterMap{
		"GET /api/lancamentos":  get,
		"POST /api/lancamentos": post,
	}
}

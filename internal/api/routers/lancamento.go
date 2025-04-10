package routers

import (
	"encoding/json"

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

// Comando:
//
//	{
//		"evento": One of [
//			LancamentoCriado,
//			LancamentoAtualizado,
//			LancamentoPago,
//			LancamentoRecebido,
//			LancamentoCancelado
//		]
//		"chave": uuid
//		"versao": int
//		"vencimento": time.Time
//		"descr": string[500]
//	}
func post(w http.ResponseWriter, r *http.Request) {
	var comando c.CriarLancamento

	if err := json.NewDecoder(r.Body).Decode(&comando); err != nil {
		h.JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var usuario *t.Usuario

	uow := i.Bootstrap(usuario)
	defer uow.Close()

	resultado, err := e.CriarLancamento(uow, &comando)

	if err != nil {
		h.JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(map[string]any{"resultado": resultado}); err != nil {
		h.JsonErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
}

// Params
//
//	chave: uuid
//	somente_versao_mais_recente: true
//	intervalo.inicio: time.Time
//	intervalo.final: time.Time
func get(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	var usuario *t.Usuario

	uow := i.Bootstrap(usuario)
	defer uow.Close()

	consulta, err := consultas.ParsearStringsParaConsultaLancamentos(
		r.FormValue("chave"),
		r.FormValue("somente_versao_mais_recente"),
		r.FormValue("intervalo.inicio"),
		r.FormValue("intervalo.final"),
	)

	if err != nil {
		h.JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	lancamentos, err := v.BuscarLancamentos(uow, consulta)

	if err != nil {
		h.JsonErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resultado := h.Resultado[entidade.Lancamento]{
		Itens: lancamentos,
		Total: len(*lancamentos),
	}

	if err := json.NewEncoder(w).Encode(resultado); err != nil {
		h.JsonErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
}

func LancamentosRouter() *map[string]http.HandlerFunc {
	return &map[string]http.HandlerFunc{
		"GET /api/lancamentos":  get,
		"POST /api/lancamentos": post,
	}
}

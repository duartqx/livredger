package routers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	e "github.com/duartqx/livro-razao/internal/application/services/executores"
	"github.com/duartqx/livro-razao/internal/application/services/visualizadores"
	"github.com/duartqx/livro-razao/internal/common"
	c "github.com/duartqx/livro-razao/internal/domain/comandos"
	i "github.com/duartqx/livro-razao/internal/infra"
)

func writeJsonResponseError(w http.ResponseWriter, err error, statusCode int) {
	res, errMarshal := json.Marshal(map[string]string{"error": err.Error()})

	if errMarshal != nil {
		http.Error(w, errMarshal.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if errors.Is(err, common.NotFoundError) {
		w.WriteHeader(http.StatusNotFound)
		w.Write(res)
		return
	}

	w.WriteHeader(statusCode)

	w.Write(res)
}

func criarLancamento(w http.ResponseWriter, r *http.Request) {
	var comando c.CriarLancamento

	if err := json.NewDecoder(r.Body).Decode(&comando); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var usuario *common.Usuario

	uow := i.Bootstrap(usuario)

	if err := e.CriarLancamento(uow, &comando); err != nil {
		writeJsonResponseError(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func pegarLancamentoPorId(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, fmt.Sprintf("Valor não é um número inteiro válido: %s", idStr), http.StatusBadRequest)
		return
	}

	var usuario *common.Usuario

	uow := i.Bootstrap(usuario)

	lancamento, err := visualizadores.BuscarLancamentoPorId(uow, id)

	if err != nil {
		writeJsonResponseError(w, err, http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(lancamento); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func LancamentosRouter() *map[string]http.HandlerFunc {
	return &map[string]http.HandlerFunc{
		"GET /lancamento/{id}": pegarLancamentoPorId,
		"POST /lancamento":     criarLancamento,
	}
}

package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	h "github.com/duartqx/livredger/internal/common/http"
	"github.com/duartqx/livredger/internal/common/mimetypes"
	"github.com/duartqx/livredger/internal/common/types"
	"github.com/duartqx/livredger/internal/infra"
)

type ApiPostHandler[Comando types.Comando, Entidade any] struct {
	Executor func(infra.UnidadeDeTrabalho, *Comando) (*Entidade, error)
}

func (ph ApiPostHandler[Comando, Entidade]) Handle(w http.ResponseWriter, r *http.Request) {
	var comando Comando

	if err := json.NewDecoder(r.Body).Decode(&comando); err != nil {
		h.JsonErrorResponse(w, fmt.Errorf("%w: %w", types.BusinessLogicError, err))
		return
	}
	defer r.Body.Close()

	if err := comando.Validar(); err != nil {
		h.JsonErrorResponse(w, err)
		return
	}

	var usuario *types.Usuario

	uow := infra.Bootstrap(usuario)
	defer uow.Close()

	resultado, err := ph.Executor(uow, &comando)

	if err != nil {
		h.JsonErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", mimetypes.JSON)
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(map[string]any{"resultado": resultado}); err != nil {
		h.JsonErrorResponse(w, fmt.Errorf("%w: %w", types.InternalError, err))
		return
	}
}

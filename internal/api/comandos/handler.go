package comandos

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/duartqx/livredger/internal/api/response"

	"github.com/duartqx/livredger/internal/domain"

	"github.com/duartqx/livredger/internal/common/mimetypes"
	"github.com/duartqx/livredger/internal/common/types"

	"github.com/duartqx/livredger/internal/infra"
)

func GenericComandoHandlerFunc[Comando domain.Comando, Entidade any](
	executor func(infra.UnidadeDeTrabalho, *Comando) (*Entidade, error),
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var comando Comando

		if err := json.NewDecoder(r.Body).Decode(&comando); err != nil {
			response.JsonErrorResponse(w, fmt.Errorf("%w: %w", types.BusinessLogicError, err))
			return
		}
		defer r.Body.Close()

		if err := comando.Validar(); err != nil {
			response.JsonErrorResponse(w, err)
			return
		}

		var usuario *types.Usuario

		uow := infra.Bootstrap(usuario)
		defer uow.Close()

		resultado, err := executor(uow, &comando)

		if err != nil {
			response.JsonErrorResponse(w, err)
			return
		}

		w.Header().Set("Content-Type", mimetypes.JSON)
		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(map[string]any{"resultado": resultado}); err != nil {
			response.JsonErrorResponse(w, fmt.Errorf("%w: %w", types.InternalError, err))
			return
		}
	}
}

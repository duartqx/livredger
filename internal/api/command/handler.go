package command

import (
	"cmp"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/duartqx/livredger/internal/api/response"
	"github.com/duartqx/livredger/internal/application/services/executores"
	"github.com/duartqx/livredger/internal/domain/entidade"

	"github.com/duartqx/livredger/internal/domain"

	"github.com/duartqx/livredger/internal/common/types"

	"github.com/duartqx/livredger/internal/infra"
)

func ComandoFromJson[Command domain.Command](body io.ReadCloser, comando *Command) error {
	if err := json.NewDecoder(body).Decode(comando); err != nil {
		return fmt.Errorf("%w: %w", types.RequestError, err)
	}
	return nil
}

func GenericCommandHandlerFunc[Command domain.Command, Entidade entidade.Entidade](
	executor func(infra.UnidadeDeTrabalho, *Command) (*Entidade, error),
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		res := &executores.Response[Command]{Comando: new(Command)}

		if err := cmp.Or(ComandoFromJson(r.Body, res.Comando), (*res.Comando).Validar()); err != nil {
			response.CommandJsonResponse(w, res.WithError(err))
			return
		}
		defer r.Body.Close()

		var usuario *types.Usuario

		uow, err := infra.Bootstrap(r.Context(), usuario)
		if err != nil {
			response.CommandJsonResponse(w, res.WithError(err))

			return
		}
		defer uow.Close()

		resultado, err := executor(uow, res.Comando)

		if err != nil {
			response.CommandJsonResponse(w, res.WithError(err))
			return
		}

		w.WriteHeader(http.StatusCreated)

		response.CommandJsonResponse(w, res.WithResultado(resultado).WithError(nil))

		return
	}
}

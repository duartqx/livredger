package routers

import (
	"cmp"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/duartqx/livredger/internal/api/decoders"
	"github.com/duartqx/livredger/internal/api/response"
	"github.com/duartqx/livredger/internal/application/services/executores"
	ce "github.com/duartqx/livredger/internal/common/errors"
	"github.com/duartqx/livredger/internal/common/types"
	"github.com/duartqx/livredger/internal/domain"
	"github.com/duartqx/livredger/internal/domain/entidade"
	"github.com/duartqx/livredger/internal/infra"
)

type DataFunc func(r *http.Request) (map[string]any, error)

type ViewContext struct {
	ViewName string
	Template *template.Template
	Data     map[string]any
	DataFunc DataFunc
}

func View(ctx *ViewContext) http.HandlerFunc {
	if ctx.Data != nil && ctx.DataFunc != nil {
		panic(fmt.Errorf("Você deve passar somente ViewContext.Data ou ViewContext.DataFunc, mas nunca ambos"))
	}

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("View", ctx.ViewName)

		auth := r.Header.Get("Authorization")

		if auth != "" {
			log.Println("Authorization", auth)
		}

		var (
			data map[string]any = ctx.Data
			err  error
		)

		if ctx.DataFunc != nil {
			data, err = ctx.DataFunc(r)
		}

		switch {
		case err == nil && r.Header.Get("HX-Request") == "true":
			if err := ctx.Template.ExecuteTemplate(w, "partial", data); err != nil {
				panic(err)
			}
		case err == nil:
			if err := ctx.Template.ExecuteTemplate(w, "base", data); err != nil {
				panic(err)
			}
		case errors.Is(err, decoders.FormDecoderError) && ctx.Template.Lookup("error") != nil:

			w.WriteHeader(400)

			data = map[string]any{"Errors": decoders.ParseFormDecodeError(err)}

			if err := ctx.Template.ExecuteTemplate(w, "error", data); err != nil {
				panic(err)
			}
		default:
			response.ErrorResponse(w, err)
		}
	}
}

func ComandoFromJson[Command domain.Command](body io.ReadCloser, comando *Command) error {
	if err := json.NewDecoder(body).Decode(comando); err != nil {
		return errors.Join(ce.RequestError, err)
	}
	return nil
}

func GenericCommandHandlerFunc[Command domain.Command, Entidade entidade.Entidade](
	executor func(infra.UnidadeDeTrabalho, *Command) (*Entidade, error),
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		res := &executores.Response[Command]{Command: new(Command)}

		if err := cmp.Or(ComandoFromJson(r.Body, res.Command), (*res.Command).Validate()); err != nil {
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

		resultado, err := executor(uow, res.Command)

		if err != nil {
			response.CommandJsonResponse(w, res.WithError(err))
			return
		}

		w.WriteHeader(http.StatusCreated)

		response.CommandJsonResponse(w, res.WithResult(resultado).WithError(nil))

		return
	}
}

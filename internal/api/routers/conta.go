package routers

import (
	"cmp"
	"errors"
	"io/fs"
	"net/http"

	"github.com/duartqx/livredger/internal/api/decoders"
	"github.com/duartqx/livredger/internal/api/response"
	"github.com/duartqx/livredger/internal/application"

	"github.com/duartqx/livredger/internal/application/services/executores"
	"github.com/duartqx/livredger/internal/application/services/visualizadores"

	ce "github.com/duartqx/livredger/internal/common/errors"
	"github.com/duartqx/livredger/internal/common/types"

	"github.com/duartqx/livredger/internal/domain/consultas"
)

func ContasRouter(templFS fs.FS) *RouterMap {
	return &RouterMap{
		"GET /api/contas":  GetApiContas,
		"POST /api/contas": GenericCommandHandlerFunc(executores.AbrirConta),
	}
}

func GetApiContas(w http.ResponseWriter, r *http.Request) {
	var usuario *types.Usuario

	consulta := consultas.ConsultaContasPadrao()

	res := &visualizadores.Response[consultas.ConsultaContas]{
		Query: &visualizadores.Query[consultas.ConsultaContas]{
			Parsed: consulta,
			Raw:    &r.Form,
		},
	}

	uow, err := application.NovaUnidadeDeTrabalho(r.Context(), usuario)
	if err != nil {
		response.QueryJsonResponse(r.Context(), w, res.WithError(err))

		return
	}
	defer uow.Close()

	if err := cmp.Or(r.ParseForm(), decoders.NewFormDecoder().Decode(consulta, r.Form)); err != nil {
		response.QueryJsonResponse(
			uow.Context(), w, res.WithError(errors.Join(ce.RequestError, err)),
		)
		return
	}

	resultado, err := visualizadores.BuscarContas(uow, consulta)

	response.QueryJsonResponse(uow.Context(), w, res.WithResult(resultado).WithError(err))

	return
}

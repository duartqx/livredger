package routers

import (
	"cmp"
	"errors"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/google/uuid"

	"github.com/duartqx/livredger/internal/api/command"
	"github.com/duartqx/livredger/internal/api/common"
	"github.com/duartqx/livredger/internal/api/decoders"
	"github.com/duartqx/livredger/internal/api/response"
	"github.com/duartqx/livredger/internal/api/templates"

	"github.com/duartqx/livredger/internal/application/services/executores"
	"github.com/duartqx/livredger/internal/application/services/visualizadores"

	ce "github.com/duartqx/livredger/internal/common/errors"
	"github.com/duartqx/livredger/internal/common/types"

	"github.com/duartqx/livredger/internal/domain/consultas"

	"github.com/duartqx/livredger/internal/infra"
)

func LancamentosRouter(fs fs.FS) *common.RouterMap {
	return &common.RouterMap{
		"GET /api/lancamentos": func(w http.ResponseWriter, r *http.Request) {
			var usuario *types.Usuario

			consulta := consultas.ConsultaLancamentosPadrao()

			res := &visualizadores.Response[consultas.ConsultaLancamentos]{
				Query: &visualizadores.Query[consultas.ConsultaLancamentos]{
					Parsed: consulta,
					Raw:    &r.Form,
				},
			}

			uow, err := infra.Bootstrap(r.Context(), usuario)
			if err != nil {
				response.QueryJsonResponse(r.Context(), w, res.WithError(err))
				return
			}
			defer uow.Close()

			if err := cmp.Or(r.ParseForm(), decoders.NewFormDecoder().Decode(consulta, r.Form)); err != nil {
				response.QueryJsonResponse(r.Context(), w, res.WithError(errors.Join(ce.RequestError, err)))
				return
			}

			resultado, err := visualizadores.BuscarLancamentos(uow, consulta)

			response.QueryJsonResponse(r.Context(), w, res.WithResult(resultado).WithError(err))

			return
		},
		"POST /api/lancamentos": command.GenericCommandHandlerFunc(executores.CriarLancamento),
		"GET /lancamentos": View(&ViewContext{
			ViewName: "ConsultarLancamentos",
			Template: templates.Templates(
				fs,
				"index.html",
				"nav.html",
				"lancamentos/consulta/lancamento.html",
				"lancamentos/consulta/form.html",
				"lancamentos/consulta/consulta.html",
			),
			DataFunc: func(r *http.Request) (map[string]any, error) {
				var usuario *types.Usuario

				uow, err := infra.Bootstrap(r.Context(), usuario)
				if err != nil {
					return nil, err
				}
				defer uow.Close()

				consulta := consultas.ConsultaLancamentosPadrao()

				if err := cmp.Or(r.ParseForm(), decoders.NewFormDecoder().Decode(consulta, r.Form)); err != nil {
					return nil, errors.Join(ce.RequestError, err)
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
			ViewName: "DetalhesLancamentos",
			Template: templates.Templates(
				fs,
				"index.html",
				"nav.html",
				"lancamentos/comando/form.html",
				"lancamentos/detalhes/detalhes.html",
			),
			DataFunc: func(r *http.Request) (map[string]any, error) {

				var usuario *types.Usuario

				uow, err := infra.Bootstrap(r.Context(), usuario)
				if err != nil {
					return nil, err
				}
				defer uow.Close()

				chave, err := uuid.Parse(r.PathValue("chave"))

				if err != nil {
					return nil, fmt.Errorf("%w: UUID Inválido: %w", ce.RequestError, err)
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

				return map[string]any{"Active": "Lancamentos", "Resultado": resultado}, nil
			},
		}),
		"GET /lancamentos/criar": View(&ViewContext{
			ViewName: "CriarLancamento",
			Template: templates.Templates(
				fs,
				"index.html",
				"nav.html",
				"lancamentos/comando/form.html",
				"lancamentos/comando/criar.html",
			),
			Data: map[string]any{"Active": "Lancamentos"},
		}),
	}
}

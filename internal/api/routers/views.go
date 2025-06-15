package routers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/duartqx/livredger/internal/api/decoders"
	"github.com/duartqx/livredger/internal/application/services/visualizadores"
	h "github.com/duartqx/livredger/internal/common/http"
	"github.com/duartqx/livredger/internal/common/types"
	"github.com/duartqx/livredger/internal/domain/consultas"
	"github.com/duartqx/livredger/internal/infra"
	"github.com/google/uuid"
)

type DataFunc func(r *http.Request) (map[string]any, error)

type ViewContext struct {
	ViewName  string
	Templates *h.Templates
	Data      map[string]any
	DataFunc  DataFunc
}

func view(ctx *ViewContext) http.HandlerFunc {
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
		case err == nil && ctx.Templates.Partial != nil && r.Header.Get("HX-Request") == "true":
			if err := ctx.Templates.Partial.ExecuteTemplate(w, "partial", data); err != nil {
				panic(err)
			}
		case err == nil:
			if err := ctx.Templates.ComBase.ExecuteTemplate(w, "base", data); err != nil {
				panic(err)
			}
		case errors.Is(err, decoders.DecoderError) && ctx.Templates.Error != nil:
			w.WriteHeader(400)
			if err := ctx.Templates.Error.ExecuteTemplate(
				w, "error", map[string]any{"Errors": decoders.ParseDecodeError(err)},
			); err != nil {
				panic(err)
			}
		default:
			h.ErrorResponse(w, err)
		}
	}

}

func viewsRouter() *RouterMap {
	registry := ObterTemplateRegistry()
	return &RouterMap{
		"GET /{$}": view(&ViewContext{
			ViewName:  "Index",
			Templates: registry.Index,
			Data:      map[string]any{"Active": "Index"},
		}),
		"GET /lancamentos": view(&ViewContext{
			ViewName:  "ConsultarLancamentos",
			Templates: registry.Lancamentos.Consulta,
			DataFunc: func(r *http.Request) (map[string]any, error) {
				var usuario *types.Usuario

				uow := infra.Bootstrap(usuario)
				defer uow.Close()

				consulta, err := parseConsultaDoForm(r)

				if err != nil {
					return nil, err
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
		"GET /lancamentos/{chave}": view(&ViewContext{
			ViewName:  "DetalhesLancamentos",
			Templates: registry.Lancamentos.Detalhes,
			DataFunc: func(r *http.Request) (map[string]any, error) {

				var usuario *types.Usuario

				uow := infra.Bootstrap(usuario)
				defer uow.Close()

				chave, err := uuid.Parse(r.PathValue("chave"))

				if err != nil {
					return nil, err
				}

				resultado, err := visualizadores.BuscarLancamentos(uow, &consultas.ConsultaLancamentos{
					SomenteVersaoMaisRecente: false,
					Chave:                    chave,
					Paginacao:                types.Paginacao{Pagina: 0, Ordenacao: types.Ordenacao{Campo: "timestamp", Direcao: "ASC"}},
				})

				if err != nil {
					return nil, err
				}

				return map[string]any{
					"Active":    "Lancamentos",
					"Resultado": resultado,
				}, nil
			},
		}),
		"GET /lancamentos/criar": view(&ViewContext{
			ViewName:  "CriarLancamento",
			Templates: registry.Lancamentos.Comando,
			Data:      map[string]any{"Active": "Lancamentos"},
		}),
	}
}

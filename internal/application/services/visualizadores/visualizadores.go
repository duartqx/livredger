package visualizadores

import (
	"github.com/duartqx/livredger/internal/common/types"
	"github.com/duartqx/livredger/internal/domain/consultas"
	"github.com/duartqx/livredger/internal/domain/entidade"
	"github.com/duartqx/livredger/internal/infra"
)

func BuscarLancamentos(uow *infra.UnidadeDeTrabalho, consulta *consultas.ConsultaLancamentos) (
	*types.Resultado[consultas.ConsultaLancamentos, entidade.Lancamento], error,
) {

	lancamentos, err := uow.Repositorios.Lancamentos.Consulta.Buscar(uow.DB, consulta)

	if err != nil {
		return nil, err
	}

	resultado := &types.Resultado[consultas.ConsultaLancamentos, entidade.Lancamento]{
		Total:    len(*lancamentos),
		Consulta: consulta,
		Itens:    lancamentos,
	}

	return resultado, err
}

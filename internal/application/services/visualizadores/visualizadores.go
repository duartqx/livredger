package visualizadores

import (
	"github.com/duartqx/livredger/internal/application"

	"github.com/duartqx/livredger/internal/domain/consultas"
	"github.com/duartqx/livredger/internal/domain/entidade"

	"github.com/duartqx/livredger/internal/infra"
)

func BuscarLancamentos(uow infra.UnidadeDeTrabalho, consulta *consultas.ConsultaLancamentos) (
	*application.Resultado[consultas.ConsultaLancamentos, entidade.Lancamento], error,
) {

	lancamentos, err := uow.GetRepositorios().Lancamentos().Consulta.Buscar(uow.GetDB(), consulta)

	if err != nil {
		return nil, err
	}

	resultado := &application.Resultado[consultas.ConsultaLancamentos, entidade.Lancamento]{
		Total:    len(*lancamentos),
		Consulta: consulta,
		Itens:    lancamentos,
	}

	return resultado, err
}

func BuscarContas(uow infra.UnidadeDeTrabalho, consulta *consultas.ConsultaContas) (
	*application.Resultado[consultas.ConsultaContas, entidade.Conta], error,
) {

	contas, err := uow.GetRepositorios().Contas().Consulta.Buscar(uow.GetDB(), consulta)

	if err != nil {
		return nil, err
	}

	resultado := &application.Resultado[consultas.ConsultaContas, entidade.Conta]{
		Total:    len(*contas),
		Consulta: consulta,
		Itens:    contas,
	}

	return resultado, err
}

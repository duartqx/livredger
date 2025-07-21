package visualizadores

import (
	"cmp"

	"github.com/duartqx/livredger/internal/domain/consultas"
	"github.com/duartqx/livredger/internal/domain/entidade"

	"github.com/duartqx/livredger/internal/infra"
)

func BuscarLancamentos(uow infra.UnidadeDeTrabalho, consulta *consultas.ConsultaLancamentos) (
	*Result[entidade.Lancamento], error,
) {
	lancamentos, err := uow.GetRepositorios().Lancamentos.Consulta.Buscar(uow.GetContext(), uow.GetDB(), consulta)

	lancamentos = cmp.Or(lancamentos, &[]*entidade.Lancamento{})

	resultado := &Result[entidade.Lancamento]{
		Total: len(*lancamentos),
		Itens: lancamentos,
	}

	return resultado, err
}

func BuscarContas(uow infra.UnidadeDeTrabalho, consulta *consultas.ConsultaContas) (
	*Result[entidade.Conta], error,
) {
	contas, err := uow.GetRepositorios().Contas.Consulta.Buscar(uow.GetContext(), uow.GetDB(), consulta)

	contas = cmp.Or(contas, &[]*entidade.Conta{})

	resultado := &Result[entidade.Conta]{
		Total: len(*contas),
		Itens: contas,
	}

	return resultado, err
}

func ConsultarDemonstrativoMensal(uow infra.UnidadeDeTrabalho, consulta *consultas.ConsultaDemonstrativoMensal) (
	*Result[entidade.DemonstrativoMensal], error,
) {
	demonstrativoSlice := make([]*entidade.DemonstrativoMensal, 0, 1)

	demonstrativo, err := uow.GetRepositorios().Demonstrativos.Consulta.DemonstrativoMensal(
		uow.GetContext(), uow.GetDB(), consulta,
	)

	if demonstrativo != nil {
		demonstrativoSlice = append(demonstrativoSlice[:0], demonstrativo)
	}

	resultado := &Result[entidade.DemonstrativoMensal]{
		Total: len(demonstrativoSlice),
		Itens: &demonstrativoSlice,
	}

	return resultado, err
}

func ConsultarDemonstrativoDosUltimosSeisMeses(uow infra.UnidadeDeTrabalho, consulta *consultas.ConsultaDemonstrativoUltimosSeisMeses) (
	*Result[entidade.DemonstrativoMensal], error,
) {
	demonstrativos, err := uow.GetRepositorios().Demonstrativos.Consulta.DemonstrativosDosUltimosSeisMeses(
		uow.GetContext(), uow.GetDB(), consulta,
	)

	demonstrativos = cmp.Or(demonstrativos, &[]*entidade.DemonstrativoMensal{})

	resultado := &Result[entidade.DemonstrativoMensal]{
		Total: len(*demonstrativos),
		Itens: demonstrativos,
	}

	return resultado, err
}

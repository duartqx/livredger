package visualizadores

import (
	"cmp"

	"github.com/duartqx/livredger/internal/application"
	"github.com/duartqx/livredger/internal/domain/consultas"
	"github.com/duartqx/livredger/internal/domain/entidade"
)

func BuscarLancamentos(uow application.UnidadeDeTrabalho, consulta *consultas.ConsultaLancamentos) (
	*Result[entidade.Lancamento], error,
) {
	lancamentos, err := uow.Repositorios().Lancamentos.Consulta.Buscar(uow.Context(), uow.DB(), consulta)

	lancamentos = cmp.Or(lancamentos, &[]*entidade.Lancamento{})

	resultado := &Result[entidade.Lancamento]{
		Total: len(*lancamentos),
		Itens: lancamentos,
	}

	return resultado, err
}

func BuscarContas(uow application.UnidadeDeTrabalho, consulta *consultas.ConsultaContas) (
	*Result[entidade.Conta], error,
) {
	contas, err := uow.Repositorios().Contas.Consulta.Buscar(uow.Context(), uow.DB(), consulta)

	contas = cmp.Or(contas, &[]*entidade.Conta{})

	resultado := &Result[entidade.Conta]{
		Total: len(*contas),
		Itens: contas,
	}

	return resultado, err
}

func ConsultarDemonstrativoMensal(
	uow application.UnidadeDeTrabalho, consulta *consultas.ConsultaDemonstrativoMensal,
) (*Result[entidade.DemonstrativoMensal], error) {
	demonstrativoSlice := make([]*entidade.DemonstrativoMensal, 0, 1)

	demonstrativo, err := uow.Repositorios().Demonstrativos.Consulta.DemonstrativoMensal(
		uow.Context(), uow.DB(), consulta,
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

func ConsultarDemonstrativoDosUltimosSeisMeses(
	uow application.UnidadeDeTrabalho, consulta *consultas.ConsultaDemonstrativoUltimosSeisMeses,
) (*Result[entidade.DemonstrativoMensal], error) {
	demonstrativos, err := uow.Repositorios().Demonstrativos.Consulta.DemonstrativosDosUltimosSeisMeses(
		uow.Context(), uow.DB(), consulta,
	)

	demonstrativos = cmp.Or(demonstrativos, &[]*entidade.DemonstrativoMensal{})

	resultado := &Result[entidade.DemonstrativoMensal]{
		Total: len(*demonstrativos),
		Itens: demonstrativos,
	}

	return resultado, err
}

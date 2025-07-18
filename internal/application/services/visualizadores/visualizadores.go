package visualizadores

import (
	"cmp"

	"github.com/duartqx/livredger/internal/application"

	"github.com/duartqx/livredger/internal/domain/consultas"
	"github.com/duartqx/livredger/internal/domain/entidade"

	"github.com/duartqx/livredger/internal/infra"
)

func BuscarLancamentos(uow infra.UnidadeDeTrabalho, consulta *consultas.ConsultaLancamentos) (
	*application.Resultado[entidade.Lancamento], error,
) {
	lancamentos, err := uow.GetRepositorios().Lancamentos().Consulta.Buscar(uow.GetDB(), consulta)

	lancamentos = cmp.Or(lancamentos, &[]*entidade.Lancamento{})

	resultado := &application.Resultado[entidade.Lancamento]{
		Total: len(*lancamentos),
		Itens: lancamentos,
	}

	return resultado, err
}

func BuscarContas(uow infra.UnidadeDeTrabalho, consulta *consultas.ConsultaContas) (
	*application.Resultado[entidade.Conta], error,
) {
	contas, err := uow.GetRepositorios().Contas().Consulta.Buscar(uow.GetDB(), consulta)

	contas = cmp.Or(contas, &[]*entidade.Conta{})

	resultado := &application.Resultado[entidade.Conta]{
		Total: len(*contas),
		Itens: contas,
	}

	return resultado, err
}

func ConsultarDemonstrativoMensal(uow infra.UnidadeDeTrabalho, consulta *consultas.ConsultaDemonstrativoMensal) (
	*application.Resultado[entidade.DemonstrativoMensal], error,
) {

	var demonstrativoSlice []*entidade.DemonstrativoMensal

	demonstrativo, err := uow.GetRepositorios().Demonstrativos().Consulta.DemonstrativoMensal(uow.GetDB(), consulta)

	if demonstrativo != nil {
		demonstrativoSlice = append(demonstrativoSlice, demonstrativo)
	}

	resultado := &application.Resultado[entidade.DemonstrativoMensal]{
		Total: 0,
		Itens: &demonstrativoSlice,
	}

	return resultado, err
}

func ConsultarDemonstrativoDosUltimosTresMeses(uow infra.UnidadeDeTrabalho, consulta *consultas.ConsultaDemonstrativoMensal) (
	*application.Resultado[entidade.DemonstrativoMensal], error,
) {
	demonstrativos, err := uow.GetRepositorios().Demonstrativos().Consulta.DemonstrativosDosUltimosTresMeses(uow.GetDB(), consulta)

	demonstrativos = cmp.Or(demonstrativos, &[]*entidade.DemonstrativoMensal{})

	resultado := &application.Resultado[entidade.DemonstrativoMensal]{
		Total: len(*demonstrativos),
		Itens: demonstrativos,
	}

	return resultado, err
}

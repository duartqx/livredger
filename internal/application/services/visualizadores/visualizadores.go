package visualizadores

import (
	c "github.com/duartqx/livredger/internal/domain/consultas"
	e "github.com/duartqx/livredger/internal/domain/entidade"
	i "github.com/duartqx/livredger/internal/infra"
)

func BuscarLancamentos(uow *i.UnidadeDeTrabalho, consulta *c.ConsultaLancamentos) (*[]*e.Lancamento, error) {
	return uow.Repositorios.Lancamentos.Consulta.Buscar(uow.DB, consulta)
}

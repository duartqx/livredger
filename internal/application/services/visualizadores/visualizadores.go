package visualizadores

import (
	e "github.com/duartqx/livredger/internal/domain/entidade"
	i "github.com/duartqx/livredger/internal/infra"
	r "github.com/duartqx/livredger/internal/infra/repositorios/sqlite/consultas"
)

func BuscarLancamentoPorId(uow *i.UnidadeDeTrabalho, id int) (*e.Lancamento, error) {

	repo := r.NewRepositorioDeConsultaLancamentos()

	return repo.BuscarPorId(uow.DB, id)
}

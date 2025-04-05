package executores

import (
	c "github.com/duartqx/livredger/internal/domain/comandos"
	i "github.com/duartqx/livredger/internal/infra"
	r "github.com/duartqx/livredger/internal/infra/repositorios/sqlite/comandos"
)

func CriarLancamento(uow *i.UnidadeDeTrabalho, comando *c.CriarLancamento) error {

	if err := comando.Validar(); err != nil {
		return err
	}

	tx, err := uow.Transaction()

	if err != nil {
		return err
	}

	repo := r.NewRepositorioDeComandoLancamentos()

	if err := repo.Criar(tx, comando); err != nil {
		return uow.Rollback()
	}

	return uow.Commit()
}

package executores

import (
	c "github.com/duartqx/livredger/internal/domain/comandos"
	e "github.com/duartqx/livredger/internal/domain/entidade"
	i "github.com/duartqx/livredger/internal/infra"
)

func CriarLancamento(uow *i.UnidadeDeTrabalho, comando *c.CriarLancamento) (*e.Lancamento, error) {

	if err := comando.Validar(); err != nil {
		return nil, err
	}

	tx, err := uow.Transaction()

	if err != nil {
		return nil, err
	}

	lancamento, err := uow.Repositorios.Lancamento.Comando.Criar(tx, comando)

	if err != nil {
		uow.Rollback()
		return nil, err
	}

	if err := uow.Commit(); err != nil {
		uow.Rollback()
		return nil, err
	}

	return lancamento, nil
}

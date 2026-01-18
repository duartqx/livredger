package executores

import (
	"database/sql"

	"github.com/duartqx/livredger/internal/application"
	"github.com/duartqx/livredger/internal/domain/entidade"
)

func TransactionalScript[Entidade entidade.Entidade](
	uow application.UnidadeDeTrabalho, executor func(*sql.Tx) (*Entidade, error),
) (*Entidade, error) {
	tx, err := uow.BeginTransaction()

	if err != nil {
		return nil, err
	}

	defer uow.Rollback()

	resultado, err := executor(tx)

	if err != nil {
		return nil, err
	}

	if err := uow.Commit(); err != nil {
		return nil, err
	}

	return resultado, nil
}

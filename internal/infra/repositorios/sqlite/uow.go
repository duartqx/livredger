package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	ce "github.com/duartqx/livredger/internal/common/errors"
	"github.com/duartqx/livredger/internal/common/types"
	"github.com/duartqx/livredger/internal/domain/repositorios"
	"github.com/duartqx/livredger/internal/infra/repositorios/sqlite/regex"
)

type UnidadeDeTrabalhoSqlite struct {
	Context      context.Context
	Usuario      *types.Usuario
	DB           *sql.DB
	Tx           *sql.Tx
	Repositorios *repositorios.Repositorios
}

func (u UnidadeDeTrabalhoSqlite) GetContext() context.Context {
	return u.Context
}

func (u UnidadeDeTrabalhoSqlite) GetUsuario() *types.Usuario {
	return u.Usuario
}

func (u UnidadeDeTrabalhoSqlite) GetRepositorios() *repositorios.Repositorios {
	return u.Repositorios
}

func (u UnidadeDeTrabalhoSqlite) GetDB() *sql.DB {
	return u.DB
}

func (u UnidadeDeTrabalhoSqlite) GetTransaction() *sql.Tx {
	if u.Tx == nil {
		panic(fmt.Errorf("%w UnidadeDeTrabalho: Já existe uma transação aberta", ce.InternalError))
	}
	return u.Tx
}

func (u *UnidadeDeTrabalhoSqlite) BeginTransaction() (tx *sql.Tx, err error) {
	if u.Tx != nil {
		return nil, fmt.Errorf("%w UnidadeDeTrabalho: Já existe uma transação aberta", ce.InternalError)
	}

	tx, err = u.DB.Begin()

	if err != nil {
		return nil, fmt.Errorf("%w UnidadeDeTrabalho: Não foi possível iniciar uma transação (%w)", ce.InternalError, err)
	}

	u.Tx = tx

	return tx, nil
}

func (u *UnidadeDeTrabalhoSqlite) Commit() error {
	if u.Tx == nil {
		return fmt.Errorf("%w UnidadeDeTrabalho: Nenhuma transação aberta", ce.InternalError)
	}

	if err := u.Tx.Commit(); err != nil {

		if match := regex.SqliteFalhouCommitar.FindStringSubmatch(err.Error()); len(match) > 1 {
			return fmt.Errorf("%w: %s", ce.BusinessLogicError, match[1])
		}

		return fmt.Errorf("%w UnidadeDeTrabalho: Não foi possível commitar a transação (%w)", ce.InternalError, err)
	}

	return nil
}

func (u *UnidadeDeTrabalhoSqlite) Rollback() error {
	if u.Tx == nil {
		return fmt.Errorf("%w UnidadeDeTrabalho: Nenhuma transação aberta", ce.InternalError)
	}

	if err := u.Tx.Rollback(); err != nil {
		return fmt.Errorf("%w UnidadeDeTrabalho: Não foi possível fazer rollback (%w)", ce.InternalError, err)
	}

	return nil
}

func (u *UnidadeDeTrabalhoSqlite) Close() {
	if u.DB != nil {
		u.DB.Close()
	}
}

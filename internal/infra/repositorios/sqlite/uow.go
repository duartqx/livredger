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
	context      context.Context
	usuario      *types.Usuario
	db           *sql.DB
	tx           *sql.Tx
	repositorios *repositorios.Repositorios
}

func NewUnidadeDeTrabalhoSqlite(
	context context.Context,
	usuario *types.Usuario,
	db *sql.DB,
	repos *repositorios.Repositorios,
) *UnidadeDeTrabalhoSqlite {
	return &UnidadeDeTrabalhoSqlite{
		context:      context,
		usuario:      usuario,
		db:           db,
		repositorios: repos,
	}
}

func (u UnidadeDeTrabalhoSqlite) Context() context.Context {
	return u.context
}

func (u UnidadeDeTrabalhoSqlite) Usuario() *types.Usuario {
	return u.usuario
}

func (u UnidadeDeTrabalhoSqlite) Repositorios() *repositorios.Repositorios {
	return u.repositorios
}

func (u UnidadeDeTrabalhoSqlite) DB() *sql.DB {
	return u.db
}

func (u UnidadeDeTrabalhoSqlite) Transaction() *sql.Tx {
	if u.tx == nil {
		panic(fmt.Errorf("%w UnidadeDeTrabalho: Já existe uma transação aberta", ce.InternalError))
	}
	return u.tx
}

func (u *UnidadeDeTrabalhoSqlite) BeginTransaction() (tx *sql.Tx, err error) {
	if u.tx != nil {
		return nil, fmt.Errorf("%w UnidadeDeTrabalho: Já existe uma transação aberta", ce.InternalError)
	}

	tx, err = u.db.Begin()

	if err != nil {
		return nil, fmt.Errorf("%w UnidadeDeTrabalho: Não foi possível iniciar uma transação (%w)", ce.InternalError, err)
	}

	u.tx = tx

	return tx, nil
}

func (u *UnidadeDeTrabalhoSqlite) Commit() error {
	if u.tx == nil {
		return fmt.Errorf("%w UnidadeDeTrabalho: Nenhuma transação aberta", ce.InternalError)
	}

	if err := u.tx.Commit(); err != nil {

		if match := regex.SqliteFalhouCommitar.FindStringSubmatch(err.Error()); len(match) > 1 {
			return fmt.Errorf("%w: %s", ce.BusinessLogicError, match[1])
		}

		return fmt.Errorf("%w UnidadeDeTrabalho: Não foi possível commitar a transação (%w)", ce.InternalError, err)
	}

	return nil
}

func (u *UnidadeDeTrabalhoSqlite) Rollback() error {
	if u.tx == nil {
		return fmt.Errorf("%w UnidadeDeTrabalho: Nenhuma transação aberta", ce.InternalError)
	}

	if err := u.tx.Rollback(); err != nil {
		return fmt.Errorf("%w UnidadeDeTrabalho: Não foi possível fazer rollback (%w)", ce.InternalError, err)
	}

	return nil
}

func (u *UnidadeDeTrabalhoSqlite) Close() {
	if u.db != nil {
		u.db.Close()
	}
}

package infra

import (
	"database/sql"
	"fmt"

	t "github.com/duartqx/livredger/internal/common/types"
)

type UnidadeDeTrabalho struct {
	Usuario      *t.Usuario
	DB           *sql.DB
	Tx           *sql.Tx
	Repositorios *Repositorios
}

func (u *UnidadeDeTrabalho) Transaction() (tx *sql.Tx, err error) {
	if u.Tx != nil {
		return nil, fmt.Errorf("%w UnidadeDeTrabalho: Já existe uma transação aberta", t.InternalError)
	}

	tx, err = u.DB.Begin()

	if err != nil {
		return nil, fmt.Errorf("%w UnidadeDeTrabalho: Não foi possível iniciar uma transação (%w)", t.InternalError, err)
	}

	u.Tx = tx

	return tx, nil
}

func (u *UnidadeDeTrabalho) Commit() error {
	if u.Tx == nil {
		return fmt.Errorf("%w UnidadeDeTrabalho: Nenhuma transação aberta", t.InternalError)
	}

	if err := u.Tx.Commit(); err != nil {
		return fmt.Errorf("%w UnidadeDeTrabalho: Não foi possível commitar a transação (%w)", t.InternalError, err)
	}

	return nil
}

func (u *UnidadeDeTrabalho) Rollback() error {
	if u.Tx == nil {
		return fmt.Errorf("%w UnidadeDeTrabalho: Nenhuma transação aberta", t.InternalError)
	}

	if err := u.Tx.Rollback(); err != nil {
		return fmt.Errorf("%w UnidadeDeTrabalho: Não foi possível fazer rollback (%w)", t.InternalError, err)
	}

	return nil
}

func (u *UnidadeDeTrabalho) Close() {
	if u.DB != nil {
		u.DB.Close()
	}
}

func Bootstrap(usuario *t.Usuario) *UnidadeDeTrabalho {
	return &UnidadeDeTrabalho{
		Usuario:      usuario,
		DB:           Connect(usuario),
		Repositorios: FabricaDeRepositorios(),
	}
}

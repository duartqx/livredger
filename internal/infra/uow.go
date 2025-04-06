package infra

import (
	"database/sql"
	"fmt"

	t "github.com/duartqx/livredger/internal/common/types"
	r "github.com/duartqx/livredger/internal/infra/repositorios/sqlite"
)

type UnidadeDeTrabalho struct {
	Usuario      *t.Usuario
	DB           *sql.DB
	Tx           *sql.Tx
	Repositorios *Repositorios
}

func (u *UnidadeDeTrabalho) Transaction() (tx *sql.Tx, err error) {
	tx, err = u.DB.Begin()

	if err != nil {
		return tx, fmt.Errorf("UnidadeDeTrabalho: Não foi possível iniciar uma transação (%w)", err)
	}

	u.Tx = tx

	return tx, err
}

func (u *UnidadeDeTrabalho) Commit() error {
	if u.Tx != nil {
		errCommit := u.Tx.Commit()

		if errCommit != nil {
			return fmt.Errorf("UnidadeDeTrabalho: Não foi possível commitar a transação (%w)", errCommit)
		}
	}
	return nil
}

func (u *UnidadeDeTrabalho) Rollback() error {
	if u.Tx != nil {
		err := u.Tx.Rollback()

		if err != nil {
			return fmt.Errorf("UnidadeDeTrabalho: Não foi possível fazer rollback (%w)", err)
		}
	}
	return nil
}

func (u *UnidadeDeTrabalho) Close() {
	if u.DB != nil {
		u.DB.Close()
	}
}

func Bootstrap(usuario *t.Usuario) *UnidadeDeTrabalho {

	conn := r.Connect(usuario)

	return &UnidadeDeTrabalho{
		Usuario:      usuario,
		DB:           conn,
		Repositorios: FabricaDeRepositorios(),
	}
}

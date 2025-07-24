package application

import (
	"context"
	"database/sql"

	"github.com/duartqx/livredger/internal/common/types"
	"github.com/duartqx/livredger/internal/domain/repositorios"
)

type fabricaDeUnidadeDeTrabalho func(context.Context, *types.Usuario) (UnidadeDeTrabalho, error)

type UnidadeDeTrabalho interface {
	Context() context.Context

	Usuario() *types.Usuario
	Repositorios() *repositorios.Repositorios

	DB() *sql.DB

	BeginTransaction() (*sql.Tx, error)
	Transaction() *sql.Tx

	Commit() error
	Rollback() error
	Close()
}

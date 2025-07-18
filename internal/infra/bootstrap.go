package infra

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/duartqx/livredger/internal/common/types"
	"github.com/duartqx/livredger/internal/domain/repositorios"
	"github.com/duartqx/livredger/internal/infra/repositorios/sqlite"
)

// TODO: Refatorar para ser Build target
const DBMS string = "sqlite"

type UnidadeDeTrabalho interface {
	GetContext() context.Context

	GetUsuario() *types.Usuario
	GetRepositorios() *repositorios.Repositorios

	GetDB() *sql.DB

	BeginTransaction() (*sql.Tx, error)
	GetTransaction() *sql.Tx

	Commit() error
	Rollback() error
	Close()
}

func Bootstrap(ctx context.Context, usuario *types.Usuario) (UnidadeDeTrabalho, error) {
	db, err := Connect(ctx, usuario)

	if err != nil {
		return nil, err
	}

	uow := &sqlite.UnidadeDeTrabalhoSqlite{
		Context:      ctx,
		Usuario:      usuario,
		DB:           db,
		Repositorios: FabricaDeRepositorios(),
	}

	return uow, nil
}

func FabricaDeRepositorios() *repositorios.Repositorios {
	switch DBMS {
	case "sqlite":
		return &repositorios.Repositorios{
			Lancamentos:    sqlite.FabricaDeRepositoriosDeLancamento(),
			Contas:         sqlite.FabricaDeRepositoriosDeContas(),
			Demonstrativos: sqlite.FabricaDeRepositoriosDeDemonstrativos(),
		}
	default:
		panic(fmt.Sprintf("Repositorios não configurados para DBMS: {%s}", DBMS))
	}
}

func Connect(ctx context.Context, usuario *types.Usuario) (*sql.DB, error) {

	type conn struct {
		Err error
		Db  *sql.DB
	}

	connChan := make(chan *conn)

	go func(dbChan chan *conn) {

		var conn conn

		switch DBMS {
		case "sqlite":
			conn.Db, conn.Err = sqlite.Connect(usuario)
		default:
			conn.Err = fmt.Errorf("%w: Conexão para {%s} não configurada", types.InternalError, DBMS)
		}

		dbChan <- &conn

	}(connChan)

	select {
	case conn := <-connChan:
		return conn.Db, conn.Err
	case <-ctx.Done():
		return nil, fmt.Errorf("%w: Não foi possível iniciar uma conexão para {%s}", types.TimeOut, DBMS)
	}
}

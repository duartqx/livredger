package infra

import (
	"context"
	"database/sql"
	"fmt"

	ce "github.com/duartqx/livredger/internal/common/errors"
	"github.com/duartqx/livredger/internal/common/types"

	"github.com/duartqx/livredger/internal/application"
	"github.com/duartqx/livredger/internal/domain/repositorios"
	"github.com/duartqx/livredger/internal/infra/repositorios/sqlite"
)

// TODO: Refatorar para ser Build target
const DBMS string = "sqlite"

func FabricaDeUnidadeDeTrabalho(ctx context.Context, usuario *types.Usuario) (application.UnidadeDeTrabalho, error) {
	db, err := Connect(ctx, usuario)

	if err != nil {
		return nil, err
	}

	switch DBMS {
	case "sqlite":
		return sqlite.NewUnidadeDeTrabalhoSqlite(ctx, usuario, db, FabricaDeRepositorios()), nil
	default:
		return nil, fmt.Errorf("%w: UnidadeDeTrabalho não configurada para DBMS: {%s}", ce.InternalError, DBMS)
	}
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
			conn.Err = fmt.Errorf("%w: Conexão para {%s} não configurada", ce.InternalError, DBMS)
		}

		dbChan <- &conn

	}(connChan)

	select {
	case conn := <-connChan:
		return conn.Db, conn.Err
	case <-ctx.Done():
		return nil, fmt.Errorf("%w: Não foi possível iniciar uma conexão para {%s}", ce.TimeOutError, DBMS)
	}
}

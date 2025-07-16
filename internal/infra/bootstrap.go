package infra

import (
	"database/sql"
	"fmt"

	"github.com/duartqx/livredger/internal/common/types"
	"github.com/duartqx/livredger/internal/domain/repositorios"
	"github.com/duartqx/livredger/internal/infra/repositorios/sqlite"
)

type Repositorios struct {
	lancamentos    *repositorios.RepositoriosLancamentos
	contas         *repositorios.RepositoriosContas
	demonstrativos *repositorios.RepositoriosDemonstrativos
}

func (r Repositorios) Lancamentos() *repositorios.RepositoriosLancamentos {
	return r.lancamentos
}

func (r Repositorios) Contas() *repositorios.RepositoriosContas {
	return r.contas
}

func (r Repositorios) Demonstrativos() *repositorios.RepositoriosDemonstrativos {
	return r.demonstrativos
}

// TODO: Refatorar para ser Build target
const DBMS string = "sqlite"

type UnidadeDeTrabalho interface {
	GetUsuario() *types.Usuario
	GetRepositorios() repositorios.Repositorios

	GetDB() *sql.DB

	BeginTransaction() (*sql.Tx, error)
	GetTransaction() *sql.Tx

	Commit() error
	Rollback() error
	Close()
}

func Bootstrap(usuario *types.Usuario) UnidadeDeTrabalho {
	return &sqlite.UnidadeDeTrabalhoSqlite{
		Usuario:      usuario,
		DB:           Connect(usuario),
		Repositorios: FabricaDeRepositorios(),
	}
}

func FabricaDeRepositorios() repositorios.Repositorios {
	switch DBMS {
	case "sqlite":
		return &Repositorios{
			lancamentos:    sqlite.FabricaDeRepositoriosDeLancamento(),
			contas:         sqlite.FabricaDeRepositoriosDeContas(),
			demonstrativos: sqlite.FabricaDeRepositoriosDeDemonstrativos(),
		}
	default:
		panic(fmt.Sprintf("Repositorios não configurados para DBMS: {%s}", DBMS))
	}
}

func Connect(usuario *types.Usuario) *sql.DB {
	switch DBMS {
	case "sqlite":
		return sqlite.Connect(usuario)
	default:
		panic(fmt.Sprintf("Conexão para {%s} não configurada", DBMS))
	}
}

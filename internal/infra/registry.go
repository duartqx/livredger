package infra

import (
	"database/sql"
	"fmt"

	"github.com/duartqx/livredger/internal/common/types"
	"github.com/duartqx/livredger/internal/domain/repositorios"
	"github.com/duartqx/livredger/internal/infra/repositorios/sqlite"
)

type Repositorios struct {
	Lancamentos *repositorios.RepositoriosLancamentos
	Contas      *repositorios.RepositoriosContas
}

// TODO: Refatorar para ser Build target
const DBMS string = "sqlite"

func FabricaDeRepositorios() *Repositorios {
	switch DBMS {
	case "sqlite":
		return &Repositorios{
			Lancamentos: sqlite.FabricaDeRepositoriosDeLancamento(),
			Contas:      sqlite.FabricaDeRepositoriosDeContas(),
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

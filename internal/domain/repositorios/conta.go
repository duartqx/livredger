package repositorios

import (
	"context"
	"database/sql"

	"github.com/duartqx/livredger/internal/domain/comandos"
	"github.com/duartqx/livredger/internal/domain/consultas"
	e "github.com/duartqx/livredger/internal/domain/entidade"
)

type RepositorioDeConsultaContas interface {
	Buscar(ctx context.Context, db *sql.DB, consulta *consultas.ConsultaContas) (*[]*e.Conta, error)
}

type RepositorioDeComandoContas interface {
	Abrir(tx *sql.Tx, comando *comandos.AbrirConta) (*e.Conta, error)
}

type RepositoriosContas struct {
	Comando  RepositorioDeComandoContas
	Consulta RepositorioDeConsultaContas
}

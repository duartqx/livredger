package repositorios

import (
	"context"
	"database/sql"

	"github.com/duartqx/livredger/internal/domain/comandos"
	"github.com/duartqx/livredger/internal/domain/consultas"
	"github.com/duartqx/livredger/internal/domain/entidade"
)

type RepositorioDeConsultaLancamentos interface {
	Buscar(ctx context.Context, db *sql.DB, consulta *consultas.ConsultaLancamentos) (*[]*entidade.Lancamento, error)
}

type RepositorioDeComandoLancamentos interface {
	Criar(ctx context.Context, tx *sql.Tx, comando *comandos.CriarLancamento) (*entidade.Lancamento, error)
	RecalcularTotais(ctx context.Context, tx *sql.Tx, comando *comandos.RecalcularTotais) (int64, error)
}

type RepositoriosLancamentos struct {
	Comando  RepositorioDeComandoLancamentos
	Consulta RepositorioDeConsultaLancamentos
}

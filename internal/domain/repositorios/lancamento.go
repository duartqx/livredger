package repositorios

import (
	"database/sql"

	"github.com/duartqx/livredger/internal/domain/comandos"
	"github.com/duartqx/livredger/internal/domain/consultas"
	"github.com/duartqx/livredger/internal/domain/entidade"
)

type RepositorioDeConsultaLancamentos interface {
	Buscar(db *sql.DB, consulta *consultas.ConsultaLancamentos) (*[]*entidade.Lancamento, error)
}

type RepositorioDeComandoLancamentos interface {
	Criar(tx *sql.Tx, comando *comandos.CriarLancamento) (*entidade.Lancamento, error)
}

package repositorios

import (
	"database/sql"

	"github.com/duartqx/livredger/internal/domain/comandos"
	"github.com/duartqx/livredger/internal/domain/consultas"
	"github.com/duartqx/livredger/internal/domain/entidade"
)

type RepositorioDeComandoDemonstrativoMensal interface {
	Gerar(tx *sql.Tx, comando *comandos.GerarDemonstrativoMensal) (*entidade.DemonstrativoMensal, error)
}

type RepositorioDeConsultaDemonstrativoMensal interface {
	DemonstrativosDosUltimosTresMeses(db *sql.DB, consulta *consultas.ConsultaDemonstrativoMensal) (*[]*entidade.DemonstrativoMensal, error)
	DemonstrativoMensal(db *sql.DB, consulta *consultas.ConsultaDemonstrativoMensal) (*entidade.DemonstrativoMensal, error)
}

type RepositoriosDemonstrativos struct {
	Comando  RepositorioDeComandoDemonstrativoMensal
	Consulta RepositorioDeConsultaDemonstrativoMensal
}

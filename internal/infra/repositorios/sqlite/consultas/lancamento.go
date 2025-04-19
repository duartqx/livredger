package consultas

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	s "github.com/duartqx/livredger/internal/common/sql"
	t "github.com/duartqx/livredger/internal/common/types"
	c "github.com/duartqx/livredger/internal/domain/consultas"
	e "github.com/duartqx/livredger/internal/domain/entidade"
)

type RepositorioDeConsultaLancamentos struct{}

func NewRepositorioDeConsultaLancamentos() *RepositorioDeConsultaLancamentos {
	return &RepositorioDeConsultaLancamentos{}
}

func (r RepositorioDeConsultaLancamentos) Buscar(db *sql.DB, consulta *c.ConsultaLancamentos) (*[]*e.Lancamento, error) {

	builder := s.NewBuilder(
		`SELECT
			id,
			evento,
			timestamp,
			chave,
			versao,
			valores,
			natureza,
			meio,
			vencimento,
			descr
		FROM lancamentos
		{{ .Where }}
		{{ .GroupBy }}
		ORDER BY timestamp DESC`,
	)

	grouping := ""
	if consulta.SomenteVersaoMaisRecente {
		grouping += `GROUP BY chave HAVING max(versao)`
	}

	conds := r.parse(consulta)

	stmt := builder.Render(
		&s.Valores{
			"GroupBy": grouping,
			"Where":   conds.Condicoes(),
		},
	)

	rows, err := r.query(db, stmt, conds.Argumentos())

	if err != nil {
		return nil, err
	}

	lancamentos := make([]*e.Lancamento, 0)

	for rows.Next() {

		var lancamento e.Lancamento

		err := rows.Scan(
			&lancamento.Id,
			&lancamento.Evento,
			&lancamento.Timestamp,
			&lancamento.Chave,
			&lancamento.Versao,
			&lancamento.Valores,
			&lancamento.Natureza,
			&lancamento.Meio,
			&lancamento.Vencimento,
			&lancamento.Descr,
		)

		if err != nil {
			return nil, fmt.Errorf("Não foi possível mapear lançamento: %w", err)
		}

		lancamentos = append(lancamentos, &lancamento)
	}

	return &lancamentos, nil
}

func (r RepositorioDeConsultaLancamentos) parse(consulta *c.ConsultaLancamentos) *s.Condicoes {

	var conds s.Condicoes

	if consulta.Chave != uuid.Nil {
		conds.Add("chave = :chave", sql.Named("chave", consulta.Chave))
	} else {

		if consulta.Description != "" {
			conds.Add("descr LIKE :descr", sql.Named("descr", "%"+consulta.Description+"%"))
		}

		conds.Add(
			"timestamp BETWEEN :inicio AND :final",
			sql.Named("inicio", consulta.Intervalo.Inicio.Format(time.DateOnly)),
			sql.Named("final", consulta.Intervalo.Final.Format(time.DateOnly)),
		)
	}

	return &conds
}

func (r RepositorioDeConsultaLancamentos) query(db *sql.DB, stmt string, args *[]any) (*sql.Rows, error) {
	rows, err := db.Query(stmt, *args...)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%w: Lançamentos não encontrados", t.NotFoundError)
		}
		return nil, err
	}

	return rows, nil
}

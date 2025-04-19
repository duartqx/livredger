package consultas

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	t "github.com/duartqx/livredger/internal/common/types"
	c "github.com/duartqx/livredger/internal/domain/consultas"
	e "github.com/duartqx/livredger/internal/domain/entidade"
)

type RepositorioDeConsultaLancamentos struct{}

func NewRepositorioDeConsultaLancamentos() *RepositorioDeConsultaLancamentos {
	return &RepositorioDeConsultaLancamentos{}
}

func (r RepositorioDeConsultaLancamentos) Buscar(db *sql.DB, consulta *c.ConsultaLancamentos) (*[]*e.Lancamento, error) {

	sql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question)

	builder := sql.
		Select(
			"id",
			"evento",
			"timestamp",
			"chave",
			"versao",
			"valores",
			"natureza",
			"meio",
			"vencimento",
			"descr",
		).
		From("lancamentos").
		OrderBy("timestamp DESC")

	if consulta.SomenteVersaoMaisRecente {
		builder = builder.GroupBy("chave").Having("max(versao)")
	}

	builder = r.parse(consulta, builder)

	stmt, args, err := builder.ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.query(db, stmt, &args)

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

func (r RepositorioDeConsultaLancamentos) parse(consulta *c.ConsultaLancamentos, builder squirrel.SelectBuilder) squirrel.SelectBuilder {

	if consulta.Chave != uuid.Nil {
		return builder.Where(squirrel.Eq{"chave": consulta.Chave})
	}

	if consulta.Description != "" {
		builder = builder.Where(squirrel.Like{"descr": "%" + consulta.Description + "%"})
	}

	return builder.Where(squirrel.And{
		squirrel.GtOrEq{"timestamp": consulta.Intervalo.Inicio.Format(time.DateOnly)},
		squirrel.LtOrEq{"timestamp": consulta.Intervalo.Final.Format(time.DateOnly)},
	})
}

func (r RepositorioDeConsultaLancamentos) query(db *sql.DB, stmt string, args *[]interface{}) (*sql.Rows, error) {
	rows, err := db.Query(stmt, *args...)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%w: Lançamentos não encontrados", t.NotFoundError)
		}
		return nil, err
	}

	return rows, nil
}

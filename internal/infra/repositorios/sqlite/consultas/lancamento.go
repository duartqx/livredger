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

	builder := squirrel.
		Select(
			"id",
			"evento",
			"timestamp",
			"chave",
			"versao",
			"valores",
			"totais",
			"natureza",
			"meio_financeiro",
			"vencimento",
			"descricao",
		).
		From("lancamentos").
		OrderBy(fmt.Sprintf("%s %s", consulta.Paginacao.Ordenacao.Campo, consulta.Paginacao.Ordenacao.Direcao))

	if consulta.SomenteVersaoMaisRecente {
		builder = builder.GroupBy("chave").Having("max(versao)")
	}

	stmt, args, err := r.condicoes(consulta, builder).ToSql()

	if err != nil {
		return nil, fmt.Errorf("%w: %w", t.InternalError, err)
	}

	rows, err := db.Query(stmt, args...)

	lancamentos := make([]*e.Lancamento, 0)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &lancamentos, nil
		}
		return nil, fmt.Errorf("%w: %w", t.InternalError, err)
	}

	defer rows.Close()

	for rows.Next() {

		var lancamento e.Lancamento

		err := rows.Scan(
			&lancamento.Id,
			&lancamento.Evento,
			&lancamento.Timestamp,
			&lancamento.Chave,
			&lancamento.Versao,
			&lancamento.Valores,
			&lancamento.Totais,
			&lancamento.Natureza,
			&lancamento.MeioFinanceiro,
			&lancamento.Vencimento,
			&lancamento.Descricao,
		)

		if err != nil {
			return nil, fmt.Errorf("%w: Não foi possível mapear lançamento: %w", t.InternalError, err)
		}

		lancamentos = append(lancamentos, &lancamento)
	}

	return &lancamentos, nil
}

func (r RepositorioDeConsultaLancamentos) condicoes(
	consulta *c.ConsultaLancamentos, builder squirrel.SelectBuilder,
) squirrel.SelectBuilder {

	if consulta.Chave != uuid.Nil {
		return builder.Where(squirrel.Eq{"chave": consulta.Chave})
	}

	if consulta.Evento != "" {
		builder = builder.Where(squirrel.Eq{"evento": consulta.Evento})
	}

	if consulta.Descricao != "" {
		builder = builder.Where(squirrel.Like{"descricao": "%" + consulta.Descricao + "%"})
	}

	if !consulta.Intervalo.IsZero() {
		if !consulta.Intervalo.Inicio.IsZero() {
			builder = builder.Where(squirrel.GtOrEq{"timestamp": consulta.Intervalo.Inicio.Format(time.DateOnly)})
		}

		if !consulta.Intervalo.Final.IsZero() {
			builder = builder.Where(squirrel.LtOrEq{"timestamp": consulta.Intervalo.Final.Format(time.DateOnly)})
		}
	}

	return builder
}

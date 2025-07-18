package consultas

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/duartqx/livredger/internal/common/types"
	"github.com/duartqx/livredger/internal/domain/consultas"
	"github.com/duartqx/livredger/internal/domain/entidade"
	"github.com/google/uuid"
)

type RepositorioDeConsultaContas struct{}

func NewRepositorioDeConsultaContas() *RepositorioDeConsultaContas {
	return &RepositorioDeConsultaContas{}
}

func (r RepositorioDeConsultaContas) Buscar(
	ctx context.Context, db *sql.DB, consulta *consultas.ConsultaContas,
) (*[]*entidade.Conta, error) {

	builder := squirrel.
		Select(
			"chave",
			"nome",
			"timestamp",
			`
			COALESCE(
				(
					SELECT
						totais
					FROM lancamentos
					WHERE
						lancamentos.evento not in ('ContaAberta')
						AND lancamentos.chave = contas.chave
					GROUP BY chave
					HAVING max(vencimento)
				),
				0
			) AS totais
			`,
		).
		From("contas").
		OrderBy("timestamp ASC")

	if consulta.Chave != uuid.Nil {
		builder = builder.Where(squirrel.Eq{"contas.chave": consulta.Chave.String()})
	} else if strings.Trim(consulta.Nome, " ") != "" {
		builder = builder.Where(squirrel.Eq{"contas.nome": consulta.Nome})
	}

	stmt, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", types.InternalError, err)
	}

	rows, err := db.QueryContext(ctx, stmt, args...)

	contas := make([]*entidade.Conta, 0)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &contas, nil
		}
		return nil, fmt.Errorf("%w: %w", types.InternalError, err)
	}

	defer rows.Close()

	for rows.Next() {

		var conta entidade.Conta

		err := rows.Scan(
			&conta.Chave,
			&conta.Nome,
			&conta.Timestamp,
			&conta.Totais,
		)

		if err != nil {
			return nil, fmt.Errorf("%w: Não foi possível mapear uma conta: %w", types.InternalError, err)
		}

		contas = append(contas, &conta)
	}

	return &contas, nil
}

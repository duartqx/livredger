package consultas

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/duartqx/livredger/internal/common/types"
	"github.com/duartqx/livredger/internal/domain/consultas"
	"github.com/duartqx/livredger/internal/domain/entidade"
)

type RepositorioDeConsultaDemonstrativos struct{}

func NewRepositorioDeConsultaDemonstrativos() *RepositorioDeConsultaDemonstrativos {
	return &RepositorioDeConsultaDemonstrativos{}
}

func (r RepositorioDeConsultaDemonstrativos) DemonstrativosDosUltimosTresMeses(
	ctx context.Context, db *sql.DB, consulta *consultas.ConsultaDemonstrativoMensal,
) (*[]*entidade.DemonstrativoMensal, error) {

	rows, err := db.QueryContext(
		ctx,
		`
		SELECT id, chave, mes, despesa, receita, saldo, timestamp
		FROM demonstrativos_mensais
		WHERE chave = :chave
		GROUP BY mes
		HAVING MAX(timestamp)
		ORDER BY mes DESC
		LIMIT 3
		`,
		sql.Named("chave", consulta.Chave.String()),
	)

	demonstrativos := make([]*entidade.DemonstrativoMensal, 0)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &demonstrativos, nil
		}
		return nil, fmt.Errorf("%w: %w", types.InternalError, err)
	}

	defer rows.Close()

	for rows.Next() {
		var demonstrativo entidade.DemonstrativoMensal

		err := rows.Scan(
			&demonstrativo.Id,
			&demonstrativo.Chave,
			&demonstrativo.Mes,
			&demonstrativo.Despesa,
			&demonstrativo.Receita,
			&demonstrativo.Saldo,
			&demonstrativo.Timestamp,
		)

		if err != nil {
			return nil, fmt.Errorf("%w: Não foi possível mapear demonstrativo: %w", types.InternalError, err)
		}

		demonstrativos = append(demonstrativos, &demonstrativo)
	}

	return &demonstrativos, nil
}

func (r RepositorioDeConsultaDemonstrativos) DemonstrativoMensal(
	ctx context.Context, db *sql.DB, consulta *consultas.ConsultaDemonstrativoMensal,
) (*entidade.DemonstrativoMensal, error) {
	var demonstrativo entidade.DemonstrativoMensal

	if err := db.QueryRowContext(
		ctx,
		`
		SELECT id, chave, mes, despesa, receita, saldo, timestamp
		FROM demonstrativos_mensais
		WHERE chave = :chave AND mes = :mes
		ORDER BY timestamp DESC
		LIMIT 1
		`,
		sql.Named("chave", consulta.Chave.String()),
		sql.Named("mes", consulta.Mes.Format("2006-01")),
	).Scan(
		&demonstrativo.Id,
		&demonstrativo.Chave,
		&demonstrativo.Mes,
		&demonstrativo.Despesa,
		&demonstrativo.Receita,
		&demonstrativo.Saldo,
		&demonstrativo.Timestamp,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf(
				"%w: Não existe demonstrativo gerado para esse mês %s: %w",
				types.NotFoundError, consulta.Mes.Format("2006-01"), err,
			)
		}
		return nil, fmt.Errorf("%w: %w", types.InternalError, err)
	}

	return &demonstrativo, nil
}

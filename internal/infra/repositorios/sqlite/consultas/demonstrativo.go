package consultas

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	ce "github.com/duartqx/livredger/internal/common/errors"
	"github.com/duartqx/livredger/internal/domain/consultas"
	"github.com/duartqx/livredger/internal/domain/entidade"
	"github.com/duartqx/livredger/internal/infra/repositorios/sqlite/helpers"
)

type RepositorioDeConsultaDemonstrativos struct{}

func NewRepositorioDeConsultaDemonstrativos() *RepositorioDeConsultaDemonstrativos {
	return &RepositorioDeConsultaDemonstrativos{}
}

func (r RepositorioDeConsultaDemonstrativos) DemonstrativosDosUltimosSeisMeses(
	ctx context.Context, db *sql.DB, consulta *consultas.ConsultaDemonstrativoUltimosSeisMeses,
) (*[]*entidade.DemonstrativoMensal, error) {

	rows, err := db.QueryContext(
		ctx,
		fmt.Sprintf(`
			SELECT
				demonstrativos.id,
				contas.chave,
				contas.nome,
				contas.timestamp,
				%s,
				demonstrativos.mes,
				demonstrativos.despesa,
				demonstrativos.receita,
				demonstrativos.saldo,
				demonstrativos.timestamp
			FROM demonstrativos_mensais AS demonstrativos
			JOIN contas AS contas ON demonstrativos.chave = contas.chave
			WHERE demonstrativos.chave = :chave
			GROUP BY demonstrativos.mes
			HAVING MAX(demonstrativos.timestamp)
			ORDER BY demonstrativos.mes DESC
			LIMIT 6`,
			helpers.CoalesceTotaisDaConta("demonstrativos.mes", "totais"),
		),
		sql.Named("chave", consulta.Chave.String()),
	)

	demonstrativos := make([]*entidade.DemonstrativoMensal, 0)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &demonstrativos, nil
		}
		return nil, fmt.Errorf("%w: %w", ce.InternalError, err)
	}

	defer rows.Close()

	for rows.Next() {
		var demonstrativo entidade.DemonstrativoMensal

		err := rows.Scan(
			&demonstrativo.Id,
			&demonstrativo.Conta.Chave,
			&demonstrativo.Conta.Nome,
			&demonstrativo.Conta.Timestamp,
			&demonstrativo.Conta.Totais,
			&demonstrativo.Mes,
			&demonstrativo.Despesa,
			&demonstrativo.Receita,
			&demonstrativo.Saldo,
			&demonstrativo.Timestamp,
		)

		if err != nil {
			return nil, fmt.Errorf("%w: Não foi possível mapear demonstrativo: %w", ce.InternalError, err)
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
		fmt.Sprintf(`
			SELECT
				demonstrativos.id,
				contas.chave,
				contas.nome,
				contas.timestamp,
				%s,
				demonstrativos.mes,
				demonstrativos.despesa,
				demonstrativos.receita,
				demonstrativos.saldo,
				demonstrativos.timestamp
			FROM demonstrativos_mensais as demonstrativos
			JOIN contas AS contas ON demonstrativos.chave = contas.chave
			WHERE
				demonstrativos.chave = :chave
				AND demonstrativos.mes = :mes
			ORDER BY demonstrativos.timestamp DESC
			LIMIT 1`,
			helpers.CoalesceTotaisDaConta("demonstrativos.mes", "totais"),
		),
		sql.Named("chave", consulta.Chave.String()),
		sql.Named("mes", consulta.Mes),
	).Scan(
		&demonstrativo.Id,
		&demonstrativo.Conta.Chave,
		&demonstrativo.Conta.Nome,
		&demonstrativo.Conta.Timestamp,
		&demonstrativo.Conta.Totais,
		&demonstrativo.Mes,
		&demonstrativo.Despesa,
		&demonstrativo.Receita,
		&demonstrativo.Saldo,
		&demonstrativo.Timestamp,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf(
				"%w: Não existe demonstrativo gerado para esse mês %s: %w",
				ce.NotFoundError, consulta.Mes, err,
			)
		}
		return nil, fmt.Errorf("%w: %w", ce.InternalError, err)
	}

	return &demonstrativo, nil
}

package comandos

import (
	"context"
	"database/sql"
	"fmt"

	ce "github.com/duartqx/livredger/internal/common/errors"
	c "github.com/duartqx/livredger/internal/domain/comandos"
	e "github.com/duartqx/livredger/internal/domain/entidade"
	"github.com/duartqx/livredger/internal/infra/repositorios/sqlite/regex"
	"github.com/google/uuid"
)

type RepositorioDeComandoLancamentos struct{}

func NewRepositorioDeComandoLancamentos() *RepositorioDeComandoLancamentos {
	return &RepositorioDeComandoLancamentos{}
}

func (r RepositorioDeComandoLancamentos) Criar(ctx context.Context, tx *sql.Tx, comando *c.CriarLancamento) (*e.Lancamento, error) {

	var exists bool

	check := tx.QueryRowContext(
		ctx,
		`SELECT EXISTS(SELECT 1 FROM contas WHERE chave = :chave)`,
		sql.Named("chave", comando.Chave),
	)

	if err := check.Scan(&exists); err != nil {
		return nil, fmt.Errorf("%w: Não foi possível identificar se conta existe: %w", ce.InternalError, err)
	}

	if !exists {
		return nil, fmt.Errorf("%w: Não encontramos uma conta aberta com essa chave '%s'", ce.BusinessLogicError, comando.Chave.String())
	}

	lancamento := e.Lancamento{
		Id:             uuid.New(),
		Evento:         comando.Evento,
		Chave:          comando.Chave,
		Versao:         comando.Versao,
		Valores:        comando.Valores,
		Natureza:       comando.Natureza,
		MeioFinanceiro: comando.MeioFinanceiro,
		Vencimento:     comando.Vencimento,
		Descricao:      comando.Descricao,
	}

	row := tx.QueryRowContext(
		ctx,
		// TODO: LancamentoCancelado aponta via fk para o lançamento cancelado
		// Ele é igual ao lançamento original, mas com valores invertido ( multiplicado por -1 )
		`
		INSERT INTO lancamentos (
			id,
			evento,
			chave,
			versao,
			valores,
			totais,
			natureza,
			meio_financeiro,
			vencimento,
			descricao
		) VALUES (
			:id,
			:evento,
			:chave,
			:versao,
			:valores,
			(
				SELECT COALESCE(SUM(valores), 0) + :valores
				FROM lancamentos WHERE chave = :chave
			),
			:natureza,
			:meio_financeiro,
			:vencimento,
			:descricao
		)
		RETURNING timestamp, totais
		`,
		sql.Named("id", lancamento.Id.String()),
		sql.Named("evento", comando.Evento),
		sql.Named("chave", comando.Chave.String()),
		sql.Named("versao", comando.Versao),
		sql.Named("valores", comando.Valores),
		sql.Named("natureza", comando.Natureza),
		sql.Named("meio_financeiro", comando.MeioFinanceiro),
		sql.Named("vencimento", comando.Vencimento),
		sql.Named("descricao", comando.Descricao),
	)

	if err := row.Scan(&lancamento.Timestamp, &lancamento.Totais); err != nil {
		if match := regex.SqliteFalhouInserirRow.FindStringSubmatch(err.Error()); len(match) > 1 {
			return nil, fmt.Errorf("%w: %s", ce.BusinessLogicError, match[1])
		}

		return nil, fmt.Errorf("%w: Não foi possível inserir novo lançamento: %w", ce.InternalError, err)
	}

	return &lancamento, nil
}

func (r RepositorioDeComandoLancamentos) RecalcularTotais(ctx context.Context, tx *sql.Tx, comando *c.RecalcularTotais) (int64, error) {

	res, err := tx.ExecContext(ctx, `
		WITH totais_recalculados_por_chave AS (
			SELECT
				l1.id AS id,
				(
					SELECT SUM(valores)
					FROM lancamentos AS l2
					WHERE l2.id <= l1.id
				) AS totais_recalculados
			FROM lancamentos AS l1
			WHERE l1.chave = :chave
		)
		UPDATE lancamentos
		SET totais = (
			SELECT trpc.totais_recalculados
			FROM totais_recalculados_por_chave AS trpc
			WHERE trpc.id = lancamentos.id
		)
		WHERE chave = :chave`,
		sql.Named("chave", comando.Chave.String()),
	)

	if err != nil {
		return 0, fmt.Errorf("%w: Não foi possível recalcular totais: %w", ce.InternalError, err)
	}

	aff, err := res.RowsAffected()

	if err != nil {
		return 0, fmt.Errorf("%w: Não foi possível recalcular totais: %w", ce.InternalError, err)
	}

	return aff, nil
}

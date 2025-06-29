package comandos

import (
	"database/sql"
	"fmt"
	"time"

	t "github.com/duartqx/livredger/internal/common/types"
	c "github.com/duartqx/livredger/internal/domain/comandos"
	e "github.com/duartqx/livredger/internal/domain/entidade"
	"github.com/google/uuid"
)

type RepositorioDeComandoLancamentos struct{}

func NewRepositorioDeComandoLancamentos() *RepositorioDeComandoLancamentos {
	return &RepositorioDeComandoLancamentos{}
}

type LancamentoCriado struct {
	Id        int       `json:"id"`
	Timestamp time.Time `json:"timestamp"`
}

func (r RepositorioDeComandoLancamentos) Criar(tx *sql.Tx, comando *c.CriarLancamento) (*e.Lancamento, error) {

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

	row := tx.QueryRow(
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
		if match := sqliteFalhouInserirRow.FindStringSubmatch(err.Error()); len(match) > 1 {
			return nil, fmt.Errorf("%w: %s", t.BusinessLogicError, match[1])
		}

		return nil, fmt.Errorf("%w: Não foi possível inserir novo lançamento: %w", t.InternalError, err)
	}

	return &lancamento, nil
}

func (r RepositorioDeComandoLancamentos) RecalcularTotais(tx *sql.Tx, comando *c.RecalcularTotais) (int64, error) {

	res, err := tx.Exec(
		`
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
		WHERE chave = :chave
		`,
		sql.Named("chave", comando.Chave.String()),
	)

	if err != nil {
		return 0, fmt.Errorf("%w: Não foi possível recalcular totais: %w", t.InternalError, err)
	}

	aff, err := res.RowsAffected()

	if err != nil {
		return 0, fmt.Errorf("%w: Não foi possível recalcular totais: %w", t.InternalError, err)
	}

	return aff, nil
}

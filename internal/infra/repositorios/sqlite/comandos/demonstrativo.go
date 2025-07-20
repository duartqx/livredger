package comandos

import (
	"context"
	"database/sql"
	e "errors"
	"fmt"

	ce "github.com/duartqx/livredger/internal/common/errors"
	"github.com/duartqx/livredger/internal/domain/comandos"
	"github.com/duartqx/livredger/internal/domain/consultas"
	"github.com/duartqx/livredger/internal/domain/entidade"
	"github.com/duartqx/livredger/internal/infra/repositorios/sqlite/helpers"
	"github.com/duartqx/livredger/internal/infra/repositorios/sqlite/regex"
	"github.com/google/uuid"
)

type RepositorioDeComandoDemonstrativos struct{}

func NewRepositorioDeComandoDemonstrativos() *RepositorioDeComandoDemonstrativos {
	return &RepositorioDeComandoDemonstrativos{}
}

func (r RepositorioDeComandoDemonstrativos) Gerar(ctx context.Context, tx *sql.Tx, comando *comandos.GerarDemonstrativoMensal) (*entidade.DemonstrativoMensal, error) {
	demonstrativo := &entidade.DemonstrativoMensal{
		Id:  uuid.New(),
		Mes: comando.Mes.Format("2006-01"),
	}

	stmt, args := helpers.SelectContaComTotais(&consultas.ConsultaContas{Chave: comando.Chave})

	err := tx.QueryRowContext(ctx, stmt, args...).Scan(
		&demonstrativo.Conta.Chave,
		&demonstrativo.Conta.Nome,
		&demonstrativo.Conta.Timestamp,
		&demonstrativo.Conta.Totais,
	)

	if err != nil {
		if e.Is(sql.ErrNoRows, err) {
			return nil, fmt.Errorf("%w: Chave inválida {%s}", ce.BusinessLogicError, comando.Chave.String())
		}
		return nil, err
	}

	err = tx.QueryRowContext(ctx, `
		WITH calculo_do_demonstrativo AS (
			SELECT
				coalesce(sum(lancamentos.valores) FILTER (WHERE lancamentos.valores < 0), 0) as despesa,
				coalesce(sum(lancamentos.valores) FILTER (WHERE lancamentos.valores > 0), 0) as receita,
				coalesce(sum(lancamentos.valores), 0) AS saldo
			FROM lancamentos
			WHERE
				lancamentos.evento NOT IN ('LancamentoTransferido', 'ContaAberta')
				AND lancamentos.chave = :chave
				AND strftime('%Y-%m', lancamentos.vencimento) = :mes
			GROUP BY lancamentos.chave, strftime('%Y-%m', lancamentos.vencimento)
		)
		INSERT INTO demonstrativos_mensais (id, chave, mes, despesa, receita, saldo)
		SELECT :id, :chave, :mes, * FROM calculo_do_demonstrativo
		RETURNING timestamp, despesa, receita, saldo;
		`,
		sql.Named("id", demonstrativo.Id.String()),
		sql.Named("chave", demonstrativo.Conta.Chave.String()),
		sql.Named("mes", demonstrativo.Mes),
	).Scan(
		&demonstrativo.Timestamp,
		&demonstrativo.Despesa,
		&demonstrativo.Receita,
		&demonstrativo.Saldo,
	)

	if err != nil {
		if match := regex.SqliteFalhouInserirRow.FindStringSubmatch(err.Error()); len(match) > 1 {
			return nil, fmt.Errorf("%w: %s", ce.BusinessLogicError, match[1])
		}

		return nil, fmt.Errorf(
			"%w: Não foi possível gerar um demonstrativo para o mês %s: %w",
			ce.InternalError, demonstrativo.Mes, err,
		)
	}

	return demonstrativo, nil
}

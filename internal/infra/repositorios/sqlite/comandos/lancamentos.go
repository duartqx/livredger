package comandos

import (
	"database/sql"
	"fmt"
	"regexp"
	"time"

	t "github.com/duartqx/livredger/internal/common/types"
	c "github.com/duartqx/livredger/internal/domain/comandos"
	e "github.com/duartqx/livredger/internal/domain/entidade"
)

var re = regexp.MustCompile("failed to get next row\nerror code = 1: Error fetching next row: SQLite failure: `(.*?)`")

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
			:evento,
			:chave,
			:versao,
			:valores,
			(
				SELECT COALESCE(SUM(totais), 0) + :valores
				FROM lancamentos WHERE chave = :chave
			),
			:natureza,
			:meio_financeiro,
			:vencimento,
			:descricao
		)
		RETURNING id, timestamp, totais
		`,
		sql.Named("evento", comando.Evento),
		sql.Named("chave", comando.Chave.String()),
		sql.Named("versao", comando.Versao),
		sql.Named("valores", comando.Valores),
		sql.Named("natureza", comando.Natureza),
		sql.Named("meio_financeiro", comando.MeioFinanceiro),
		sql.Named("vencimento", comando.Vencimento),
		sql.Named("descricao", comando.Descricao),
	)

	if err := row.Scan(&lancamento.Id, &lancamento.Timestamp, &lancamento.Totais); err != nil {
		if match := re.FindStringSubmatch(err.Error()); len(match) > 1 {
			return nil, fmt.Errorf("%w: %s", t.BusinessLogicError, match[1])
		}

		return nil, fmt.Errorf("%w: Não foi possível inserir novo lançamento: %w", t.InternalError, err)
	}

	return &lancamento, nil
}

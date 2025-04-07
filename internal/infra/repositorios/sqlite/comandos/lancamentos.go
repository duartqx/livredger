package comandos

import (
	"database/sql"
	"fmt"
	"regexp"
	"time"

	"github.com/duartqx/livredger/internal/domain/comandos"
	e "github.com/duartqx/livredger/internal/domain/entidade"
)

type RepositorioDeComandoLancamentos struct{}

func NewRepositorioDeComandoLancamentos() *RepositorioDeComandoLancamentos {
	return &RepositorioDeComandoLancamentos{}
}

type LancamentoCriado struct {
	Id        int       `json:"id"`
	Timestamp time.Time `json:"timestamp"`
}

func (r RepositorioDeComandoLancamentos) Criar(tx *sql.Tx, comando *comandos.CriarLancamento) (*e.Lancamento, error) {

	lancamento := e.Lancamento{
		Evento:     comando.Evento,
		Chave:      *comando.Chave,
		Versao:     comando.Versao,
		Valores:    comando.Valores,
		Natureza:   comando.Natureza,
		Meio:       comando.Meio,
		Vencimento: comando.Vencimento,
		Descr:      comando.Descr,
	}

	row := tx.QueryRow(
		`
		INSERT INTO lancamentos (
			evento,
			chave,
			versao,
			valores,
			natureza,
			meio,
			vencimento,
			descr
		) VALUES (
			:evento,
			:chave,
			:versao,
			:valores,
			:natureza,
			:meio,
			:vencimento,
			:descr
		)
		RETURNING id, timestamp
		`,
		sql.Named("evento", comando.Evento),
		sql.Named("chave", comando.Chave.String()),
		sql.Named("versao", comando.Versao),
		sql.Named("valores", comando.Valores),
		sql.Named("natureza", comando.Natureza),
		sql.Named("meio", comando.Meio),
		sql.Named("vencimento", comando.Vencimento),
		sql.Named("descr", comando.Descr),
	)

	if err := row.Scan(&lancamento.Id, &lancamento.Timestamp); err != nil {
		re := regexp.MustCompile("failed to get next row\nerror code = 1: Error fetching next row: SQLite failure: `(.*?)`")

		if match := re.FindStringSubmatch(err.Error()); len(match) > 1 {
			return nil, fmt.Errorf("Integridade: %s", match[1])
		}

		return nil, fmt.Errorf("%w: Não foi possível inserir novo lançamento", err)
	}

	return &lancamento, nil
}

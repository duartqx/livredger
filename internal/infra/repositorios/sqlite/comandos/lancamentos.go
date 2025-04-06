package comandos

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/duartqx/livredger/internal/domain/comandos"
)

type RepositorioDeComandoLancamentos struct{}

func NewRepositorioDeComandoLancamentos() *RepositorioDeComandoLancamentos {
	return &RepositorioDeComandoLancamentos{}
}

type LancamentoCriado struct {
	Id        int       `json:"id"`
	Timestamp time.Time `json:"timestamp"`
}

func (r RepositorioDeComandoLancamentos) Criar(tx *sql.Tx, comando *comandos.CriarLancamento) (*LancamentoCriado, error) {

	var lancamento LancamentoCriado

	row := tx.QueryRow(
		`
		-- Chave deve existir para poder criar versao > 1
		-- Só permitir criar nova chave se versao = 1
		-- Se versao != 1 só permitir criar se existir versao = :versao - 1
		INSERT INTO lancamentos (
			evento,
			chave,
			versao,
			valores,
			vencimento,
			descr
		) 
		SELECT
			:evento,
			:chave,
			:versao,
			:valores,
			:vencimento,
			:descr
		WHERE (
			SELECT
				-- Só permite criar se não existir versão com o
				-- mesmo valor e que existe uma versão anterior
				COUNT(*) FILTER (WHERE versao = :versao) = 0
				AND
				(COUNT(*) filter (WHERE versao = :versao - 1) = 1 OR :versao = 1)
			FROM lancamentos
			WHERE chave = :chave
		)
		RETURNING id, timestamp
		`,
		sql.Named("evento", comando.Evento),
		sql.Named("chave", comando.Chave.String()),
		sql.Named("versao", comando.Versao),
		sql.Named("valores", comando.Valores),
		sql.Named("vencimento", comando.Vencimento),
		sql.Named("descr", comando.Descr),
	)

	if err := row.Scan(&lancamento.Id, &lancamento.Timestamp); err != nil {
		return nil, fmt.Errorf("%w: Não foi possível inserir novo lançamento", err)
	}

	return &lancamento, nil
}

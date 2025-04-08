package consultas

import (
	"database/sql"
	"errors"
	"fmt"

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
	base := `
		SELECT
			id,
			evento,
			timestamp,
			chave,
			versao,
			valores,
			natureza,
			meio,
			vencimento,
			descr
		FROM lancamentos
		%s
		%s
		ORDER BY timestamp DESC
	`

	groupBy := ""
	if consulta.SomenteVersaoMaisRecente {
		groupBy += `GROUP BY chave HAVING max(versao)`
	}

	var (
		condicoes  string
		argumentos []any
	)

	if consulta.Chave != uuid.Nil {
		condicoes = "WHERE chave = :chave"

		argumentos = []any{sql.Named("chave", consulta.Chave)}
	} else {
		condicoes = "WHERE timestamp BETWEEN :inicio AND :final "

		argumentos = []any{
			sql.Named("inicio", consulta.Intervalo.Inicio.Format("2006-01-02 15:04:05")),
			sql.Named("final", consulta.Intervalo.Final.Format("2006-01-02 15:04:05")),
		}
	}

	rows, err := db.Query(fmt.Sprintf(base, condicoes, groupBy), argumentos...)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%w: Lançamentos não encontrados", t.NotFoundError)
		}
		return nil, err
	}

	lancamentos := make([]*e.Lancamento, 0)

	for rows.Next() {

		var lancamento e.Lancamento

		err := rows.Scan(
			&lancamento.Id,
			&lancamento.Evento,
			&lancamento.Timestamp,
			&lancamento.Chave,
			&lancamento.Versao,
			&lancamento.Valores,
			&lancamento.Natureza,
			&lancamento.Meio,
			&lancamento.Vencimento,
			&lancamento.Descr,
		)

		if err != nil {
			return nil, fmt.Errorf("Não foi possível mapear lançamento: %w", err)
		}

		lancamentos = append(lancamentos, &lancamento)
	}

	return &lancamentos, nil
}

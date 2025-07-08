package consultas

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/duartqx/livredger/internal/common/types"
	"github.com/duartqx/livredger/internal/domain/consultas"
	"github.com/duartqx/livredger/internal/domain/entidade"
)

type RepositorioDeConsultaContas struct{}

func NewRepositorioDeConstaContas() *RepositorioDeConsultaContas {
	return &RepositorioDeConsultaContas{}
}

func (r RepositorioDeConsultaContas) Buscar(
	db *sql.DB, consulta *consultas.ConsultaContas,
) (*[]*entidade.Conta, error) {

	totais, _, err := squirrel.
		Select("totais").
		From("lancamentos").
		Where("lancamentos.chave = contas.chave").
		OrderBy("lancamentos.timestamp DESC").
		Limit(1).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("%w: %w", types.InternalError, err)
	}

	builder := squirrel.
		Select(
			"chave",
			"nome",
			"timestamp",
			fmt.Sprintf("COALESCE((%s), 0) as totais", totais),
		).
		From("contas").
		OrderBy("timestamp ASC")

	if strings.Trim(consulta.Nome, " ") != "" {
		builder = builder.Where(squirrel.Eq{"contas.nome": consulta.Nome})
	}

	stmt, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", types.InternalError, err)
	}

	rows, err := db.Query(stmt, args...)

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

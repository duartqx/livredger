package consultas

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/duartqx/livredger/internal/common/types"
	"github.com/duartqx/livredger/internal/domain/consultas"
	"github.com/duartqx/livredger/internal/domain/entidade"
	"github.com/duartqx/livredger/internal/infra/repositorios/sqlite/helpers"
)

type RepositorioDeConsultaContas struct{}

func NewRepositorioDeConsultaContas() *RepositorioDeConsultaContas {
	return &RepositorioDeConsultaContas{}
}

func (r RepositorioDeConsultaContas) Buscar(
	ctx context.Context, db *sql.DB, consulta *consultas.ConsultaContas,
) (*[]*entidade.Conta, error) {

	stmt, args := helpers.SelectContaComTotais(consulta)

	rows, err := db.QueryContext(ctx, stmt, args...)

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

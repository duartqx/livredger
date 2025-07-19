package comandos

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/duartqx/livredger/internal/common/types"
	c "github.com/duartqx/livredger/internal/domain/comandos"
	e "github.com/duartqx/livredger/internal/domain/entidade"
	"github.com/duartqx/livredger/internal/infra/repositorios/sqlite/regex"
	"github.com/google/uuid"
)

type RepositorioDeComandoContas struct{}

func NewRepositorioDeComandoContas() *RepositorioDeComandoContas {
	return &RepositorioDeComandoContas{}
}

func (r RepositorioDeComandoContas) Abrir(ctx context.Context, tx *sql.Tx, comando *c.AbrirConta) (*e.Conta, error) {

	conta := e.Conta{
		Chave: uuid.New(),
		Nome:  comando.Nome,
	}

	err := tx.QueryRowContext(ctx, `
		INSERT INTO contas (chave, nome)
		VALUES (:chave, :nome)
		RETURNING timestamp`,
		sql.Named("chave", conta.Chave.String()),
		sql.Named("nome", conta.Nome),
	).Scan(&conta.Timestamp)

	if err != nil {
		if match := regex.SqliteFalhouInserirRow.FindStringSubmatch(err.Error()); len(match) > 1 {
			return nil, fmt.Errorf("%w: %s", types.BusinessLogicError, match[1])
		}

		return nil, fmt.Errorf("%w: Não foi possível abrir uma nova conta: %w", types.InternalError, err)
	}

	return &conta, nil
}

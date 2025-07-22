package helpers

import (
	"cmp"
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"

	"github.com/duartqx/livredger/internal/common/op"
	"github.com/duartqx/livredger/internal/domain/consultas"
)

func SelectContaComTotais(consulta *consultas.ConsultaContas) (string, []any) {
	where_chave := op.Ternary(
		consulta.Chave != uuid.Nil, "WHERE contas.chave = :contas__chave", "",
	)

	where_nome := op.Ternary(
		strings.Trim(consulta.Nome, " ") != "", "WHERE contas.nome = :contas__nome", "",
	)

	query := fmt.Sprintf(
		`SELECT chave, nome, timestamp, %s FROM contas %s ORDER BY timestamp ASC`,
		CoalesceTotaisDaConta("", "totais"),
		*cmp.Or(&where_chave, &where_nome),
	)

	args := []any{
		sql.Named("contas__chave", consulta.Chave.String()),
		sql.Named("contas__nome", consulta.Nome),
	}

	return query, args
}

func CoalesceTotaisDaConta(mes string, alias string) string {
	return fmt.Sprintf(`
		COALESCE(
			(
				SELECT totais FROM lancamentos AS l
				WHERE l.evento NOT IN ('ContaAberta') AND l.chave = contas.chave %s
				ORDER BY vencimento DESC
				LIMIT 1
			),
			0
		) AS %s`,
		op.Ternary(mes != "", "AND strftime('%Y-%m', l.vencimento) = "+mes, ""),
		alias,
	)
}

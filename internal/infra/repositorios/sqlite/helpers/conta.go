package helpers

import (
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/duartqx/livredger/internal/domain/consultas"
	"github.com/google/uuid"
)

func SelectContaComTotais(consulta *consultas.ConsultaContas) (string, []any) {

	builder := squirrel.
		Select(
			"chave",
			"nome",
			"timestamp",
			CoalesceTotaisDaConta("", "totais"),
		).
		From("contas").
		OrderBy("timestamp ASC")

	if consulta.Chave != uuid.Nil {
		builder = builder.Where(squirrel.Eq{"contas.chave": consulta.Chave.String()})
	} else if strings.Trim(consulta.Nome, " ") != "" {
		builder = builder.Where(squirrel.Eq{"contas.nome": consulta.Nome})
	}

	stmt, args, _ := builder.ToSql()

	return stmt, args
}

func CoalesceTotaisDaConta(noMes string, alias string) string {

	fMes := ""
	if noMes != "" {
		fMes = fmt.Sprintf("AND strftime('%%Y-%%m', lancamentos.vencimento) = %s", noMes)
	}

	return fmt.Sprintf(`
		COALESCE(
			(
				SELECT
					totais
				FROM lancamentos
				WHERE
					lancamentos.evento not in ('ContaAberta')
					AND lancamentos.chave = contas.chave
					%s
				ORDER BY vencimento DESC
				LIMIT 1
			),
			0
		) AS %s`,
		fMes,
		alias,
	)
}

package comandos

import (
	"fmt"
	"slices"
	"strings"
	"time"

	ce "github.com/duartqx/livredger/internal/common/errors"
	"github.com/duartqx/livredger/internal/domain/eventos"
	"github.com/duartqx/livredger/internal/domain/value/meios"
	"github.com/duartqx/livredger/internal/domain/value/naturezas"
	"github.com/google/uuid"
)

type CriarLancamento struct {
	Evento         string    `json:"evento"`
	Chave          uuid.UUID `json:"chave"`
	Versao         int       `json:"versao"`
	Valores        float64   `json:"valores"`
	Natureza       string    `json:"natureza"`
	MeioFinanceiro string    `json:"meio_financeiro"`
	Vencimento     time.Time `json:"vencimento"`
	Descricao      string    `json:"descricao"`
}

func (c CriarLancamento) Validate() error {
	if uuid.Nil == c.Chave {
		return fmt.Errorf("%w: Chave é obrigatória", ce.BusinessLogicError)
	}

	if c.Descricao == "" {
		return fmt.Errorf("%w: Descrição é obrigatória", ce.BusinessLogicError)
	}

	if len(c.Descricao) > 500 {
		return fmt.Errorf(
			"%w: Descrição muito longa, deve ter no máximo 500 caracteres",
			ce.BusinessLogicError,
		)
	}

	if !slices.Contains(eventos.EVENTOS_DE_LANCAMENTOS, c.Evento) {
		return fmt.Errorf(
			"%w: Evento não é válido, opções: [%s]",
			ce.BusinessLogicError,
			strings.Join(eventos.EVENTOS_DE_LANCAMENTOS, ", "),
		)
	}

	if c.Versao == 0 {
		return fmt.Errorf("%w: Versão não pode ser igual a 0", ce.BusinessLogicError)
	}

	if !slices.Contains(meios.MEIOS_FINANCEIRO, c.MeioFinanceiro) {
		return fmt.Errorf("%w: Meio Financeiro inválido: %s", ce.BusinessLogicError, c.MeioFinanceiro)
	}

	if !slices.Contains(naturezas.NATUREZAS, c.Natureza) {
		return fmt.Errorf("%w: Natureza da transação inválida: %s", ce.BusinessLogicError, c.Natureza)
	}

	return nil
}

type RecalcularTotais struct {
	Chave uuid.UUID `json:"chave"`
}

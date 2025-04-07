package comandos

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/duartqx/livredger/internal/domain/eventos"
	"github.com/duartqx/livredger/internal/domain/value/meios"
	"github.com/duartqx/livredger/internal/domain/value/naturezas"
	"github.com/google/uuid"
)

var eventosLancamento []string = []string{
	string(eventos.LancamentoCriado),
	string(eventos.LancamentoAtualizado),
	string(eventos.LancamentoPago),
	string(eventos.LancamentoRecebido),
	string(eventos.LancamentoCancelado),
}

type CriarLancamento struct {
	Evento     string     `json:"evento"`
	Chave      *uuid.UUID `json:"chave"`
	Versao     int        `json:"versao"`
	Valores    float64    `json:"valores"`
	Natureza   string     `json:"natureza"`
	Meio       string     `json:"meio"`
	Vencimento time.Time  `json:"vencimento"`
	Descr      string     `json:"descr"`
}

func (c CriarLancamento) Validar() error {
	if c.Descr == "" {
		return fmt.Errorf("Descrição é obrigatória")
	}

	if len(c.Descr) > 500 {
		return fmt.Errorf("Descrição muito longa, deve ter no máximo 500 caracteres")
	}

	if !slices.Contains(eventosLancamento, c.Evento) {
		return fmt.Errorf(
			"Evento não é válido, opções: [%s]",
			strings.Join(eventosLancamento, ", "),
		)
	}

	if c.Versao == 0 {
		return fmt.Errorf("Versão não pode ser igual a 0")
	}

	if !slices.Contains(meios.MEIOS, c.Meio) {
		return fmt.Errorf("Meio de transação inválido: %s", c.Meio)
	}

	if !slices.Contains(naturezas.NATUREZAS, c.Natureza) {
		return fmt.Errorf("Natureza da transação inválida: %s", c.Natureza)
	}

	return nil
}

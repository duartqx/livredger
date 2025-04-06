package comandos

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/duartqx/livredger/internal/domain/eventos"
	"github.com/google/uuid"
)

type CriarLancamento struct {
	Evento     string     `json:"evento"`
	Chave      *uuid.UUID `json:"chave"`
	Versao     int        `json:"versao"`
	Valores    float64    `json:"valores"`
	Vencimento time.Time  `json:"vencimento"`
	Descr      string     `json:"descr"`
}

var eventosLancamentoValido []string = []string{
	string(eventos.LancamentoCriado),
	string(eventos.LancamentoAtualizado),
	string(eventos.LancamentoPago),
	string(eventos.LancamentoRecebido),
	string(eventos.LancamentoCancelado),
}

var errEventoInvalido error = errors.New(
	fmt.Sprintf(
		"Evento não é válido, opções: [%s]",
		strings.Join(eventosLancamentoValido, ", "),
	),
)

func (c CriarLancamento) Validar() error {
	if c.Descr == "" {
		return fmt.Errorf("Descrição é obrigatória")
	}

	if len(c.Descr) > 500 {
		return fmt.Errorf("Descrição muito longa, deve ter no máximo 500 caracteres")
	}

	if !slices.Contains(eventosLancamentoValido, c.Evento) {
		return errEventoInvalido
	}

	if c.Versao == 0 {
		return fmt.Errorf("Versão não pode ser igual a 0")
	}

	return nil
}

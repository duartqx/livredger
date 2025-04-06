package comandos

import (
	"fmt"
	"time"

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

func (c CriarLancamento) Validar() error {
	if c.Descr == "" {
		return fmt.Errorf("Descrição é obrigatória")
	}

	if c.Evento == "" {
		return fmt.Errorf("Tipo inválido")
	}

	if c.Versao == 0 {
		return fmt.Errorf("Versão não pode ser igual a 0")
	}

	return nil
}

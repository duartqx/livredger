package comandos

import (
	"fmt"
	"strings"

	ce "github.com/duartqx/livredger/internal/common/errors"
)

type AbrirConta struct {
	Nome string `json:"nome"`
}

func (a AbrirConta) Validar() error {
	if strings.Trim(a.Nome, " ") == "" {
		return fmt.Errorf("%w: '%s' não é um nome de conta válido", ce.BusinessLogicError, a.Nome)
	}

	return nil
}

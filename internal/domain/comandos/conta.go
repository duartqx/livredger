package comandos

import (
	"fmt"
	"strings"

	"github.com/duartqx/livredger/internal/common/types"
)

type AbrirConta struct {
	Nome string `json:"nome"`
}

func (a AbrirConta) Validar() error {
	if strings.Trim(a.Nome, " ") == "" {
		return fmt.Errorf("%w: '%s' não é um nome de conta válido", types.BusinessLogicError, a.Nome)
	}

	return nil
}

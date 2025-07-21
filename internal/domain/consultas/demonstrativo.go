package consultas

import (
	"fmt"
	"time"

	ce "github.com/duartqx/livredger/internal/common/errors"
	"github.com/google/uuid"
)

type ConsultaDemonstrativoMensal struct {
	Chave uuid.UUID `json:"chave" form:"chave"`
	Mes   string    `json:"mes" form:"mes" help:"Date string formatted as YYYY-MM"`
}

func (c ConsultaDemonstrativoMensal) Validate() error {
	if c.Chave == uuid.Nil {
		return fmt.Errorf("%w: É necessário informar a chave da conta para obter um demonstrativo", ce.BusinessLogicError)
	}

	if _, err := time.Parse("2006-01", c.Mes); err != nil {
		return fmt.Errorf("%w: Mês no formato errado, deve ser YYYY-MM", ce.BusinessLogicError)
	}

	return nil
}

type ConsultaDemonstrativoUltimosSeisMeses struct {
	Chave uuid.UUID `json:"chave" form:"chave"`
}

func (c ConsultaDemonstrativoUltimosSeisMeses) Validate() error {
	if c.Chave == uuid.Nil {
		return fmt.Errorf("%w: É necessário informar a chave da conta para obter um demonstrativo", ce.BusinessLogicError)
	}

	return nil
}

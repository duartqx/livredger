package consultas

import "github.com/google/uuid"

type ConsultaContas struct {
	Nome  string    `json:"nome" form:"nome"`
	Chave uuid.UUID `json:"chave" form:"chave"`
}

func ConsultaContasPadrao() *ConsultaContas {
	return &ConsultaContas{}
}

package entidade

import (
	"time"

	"github.com/google/uuid"
)

type DemonstrativoMensal struct {
	Id        uuid.UUID `json:"id"`
	Chave     uuid.UUID `json:"chave"`
	Mes       string    `json:"mes"`
	Despesa   float64   `json:"despesa"`
	Receita   float64   `json:"receita"`
	Saldo     float64   `json:"saldo"`
	Timestamp time.Time `json:"timestamp"`
}

package entidade

import (
	"time"

	"github.com/google/uuid"
)

type Conta struct {
	Chave     uuid.UUID `json:"chave"`
	Nome      string    `json:"nome"`
	Totais    float64   `json:"totais"`
	Timestamp time.Time `json:"timestamp"`
}

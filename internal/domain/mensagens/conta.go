package mensagens

import (
	"time"

	"github.com/google/uuid"
)

type ContaAberta struct {
	Chave     uuid.UUID `json:"chave"`
	Timestamp time.Time `json:"timestamp"`
}

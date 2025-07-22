package mensagens

import (
	"time"

	"github.com/google/uuid"
)

type ContaAberta struct {
	EId       uuid.UUID `json:"event_identifier"`
	Chave     uuid.UUID `json:"chave"`
	Timestamp time.Time `json:"timestamp"`
}

func (ca ContaAberta) GetEventIdentifier() uuid.UUID {
	return ca.EId
}

func (ca ContaAberta) GetEntityIdentifier() any {
	return ca.Chave.String()
}

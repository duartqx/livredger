package mensagens

import (
	"time"

	"github.com/google/uuid"

	"github.com/duartqx/livredger/internal/domain"
)

type LancamentoCriado struct {
	EId       uuid.UUID    `json:"event_identifier"`
	Id        uuid.UUID    `json:"id"`
	Evento    domain.Event `json:"evento"`
	Timestamp time.Time    `json:"timestamp"`
}

func (lc LancamentoCriado) GetEventIdentifier() uuid.UUID {
	return lc.EId
}

func (lc LancamentoCriado) GetEntityIdentifier() any {
	return lc.Id.String()
}

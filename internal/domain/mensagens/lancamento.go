package mensagens

import (
	"time"

	"github.com/google/uuid"

	"github.com/duartqx/livredger/internal/domain"
)

type LancamentoCriado struct {
	Id        uuid.UUID     `json:"id"`
	Evento    domain.Evento `json:"evento"`
	Timestamp time.Time     `json:"timestamp"`
}

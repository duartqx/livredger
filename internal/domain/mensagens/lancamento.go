package mensagens

import (
	"time"

	"github.com/duartqx/livredger/internal/common/types"
	"github.com/google/uuid"
)

type LancamentoCriado struct {
	Id        uuid.UUID    `json:"id"`
	Evento    types.Evento `json:"evento"`
	Timestamp time.Time    `json:"timestamp"`
}

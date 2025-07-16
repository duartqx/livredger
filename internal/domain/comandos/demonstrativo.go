package comandos

import (
	"time"

	"github.com/google/uuid"
)

type GerarDemonstrativoMensal struct {
	Chave uuid.UUID `json:"chave"`
	Mes   time.Time `json:"mes"`
}

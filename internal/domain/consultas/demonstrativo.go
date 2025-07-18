package consultas

import (
	"time"

	"github.com/google/uuid"
)

type ConsultaDemonstrativoMensal struct {
	Chave uuid.UUID `json:"chave"`
	Mes   time.Time `json:"mes"`
}

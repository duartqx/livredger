package consultas

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	t "github.com/duartqx/livredger/internal/common/types"
)

type ConsultaLancamentos struct {
	Chave                    uuid.UUID    `json:"chave"`
	SomenteVersaoMaisRecente bool         `json:"somente_versao_mais_recente"`
	Intervalo                *t.Intervalo `json:"intervalo"`
	Description              string       `json:"description"`
}

func ParsearStringsParaConsultaLancamentos(
	chaveStr string,
	somenteVersaoMaisRecenteStr string,
	intervaloInicioStr string,
	intervaloFinalStr string,
	description string,
) (*ConsultaLancamentos, error) {

	consulta := &ConsultaLancamentos{
		SomenteVersaoMaisRecente: somenteVersaoMaisRecenteStr == "true",
		Description:              description,
	}

	if chaveStr != "" {
		if chaveUUID, err := uuid.Parse(chaveStr); err == nil {
			consulta.Chave = chaveUUID
		} else {
			return nil, fmt.Errorf("%w: Chave inválida", err)
		}
	}

	intervalo, err := t.ParseIntervalo(intervaloInicioStr, intervaloFinalStr)

	if err != nil {
		return nil, err
	}

	if intervalo.IsZero() {
		now := time.Now()

		intervalo = &t.Intervalo{
			Inicio: time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()),
			Final:  now,
		}
	}

	consulta.Intervalo = intervalo

	return consulta, nil
}

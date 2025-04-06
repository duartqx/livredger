package consultas

import (
	"fmt"

	"github.com/google/uuid"

	t "github.com/duartqx/livredger/internal/common/types"
)

type ConsultaLancamentos struct {
	Chave                    uuid.UUID    `json:"chave"`
	SomenteVersaoMaisRecente bool         `json:"somente_versao_mais_recente"`
	Intervalo                *t.Intervalo `json:"intervalo"`
}

func ParsearStringsParaConsultaLancamentos(
	chaveStr string,
	SomenteVersaoMaisRecenteStr string,
	IntervaloInicioStr string,
	IntervaloFinalStr string,
) (*ConsultaLancamentos, error) {

	consulta := &ConsultaLancamentos{
		SomenteVersaoMaisRecente: SomenteVersaoMaisRecenteStr == "true",
	}

	if chaveStr != "" {
		if chaveUUID, err := uuid.Parse(chaveStr); err == nil {
			consulta.Chave = chaveUUID
		} else {
			return nil, fmt.Errorf("%w: Chave inválida", err)
		}
	}

	intervalo, err := t.ParseIntervalo(IntervaloInicioStr, IntervaloFinalStr)

	if err != nil {
		return nil, err
	}

	consulta.Intervalo = intervalo

	return consulta, nil
}

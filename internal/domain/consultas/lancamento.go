package consultas

import (
	"encoding/json"

	"github.com/google/uuid"

	t "github.com/duartqx/livredger/internal/common/types"
)

type ConsultaLancamentos struct {
	Chave                    uuid.UUID   `json:"chave"`
	SomenteVersaoMaisRecente bool        `json:"somente_versao_mais_recente"`
	Intervalo                t.Intervalo `json:"intervalo"`
	Descricao                string      `json:"descricao"`
}

func ConsultaLancamentosPadrao() *ConsultaLancamentos {
	return &ConsultaLancamentos{
		SomenteVersaoMaisRecente: true,
		Intervalo:                t.IntervaloDesseMes(),
	}
}

func (cl *ConsultaLancamentos) UnmarshalJSON(data []byte) error {
	type Alias ConsultaLancamentos

	aux := &struct{ *Alias }{Alias: (*Alias)(cl)}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.Intervalo.IsZero() {
		aux.Intervalo = t.IntervaloDesseMes()
	}

	cl = (*ConsultaLancamentos)(aux.Alias)

	return nil
}

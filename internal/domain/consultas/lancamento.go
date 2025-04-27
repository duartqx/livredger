package consultas

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"github.com/google/uuid"

	"github.com/duartqx/livredger/internal/common/types"
	t "github.com/duartqx/livredger/internal/common/types"
	e "github.com/duartqx/livredger/internal/domain/eventos"
)

type ConsultaLancamentos struct {
	Chave                    uuid.UUID   `json:"chave"`
	SomenteVersaoMaisRecente bool        `json:"somente_versao_mais_recente"`
	Evento                   string      `json:"evento"`
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

	if aux.Evento != "" && !slices.Contains(e.EVENTOS_DE_LANCAMENTOS, aux.Evento) {
		return fmt.Errorf("%w: %s não é um evento válido [%s]", types.BusinessLogicError, aux.Evento, strings.Join(e.EVENTOS_DE_LANCAMENTOS, ", "))
	}

	if aux.Intervalo.IsZero() {
		aux.Intervalo = t.IntervaloDesseMes()
	}

	cl = (*ConsultaLancamentos)(aux.Alias)

	return nil
}

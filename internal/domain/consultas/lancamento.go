package consultas

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"github.com/google/uuid"

	ce "github.com/duartqx/livredger/internal/common/errors"
	"github.com/duartqx/livredger/internal/common/types"
	"github.com/duartqx/livredger/internal/domain/eventos"
)

type ConsultaLancamentos struct {
	Chave                    uuid.UUID       `json:"chave" form:"chave"`
	SomenteVersaoMaisRecente bool            `json:"somente_versao_mais_recente" form:"somente_versao_mais_recente"`
	Evento                   string          `json:"evento" form:"evento"`
	Intervalo                types.Intervalo `json:"intervalo" form:"intervalo"`
	Descricao                string          `json:"descricao" form:"descricao"`
	Paginacao                types.Paginacao `json:"paginacao" form:"paginacao"`
}

func ConsultaLancamentosPadrao() *ConsultaLancamentos {
	return &ConsultaLancamentos{
		SomenteVersaoMaisRecente: true,
		Intervalo:                types.Intervalo{},
		Paginacao:                types.Paginacao{Pagina: 0, Ordenacao: types.Ordenacao{Campo: "timestamp", Direcao: "DESC"}},
	}
}

func (cl *ConsultaLancamentos) UnmarshalJSON(data []byte) error {
	type Alias ConsultaLancamentos

	aux := &struct{ *Alias }{Alias: (*Alias)(cl)}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.Evento != "" && !slices.Contains(eventos.EVENTOS_DE_LANCAMENTOS, aux.Evento) {
		return fmt.Errorf(
			"%w: %s não é um evento válido [%s]",
			ce.BusinessLogicError,
			aux.Evento,
			strings.Join(eventos.EVENTOS_DE_LANCAMENTOS, ", "),
		)
	}

	if aux.Intervalo.IsZero() {
		aux.Intervalo = types.IntervaloDesseMes()
	}

	cl = (*ConsultaLancamentos)(aux.Alias)

	return nil
}

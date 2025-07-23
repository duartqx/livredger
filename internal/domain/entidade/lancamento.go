package entidade

import (
	"time"

	m "github.com/duartqx/livredger/internal/domain/value/meios"
	n "github.com/duartqx/livredger/internal/domain/value/naturezas"
	"github.com/google/uuid"
)

type Lancamento struct {
	Id             uuid.UUID        `json:"id"`
	Evento         string           `json:"evento"`
	Chave          uuid.UUID        `json:"chave"`
	Timestamp      time.Time        `json:"timestamp"`
	Vencimento     time.Time        `json:"vencimento"`
	Versao         int              `json:"versao"`
	Valores        float64          `json:"valores"`
	Totais         float64          `json:"totais"`
	Natureza       n.Natureza       `json:"natureza"`
	MeioFinanceiro m.MeioFinanceiro `json:"meio_financeiro"`
	Descricao      string           `json:"descricao"`
}

func NovoLancamento() *Lancamento {
	return &Lancamento{
		Timestamp: time.Now(),
	}
}

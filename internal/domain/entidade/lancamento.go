package entidade

import (
	"time"

	"github.com/google/uuid"
)

type Lancamento struct {
	Id        int       `json:"id"`
	Evento    string    `json:"evento"`
	Timestamp time.Time `json:"timestamp"`

	Chave  uuid.UUID `json:"chave"`
	Versao int       `json:"versao"`

	Valores        float64   `json:"valores"`
	Natureza       string    `json:"natureza"`
	MeioFinanceiro string    `json:"meio_financeiro"`
	Vencimento     time.Time `json:"vencimento"`
	Descricao      string    `json:"descricao"`
}

func NovoLancamento() *Lancamento {
	return &Lancamento{
		Timestamp: time.Now(),
	}
}

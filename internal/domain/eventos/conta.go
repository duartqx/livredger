package eventos

import "github.com/duartqx/livredger/internal/domain"

const (
	ContaAberta  = domain.Event("ContaAberta")
	ContaFechada = domain.Event("ContaFechada")
	ContaPausada = domain.Event("ContaPausada")
)

var EVENTOS_DE_CONTA []string = []string{
	string(ContaAberta),
	string(ContaFechada),
	string(ContaPausada),
}

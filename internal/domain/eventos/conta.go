package eventos

import "github.com/duartqx/livredger/internal/domain"

const (
	ContaAberta  = domain.Evento("ContaAberta")
	ContaFechada = domain.Evento("ContaFechada")
	ContaPausada = domain.Evento("ContaPausada")
)

var EVENTOS_DE_CONTA []string = []string{
	string(ContaAberta),
	string(ContaFechada),
	string(ContaPausada),
}

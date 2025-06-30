package eventos

import "github.com/duartqx/livredger/internal/common/types"

const (
	ContaAberta  = types.Evento("ContaAberta")
	ContaFechada = types.Evento("ContaFechada")
	ContaPausada = types.Evento("ContaPausada")
)

var EVENTOS_DE_CONTA []string = []string{
	string(ContaAberta),
	string(ContaFechada),
	string(ContaPausada),
}

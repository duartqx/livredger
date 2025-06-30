package eventos

import (
	"github.com/duartqx/livredger/internal/domain"
)

const (
	LancamentoPrevisto    = domain.Evento("LancamentoPrevisto")
	LancamentoPago        = domain.Evento("LancamentoPago")
	LancamentoRecebido    = domain.Evento("LancamentoRecebido")
	LancamentoCancelado   = domain.Evento("LancamentoCancelado")
	LancamentoCorrigido   = domain.Evento("LancamentoCorrigido")
	LancamentoTransferido = domain.Evento("LancamentoTransferido")
)

var EVENTOS_DE_LANCAMENTOS []string = []string{
	string(LancamentoPrevisto),
	string(LancamentoPago),
	string(LancamentoRecebido),
	string(LancamentoCancelado),
	string(LancamentoCorrigido),
	string(LancamentoTransferido),
}

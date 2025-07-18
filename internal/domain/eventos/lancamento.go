package eventos

import (
	"github.com/duartqx/livredger/internal/domain"
)

const (
	LancamentoPrevisto    = domain.Event("LancamentoPrevisto")
	LancamentoPago        = domain.Event("LancamentoPago")
	LancamentoRecebido    = domain.Event("LancamentoRecebido")
	LancamentoCancelado   = domain.Event("LancamentoCancelado")
	LancamentoCorrigido   = domain.Event("LancamentoCorrigido")
	LancamentoTransferido = domain.Event("LancamentoTransferido")
)

var EVENTOS_DE_LANCAMENTOS []string = []string{
	string(LancamentoPrevisto),
	string(LancamentoPago),
	string(LancamentoRecebido),
	string(LancamentoCancelado),
	string(LancamentoCorrigido),
	string(LancamentoTransferido),
}

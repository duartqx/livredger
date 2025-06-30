package eventos

import (
	"github.com/duartqx/livredger/internal/common/types"
)

const (
	LancamentoPrevisto    = types.Evento("LancamentoPrevisto")
	LancamentoPago        = types.Evento("LancamentoPago")
	LancamentoRecebido    = types.Evento("LancamentoRecebido")
	LancamentoCancelado   = types.Evento("LancamentoCancelado")
	LancamentoCorrigido   = types.Evento("LancamentoCorrigido")
	LancamentoTransferido = types.Evento("LancamentoTransferido")
)

var EVENTOS_DE_LANCAMENTOS []string = []string{
	string(LancamentoPrevisto),
	string(LancamentoPago),
	string(LancamentoRecebido),
	string(LancamentoCancelado),
	string(LancamentoCorrigido),
	string(LancamentoTransferido),
}

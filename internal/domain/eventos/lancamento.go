package eventos

import t "github.com/duartqx/livredger/internal/common/types"

const (
	LancamentoCriado     = t.Event("LancamentoCriado")
	LancamentoPrevisto   = t.Event("LancamentoPrevisto")
	LancamentoAtualizado = t.Event("LancamentoAtualizado")
	LancamentoPago       = t.Event("LancamentoPago")
	LancamentoRecebido   = t.Event("LancamentoRecebido")
	LancamentoCancelado  = t.Event("LancamentoCancelado")
)

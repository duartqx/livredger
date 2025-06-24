package eventos

const (
	LancamentoPrevisto    = "LancamentoPrevisto"
	LancamentoPago        = "LancamentoPago"
	LancamentoRecebido    = "LancamentoRecebido"
	LancamentoCancelado   = "LancamentoCancelado"
	LancamentoCorrigido   = "LancamentoCorrigido"
	LancamentoTransferido = "LancamentoTransferido"
)

var EVENTOS_DE_LANCAMENTOS []string = []string{
	LancamentoPrevisto,
	LancamentoPago,
	LancamentoRecebido,
	LancamentoCancelado,
	LancamentoCorrigido,
	LancamentoTransferido,
}

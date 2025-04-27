package eventos

const (
	LancamentoPrevisto  = "LancamentoPrevisto"
	LancamentoPago      = "LancamentoPago"
	LancamentoRecebido  = "LancamentoRecebido"
	LancamentoCancelado = "LancamentoCancelado"
)

var EVENTOS_DE_LANCAMENTOS []string = []string{
	LancamentoPrevisto,
	LancamentoPago,
	LancamentoRecebido,
	LancamentoCancelado,
}

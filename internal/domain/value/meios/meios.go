package meios

type MeioFinanceiro string

const (
	BENEFICIOS             MeioFinanceiro = "Cartão de Benefícios"
	BOLETO_BANCARIO        MeioFinanceiro = "Boleto Bancário"
	CARAO_DEBITO           MeioFinanceiro = "Cartão de Débito"
	CARTAO_CREDITO         MeioFinanceiro = "Cartão de Crédito"
	DINHEIRO               MeioFinanceiro = "Dinheiro"
	PIX                    MeioFinanceiro = "PIX"
	POUPANCA               MeioFinanceiro = "Poupança"
	TRANSFERENCIA_BANCARIA MeioFinanceiro = "Transferência Bancária"
	OUTRO                  MeioFinanceiro = "Outro"
)

var MEIOS_FINANCEIRO []MeioFinanceiro = []MeioFinanceiro{
	BENEFICIOS,
	BOLETO_BANCARIO,
	CARAO_DEBITO,
	CARTAO_CREDITO,
	DINHEIRO,
	OUTRO,
	PIX,
	POUPANCA,
	TRANSFERENCIA_BANCARIA,
}

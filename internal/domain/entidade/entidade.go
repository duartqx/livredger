package entidade

type Entidade interface {
	Lancamento | Conta | DemonstrativoMensal
}

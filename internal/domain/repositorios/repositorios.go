package repositorios

type Repositorios interface {
	Lancamentos() *RepositoriosLancamentos
	Contas() *RepositoriosContas
	Demonstrativos() *RepositoriosDemonstrativos
}

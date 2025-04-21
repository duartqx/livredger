package sqlite

import (
	"github.com/duartqx/livredger/internal/domain/repositorios"

	"github.com/duartqx/livredger/internal/infra/repositorios/sqlite/comandos"
	"github.com/duartqx/livredger/internal/infra/repositorios/sqlite/consultas"
)

func FabricaDeRepositoriosDeLancamento() *repositorios.RepositoriosLancamentos {
	return &repositorios.RepositoriosLancamentos{
		Comando:  comandos.NewRepositorioDeComandoLancamentos(),
		Consulta: consultas.NewRepositorioDeConsultaLancamentos(),
	}
}

package infra

import (
	"fmt"

	"github.com/duartqx/livredger/internal/domain/repositorios"
	"github.com/duartqx/livredger/internal/infra/repositorios/sqlite/comandos"
	"github.com/duartqx/livredger/internal/infra/repositorios/sqlite/consultas"
)

type RepositoriosLancamento struct {
	Comando  repositorios.RepositorioDeComandoLancamentos
	Consulta repositorios.RepositorioDeConsultaLancamentos
}

type Repositorios struct {
	Lancamento *RepositoriosLancamento
}

const DBMS string = "sqlite"

func FabricaDeRepositorios() *Repositorios {
	switch DBMS {
	case "sqlite":
		return &Repositorios{
			Lancamento: &RepositoriosLancamento{
				Comando:  comandos.NewRepositorioDeComandoLancamentos(),
				Consulta: consultas.NewRepositorioDeConsultaLancamentos(),
			},
		}
	default:
		panic(fmt.Sprintf("Repositorios não configurados para DBMS: {%s}", DBMS))
	}
}

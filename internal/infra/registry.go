package infra

import (
	"fmt"

	"github.com/duartqx/livredger/internal/domain/repositorios"
	"github.com/duartqx/livredger/internal/infra/repositorios/sqlite/comandos"
	"github.com/duartqx/livredger/internal/infra/repositorios/sqlite/consultas"
)

type RepositoriosLancamentos struct {
	Comando  repositorios.RepositorioDeComandoLancamentos
	Consulta repositorios.RepositorioDeConsultaLancamentos
}

type Repositorios struct {
	Lancamentos *RepositoriosLancamentos
}

const DBMS string = "sqlite"

func FabricaDeRepositorios() *Repositorios {
	switch DBMS {
	case "sqlite":
		return &Repositorios{
			Lancamentos: &RepositoriosLancamentos{
				Comando:  comandos.NewRepositorioDeComandoLancamentos(),
				Consulta: consultas.NewRepositorioDeConsultaLancamentos(),
			},
		}
	default:
		panic(fmt.Sprintf("Repositorios não configurados para DBMS: {%s}", DBMS))
	}
}

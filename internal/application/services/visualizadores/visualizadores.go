package visualizadores

import (
	"fmt"

	"github.com/google/uuid"

	t "github.com/duartqx/livredger/internal/common/types"
	c "github.com/duartqx/livredger/internal/domain/consultas"
	e "github.com/duartqx/livredger/internal/domain/entidade"
	i "github.com/duartqx/livredger/internal/infra"
)

func BuscarLancamentos(uow *i.UnidadeDeTrabalho, consulta *c.ConsultaLancamentos) (*[]*e.Lancamento, error) {

	if consulta.Chave == uuid.Nil && consulta.Intervalo.IsZero() {
		return nil, fmt.Errorf("%w: Proibido buscar lançamentos sem chave sem a presença de intervalo de tempo", t.BusinessLogicError)
	}

	return uow.Repositorios.Lancamento.Consulta.Buscar(uow.DB, consulta)
}

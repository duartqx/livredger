package visualizadores

import (
	"fmt"

	t "github.com/duartqx/livredger/internal/common/types"
	c "github.com/duartqx/livredger/internal/domain/consultas"
	e "github.com/duartqx/livredger/internal/domain/entidade"
	i "github.com/duartqx/livredger/internal/infra"
	r "github.com/duartqx/livredger/internal/infra/repositorios/sqlite/consultas"
	"github.com/google/uuid"
)

func BuscarLancamentos(uow *i.UnidadeDeTrabalho, consulta *c.ConsultaLancamentos) (*[]*e.Lancamento, error) {

	if consulta.Chave == uuid.Nil && (consulta.Intervalo.Inicio.IsZero() && consulta.Intervalo.Final.IsZero()) {
		return nil, fmt.Errorf("%w: Proibido buscar lançamentos sem chave sem a presença de intervalo de tempo", t.BusinessLogicError)
	}

	repo := r.NewRepositorioDeConsultaLancamentos()

	return repo.Buscar(uow.DB, consulta)
}

package executores

import (
	"database/sql"

	"github.com/duartqx/livredger/internal/common/types"
	c "github.com/duartqx/livredger/internal/domain/comandos"
	e "github.com/duartqx/livredger/internal/domain/entidade"
	i "github.com/duartqx/livredger/internal/infra"
)

type Executor[Comando types.Comando, Entidade any] struct{}

func (e Executor[Comando, Entidade]) TransactionalScript(
	uow *i.UnidadeDeTrabalho,
	comando Comando,
	fn func(*sql.Tx, Comando) (*Entidade, error),
) (*Entidade, error) {
	if err := comando.Validar(); err != nil {
		return nil, err
	}

	tx, err := uow.Transaction()

	if err != nil {
		return nil, err
	}

	resultado, err := fn(tx, comando)

	if err != nil {
		uow.Rollback()
		return nil, err
	}

	if err := uow.Commit(); err != nil {
		uow.Rollback()
		return nil, err
	}

	return resultado, nil
}

func CriarLancamento(uow *i.UnidadeDeTrabalho, comando *c.CriarLancamento) (*e.Lancamento, error) {
	var executor Executor[*c.CriarLancamento, e.Lancamento]

	return executor.TransactionalScript(
		uow, comando, uow.Repositorios.Lancamentos.Comando.Criar,
	)
}

func AbrirConta(uow *i.UnidadeDeTrabalho, comando *c.AbrirConta) (*e.Conta, error) {
	var executor Executor[*c.AbrirConta, e.Conta]

	return executor.TransactionalScript(
		uow, comando, uow.Repositorios.Contas.Comando.Abrir,
	)
}

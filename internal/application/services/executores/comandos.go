package executores

import (
	"database/sql"

	"github.com/duartqx/livredger/internal/application/messagebus"

	"github.com/duartqx/livredger/internal/domain"
	"github.com/duartqx/livredger/internal/domain/comandos"
	"github.com/duartqx/livredger/internal/domain/entidade"
	"github.com/duartqx/livredger/internal/domain/mensagens"

	"github.com/duartqx/livredger/internal/infra"
)

func TransactionalScript[Entidade entidade.Entidade](
	uow infra.UnidadeDeTrabalho, fn func(*sql.Tx) (*Entidade, error),
) (*Entidade, error) {
	tx, err := uow.BeginTransaction()

	if err != nil {
		return nil, err
	}

	resultado, err := fn(tx)

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

func CriarLancamento(uow infra.UnidadeDeTrabalho, comando *comandos.CriarLancamento) (*entidade.Lancamento, error) {
	return TransactionalScript(
		uow,
		func(tx *sql.Tx) (*entidade.Lancamento, error) {
			lancamento, err := uow.GetRepositorios().Lancamentos.Comando.Criar(tx, comando)

			if err != nil {
				return nil, err
			}

			errCh := messagebus.MessageBus.Publish(
				uow, mensagens.LancamentoCriado{
					Id:        lancamento.Id,
					Evento:    domain.Event(lancamento.Evento),
					Timestamp: lancamento.Timestamp,
				},
			)

			if err := <-errCh; err != nil {
				return nil, err
			}

			return lancamento, nil
		},
	)
}

func AbrirConta(uow infra.UnidadeDeTrabalho, comando *comandos.AbrirConta) (*entidade.Conta, error) {
	return TransactionalScript(
		uow,
		func(tx *sql.Tx) (*entidade.Conta, error) {
			conta, err := uow.GetRepositorios().Contas.Comando.Abrir(tx, comando)

			if err != nil {
				return nil, err
			}

			errCh := messagebus.MessageBus.Publish(
				uow, mensagens.ContaAberta{Chave: conta.Chave, Timestamp: conta.Timestamp},
			)

			if err := <-errCh; err != nil {
				return nil, err
			}

			return conta, nil
		},
	)
}

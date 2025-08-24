package executores

import (
	"database/sql"

	"github.com/duartqx/livredger/internal/application"
	"github.com/duartqx/livredger/internal/application/messagebus"
	"github.com/google/uuid"

	"github.com/duartqx/livredger/internal/domain"
	"github.com/duartqx/livredger/internal/domain/comandos"
	"github.com/duartqx/livredger/internal/domain/entidade"
	"github.com/duartqx/livredger/internal/domain/mensagens"
)

func TransactionalScript[Entidade entidade.Entidade](
	uow application.UnidadeDeTrabalho, executor func(*sql.Tx) (*Entidade, error),
) (*Entidade, error) {
	tx, err := uow.BeginTransaction()

	if err != nil {
		return nil, err
	}

	defer uow.Rollback()

	resultado, err := executor(tx)

	if err != nil {
		return nil, err
	}

	if err := uow.Commit(); err != nil {
		return nil, err
	}

	return resultado, nil
}

func CriarLancamento(uow application.UnidadeDeTrabalho, comando *comandos.CriarLancamento) (*entidade.Lancamento, error) {
	return TransactionalScript(
		uow,
		func(tx *sql.Tx) (*entidade.Lancamento, error) {
			lancamento, err := uow.Repositorios().Lancamentos.Comando.Criar(uow.Context(), tx, comando)

			if err != nil {
				return nil, err
			}

			err = messagebus.MessageBus.Publish(
				uow, &mensagens.LancamentoCriado{
					EId:       uuid.New(),
					Id:        lancamento.Id,
					Evento:    domain.Event(lancamento.Evento),
					Timestamp: lancamento.Timestamp,
				},
			)

			if err != nil {
				return nil, err
			}

			return lancamento, nil
		},
	)
}

func AbrirConta(uow application.UnidadeDeTrabalho, comando *comandos.AbrirConta) (*entidade.Conta, error) {
	return TransactionalScript(
		uow,
		func(tx *sql.Tx) (*entidade.Conta, error) {
			conta, err := uow.Repositorios().Contas.Comando.Abrir(uow.Context(), tx, comando)

			if err != nil {
				return nil, err
			}

			err = messagebus.MessageBus.Publish(
				uow, &mensagens.ContaAberta{
					EId:       uuid.New(),
					Chave:     conta.Chave,
					Timestamp: conta.Timestamp,
				},
			)

			if err != nil {
				return nil, err
			}

			return conta, nil
		},
	)
}

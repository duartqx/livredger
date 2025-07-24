package executores

import (
	"github.com/duartqx/livredger/internal/application"
	"github.com/duartqx/livredger/internal/application/messagebus"

	"github.com/duartqx/livredger/internal/domain/comandos"
	"github.com/duartqx/livredger/internal/domain/eventos"
	"github.com/duartqx/livredger/internal/domain/mensagens"
	"github.com/duartqx/livredger/internal/domain/value/meios"
	"github.com/duartqx/livredger/internal/domain/value/naturezas"
)

func LancamentoContaCriada(uow application.UnidadeDeTrabalho, mensagem messagebus.IdentifiableMessage) error {
	evento, err := messagebus.CastMessage[mensagens.ContaAberta](mensagem)

	if err != nil {
		return err
	}

	_, err = uow.Repositorios().Lancamentos.Comando.Criar(
		uow.Context(),
		uow.Transaction(),
		&comandos.CriarLancamento{
			Evento:         string(eventos.ContaAberta),
			Chave:          evento.Chave,
			Versao:         1,
			Valores:        0,
			Natureza:       naturezas.OUTRO,
			MeioFinanceiro: meios.OUTRO,
			Descricao:      "Abertura da Conta",
			Vencimento:     evento.Timestamp,
		},
	)

	return err
}

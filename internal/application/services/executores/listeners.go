package executores

import (
	"fmt"
	"reflect"

	"github.com/duartqx/livredger/internal/application/messagebus"

	"github.com/duartqx/livredger/internal/domain/comandos"
	"github.com/duartqx/livredger/internal/domain/eventos"
	"github.com/duartqx/livredger/internal/domain/mensagens"
	"github.com/duartqx/livredger/internal/domain/value/meios"
	"github.com/duartqx/livredger/internal/domain/value/naturezas"

	"github.com/duartqx/livredger/internal/infra"
)

func LancamentoContaCriada(uow infra.UnidadeDeTrabalho, mensagem messagebus.IdentifiableMessage) error {
	evento, err := messagebus.CastMessage[mensagens.ContaAberta](mensagem)

	if err != nil {
		return fmt.Errorf("Mensagem inesperada: %s", reflect.TypeOf(mensagem).String())
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

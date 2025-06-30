package executores

import (
	"fmt"
	"reflect"

	"github.com/duartqx/livredger/internal/domain/comandos"
	"github.com/duartqx/livredger/internal/domain/eventos"
	"github.com/duartqx/livredger/internal/domain/mensagens"
	"github.com/duartqx/livredger/internal/infra"
)

func LancamentoContaCriada(uow infra.UnidadeDeTrabalho, mensagem any) error {
	evento, ok := mensagem.(mensagens.ContaAberta)

	if !ok {
		return fmt.Errorf("Mensagem inesperada: %s", reflect.TypeOf(mensagem).String())
	}

	tx := uow.GetTransaction()

	_, err := uow.GetRepositorios().Lancamentos().Comando.Criar(tx, &comandos.CriarLancamento{
		Evento:         string(eventos.ContaAberta),
		Chave:          evento.Chave,
		Versao:         1,
		Valores:        0,
		Natureza:       "Outro",
		MeioFinanceiro: "Outro",
		Descricao:      "Abertura da Conta",
	})

	return err
}

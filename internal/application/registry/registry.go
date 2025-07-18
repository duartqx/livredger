package registry

import (
	"reflect"

	"github.com/duartqx/livredger/internal/domain/mensagens"

	"github.com/duartqx/livredger/internal/application/messagebus"
	"github.com/duartqx/livredger/internal/application/services/executores"
)

func SetupEventHandlers() {
	messagebus.MessageBus.Subscribe(reflect.TypeOf(mensagens.ContaAberta{}), executores.LancamentoContaCriada)
}

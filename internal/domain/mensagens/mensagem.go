package mensagens

import (
	"fmt"
	"reflect"
)

func CastMensagem[M any](msg *any) (*M, error) {
	if msg == nil {
		return nil, fmt.Errorf("Mensagem inesperada: <nil>")
	}

	if mensagem, ok := (*msg).(*M); ok {
		return mensagem, nil
	}

	return nil, fmt.Errorf("Mensagem inesperada: %s", reflect.TypeOf(msg).Elem().Name())
}

package messagebus

import (
	"fmt"
	"reflect"
)

func CastMessage[M IdentifiableMessage](msg IdentifiableMessage) (*M, error) {
	if msg == nil {
		return nil, fmt.Errorf("Mensagem inesperada: <nil>")
	}

	if mensagem, ok := any(msg).(*M); ok {
		return mensagem, nil
	}

	return nil, fmt.Errorf("Unexpected message type: %s", reflect.TypeOf(msg).Elem().Name())
}

func GenerateMessageKey(msg reflect.Type) (string, error) {
	if msg == nil {
		return "", fmt.Errorf("Message can't be <nil>")
	}

	if msg.Kind() == reflect.Ptr {
		msg = msg.Elem()
	}

	if msg.Kind() != reflect.Struct {
		return "", fmt.Errorf("Message is not a struct {%s}", msg.Kind().String())
	}

	return msg.Name(), nil
}

func GetMessageKey(msg any) string {
	typ := reflect.TypeOf(msg)
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}

	return typ.Name()
}

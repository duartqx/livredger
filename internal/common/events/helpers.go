package events

import (
	"fmt"
	"reflect"
)

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

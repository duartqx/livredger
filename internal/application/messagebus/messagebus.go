package messagebus

import (
	"context"
	"fmt"
	"reflect"
	"runtime"
	"sync"

	"github.com/duartqx/livredger/internal/common/events"
	"github.com/duartqx/livredger/internal/common/logger"
	"github.com/duartqx/livredger/internal/domain"
	"github.com/google/uuid"

	"github.com/duartqx/livredger/internal/infra"
)

type IdentifiableMessage interface {
	GetEventIdentifier() uuid.UUID
	GetEntityIdentifier() any
}

var MessageBus = bus{
	registry: make(map[domain.Message][]MessageHandler),
}

type MessageHandler struct {
	Handle func(infra.UnidadeDeTrabalho, any) error
	Name   string
}

type bus struct {
	registry map[domain.Message][]MessageHandler
}

func (mb *bus) Subscribe(typ reflect.Type, handler func(infra.UnidadeDeTrabalho, any) error) {

	key, err := events.GenerateMessageKey(typ)

	if err != nil {
		panic(err)
	}

	handlers, ok := mb.registry[domain.Message(key)]

	if !ok {
		handlers = []MessageHandler{}
	}

	mb.registry[domain.Message(key)] = append(handlers, MessageHandler{
		Handle: handler,
		Name:   runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name(),
	})
}

func (mb *bus) Publish(uow infra.UnidadeDeTrabalho, mensagem IdentifiableMessage) error {

	key := events.GetMessageKey(mensagem)

	if mensagem.GetEventIdentifier() == uuid.Nil {
		return fmt.Errorf("Mensagem com identificador inválido {%s}", key)
	}

	handlers, ok := mb.registry[domain.Message(key)]

	if !ok || len(handlers) == 0 {
		return nil
	}

	errCh := make(chan error, 1)
	defer close(errCh)

	logger.SLogger.Debug(
		key,
		"status", "PUBLISHING",
		"message", mensagem.GetEventIdentifier().String(),
		"entity", mensagem.GetEntityIdentifier(),
	)

	ctx, cancel := context.WithCancel(uow.GetContext())
	defer cancel()

	var wg sync.WaitGroup

	for _, handler := range handlers {

		wg.Add(1)

		go func(handler MessageHandler) {
			logger.SLogger.Debug(
				key,
				"status", "HANDLING",
				"handler", handler.Name,
				"message", mensagem.GetEventIdentifier().String(),
				"entity", mensagem.GetEntityIdentifier(),
			)

			defer wg.Done()

			select {
			case <-ctx.Done():
				return
			default:
			}

			if err := handler.Handle(uow, mensagem); err != nil {

				logger.SLogger.Error(
					key,
					"status", "ERROR",
					"reason", err.Error(),
					"handler", handler.Name,
					"message", mensagem.GetEventIdentifier().String(),
					"entity", mensagem.GetEntityIdentifier(),
				)

				select {
				case errCh <- err:
					cancel()
				default:
				}

				return
			}

			logger.SLogger.Debug(
				key,
				"status", "SUCCESS",
				"handler", handler.Name,
				"message", mensagem.GetEventIdentifier().String(),
				"entity", mensagem.GetEntityIdentifier(),
			)
		}(handler)
	}

	wg.Wait()

	select {
	case err := <-errCh:
		return err
	default:
		return nil
	}
}

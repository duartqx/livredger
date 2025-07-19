package messagebus

import (
	"context"
	"log/slog"
	"os"
	"reflect"
	"runtime"
	"sync"

	"github.com/duartqx/livredger/internal/common/events"
	"github.com/duartqx/livredger/internal/domain"

	"github.com/duartqx/livredger/internal/infra"
)

var logger = slog.New(
	slog.NewJSONHandler(
		os.Stderr,
		&slog.HandlerOptions{
			Level: func() *slog.LevelVar {
				lv := slog.LevelVar{}
				lv.Set(slog.LevelDebug)
				return &lv
			}(),
		},
	),
)

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

func (mb *bus) Publish(uow infra.UnidadeDeTrabalho, mensagem any) error {

	key := events.GetMessageKey(mensagem)

	handlers, ok := mb.registry[domain.Message(key)]

	if !ok || len(handlers) == 0 {
		return nil
	}

	errCh := make(chan error, 1)
	defer close(errCh)

	logger.Debug(key, "Status", "Publishing")

	ctx, cancel := context.WithCancel(uow.GetContext())
	defer cancel()

	var wg sync.WaitGroup

	for _, handler := range handlers {

		wg.Add(1)

		go func(handler MessageHandler) {
			logger.Debug(key, "Status", "Handling", "Handler", handler.Name)

			defer wg.Done()

			select {
			case <-ctx.Done():
				return
			default:
			}

			if err := handler.Handle(uow, mensagem); err != nil {

				logger.Error(key, "Status", "Error", "Reason", err.Error(), "Handler", handler.Name)

				select {
				case errCh <- err:
					cancel()
				default:
				}

				return
			}

			logger.Debug(key, "Status", "Success", "Handler", handler.Name)
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

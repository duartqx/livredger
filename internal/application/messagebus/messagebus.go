package messagebus

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"reflect"
	"runtime"
	"sync"

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
	registry: make(map[domain.Mensagem][]MessageHandler),
}

type MessageHandler struct {
	Handle func(infra.UnidadeDeTrabalho, any) error
	Name   string
}

type bus struct {
	registry map[domain.Mensagem][]MessageHandler
}

func (mb *bus) Subscribe(mensagem any, handler func(infra.UnidadeDeTrabalho, any) error) {

	typ := reflect.TypeOf(mensagem)

	if typ.Kind() != reflect.Struct {
		panic(fmt.Sprintf("Mensagem não permitida, não é Struct"))
	}

	key := domain.Mensagem(typ.Name())

	handlers := mb.registry[key]

	if handlers == nil {
		handlers = []MessageHandler{}
	}

	mb.registry[key] = append(handlers, MessageHandler{
		Handle: handler,
		Name:   runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name(),
	})
}

func (mb *bus) Publish(uow infra.UnidadeDeTrabalho, mensagem any) <-chan error {
	errCh := make(chan error, 1)

	go func() {
		defer close(errCh)

		key := domain.Mensagem(reflect.TypeOf(mensagem).Name())

		handlers, ok := mb.registry[key]

		if !ok || len(handlers) == 0 {
			errCh <- nil
			return
		} else {
			logger.Debug(string(key), "Status", "Publishing")
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var wg sync.WaitGroup

		for _, handler := range handlers {

			wg.Add(1)

			go func() {
				logger.Debug(string(key), "Status", "Handling", "Handler", handler.Name)

				defer wg.Done()

				select {
				case <-ctx.Done():
					return
				default:
				}

				if err := handler.Handle(uow, mensagem); err != nil {

					logger.Error(string(key), "Status", "Error", "Reason", err.Error(), "Handler", handler.Name)

					select {
					case errCh <- err:
						cancel()
					default:
					}

					return
				}

				logger.Debug(string(key), "Status", "Success", "Handler", handler.Name)
			}()
		}

		wg.Wait()

		select {
		case errCh <- nil:
		default:
		}
	}()

	return errCh
}

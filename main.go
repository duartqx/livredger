package main

import (
	"context"
	"fmt"
	"os"

	"flag"
	"io/fs"
	"log"
	"net/http"
	"os/signal"
	"time"

	"github.com/duartqx/livredger/internal/api/routers"
	"github.com/duartqx/livredger/internal/application"
	"github.com/duartqx/livredger/internal/application/messagebus/registry"
	"github.com/duartqx/livredger/internal/infra"
)

const (
	staticPath    = "internal/presentation/static"
	templatesPath = "internal/presentation/templates"
)

type config struct {
	Port           int
	ServerTimeout  int
	RequestTimeout int
}

var cfg config

var (
	static    fs.FS
	templates fs.FS
)

func main() {

	flag.IntVar(&cfg.Port, "port", 8000, "The port the server will run at")
	flag.IntVar(&cfg.ServerTimeout, "servertm", 5, "The time in seconds for server timeout")
	flag.IntVar(&cfg.RequestTimeout, "requesttm", 1, "The time in seconds for request timeout")

	flag.Parse()

	application.SetupApplication(infra.FabricaDeUnidadeDeTrabalho)
	registry.SetupEventHandlers()

	srv := &http.Server{
		Handler: routers.Router(
			&routers.Dependencies{
				Templates: templates,
				Static: &[]routers.Static{
					{Path: "/static/", Fs: static},
				},
				RequestTimeout: time.Duration(cfg.RequestTimeout),
			},
		),
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		WriteTimeout: time.Duration(cfg.ServerTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.ServerTimeout) * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	log.Println("Listening and Serving at:", cfg.Port)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(cfg.ServerTimeout))
	defer cancel()

	srv.Shutdown(ctx)

	os.Exit(0)
}

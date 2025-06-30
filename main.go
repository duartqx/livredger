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
	"github.com/duartqx/livredger/internal/application/registry"
)

const (
	staticPath    = "internal/presentation/static"
	templatesPath = "internal/presentation/templates"
)

var (
	static    fs.FS
	templates fs.FS
	port      int
)

func main() {
	flag.IntVar(&port, "port", 8000, "The port the server will run at")

	flag.Parse()

	registry.SetupEventHandlers()

	srv := &http.Server{
		Handler: routers.Router(
			&routers.Dependencies{
				Templates: templates,
				Static: &[]routers.Static{
					{Path: "/static/", Fs: static},
				},
			},
		),
		Addr:         fmt.Sprintf(":%d", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	log.Println("Listening and Serving at:", port)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	srv.Shutdown(ctx)

	os.Exit(0)
}

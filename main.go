package main

import (
	"context"
	"fmt"
	"os"

	"embed"
	"flag"
	"io/fs"
	"log"
	"net/http"
	"os/signal"
	"time"

	"github.com/duartqx/livredger/internal/api/routers"
)

var (
	err       error
	static    fs.FS
	templates fs.FS
	port      int
	srv       *http.Server
)

//go:embed internal/presentation/static/*
var staticFS embed.FS

//go:embed internal/presentation/templates/*
var templatesFS embed.FS

func init() {
	flag.IntVar(&port, "port", 8000, "The port the server will run at")

	flag.Parse()

	static, err = fs.Sub(staticFS, "internal/presentation/static")
	if err != nil {
		log.Fatalln(err)
	}

	templates, err = fs.Sub(templatesFS, "internal/presentation/templates")
	if err != nil {
		log.Fatalln(err)
	}

	srv = &http.Server{
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
}

func main() {
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

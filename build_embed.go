//go:build embed
// +build embed

package main

import (
	"embed"
)

//go:embed internal/presentation/static/*
var staticFS embed.FS

//go:embed internal/presentation/templates/*
var templatesFS embed.FS

func init() {
	var err error

	static, err = fs.Sub(staticFS, staticPath)
	if err != nil {
		log.Fatalln(err)
	}

	templates, err = fs.Sub(templatesFS, templatesPath)
	if err != nil {
		log.Fatalln(err)
	}
}

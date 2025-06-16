//go:build !embed
// +build !embed

package main

import (
	"os"
)

func init() {
	static = os.DirFS(staticPath)

	templates = os.DirFS(templatesPath)
}

//go:build !dev
// +build !dev

package main

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed public
var publicFS embed.FS

func public() http.Handler {
	fs, err := fs.Sub(publicFS, "public")
	if err != nil {
		panic(err)
	}
	return http.FileServer(http.FS(fs))
}

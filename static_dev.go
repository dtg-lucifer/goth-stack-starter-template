//go:build dev
// +build dev

package main

import "net/http"

func public() http.Handler {
	return http.FileServer(http.Dir("public"))
}

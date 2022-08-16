// Package main
// Time    : 2022/8/15 23:00
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("welcome"))
	})
	log.Fatalln(http.ListenAndServe(":3000", r))
}

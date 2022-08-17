// Package main
// Time    : 2022/8/17 21:55
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here.
		w.WriteHeader(503)
		_, _ = w.Write([]byte("bad"))
	}
}

func main() {
	r := chi.NewRouter()
	r.Method("GET", "/", Handler(customHandler))
	log.Fatalln(http.ListenAndServe(":3333", r))
}

func customHandler(w http.ResponseWriter, r *http.Request) error {
	q := r.URL.Query().Get("err")

	if q != "" {
		return errors.New(q)
	}

	_, _ = w.Write([]byte("foo"))
	return nil
}

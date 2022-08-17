// Package main
// Time    : 2022/8/17 22:00
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	// Index handler
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	// Create a route along /files that will serve contents from
	// the ./data/ folder.
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "data"))
	FileServer(r, "/files", filesDir)

	err := http.ListenAndServe(":3333", r)
	if err != nil {
		return
	}
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

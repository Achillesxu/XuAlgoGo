// Package main
// Time    : 2022/8/17 21:26
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/_examples/versions/data"
	v1 "github.com/go-chi/chi/v5/_examples/versions/presenter/v1"
	v2 "github.com/go-chi/chi/v5/_examples/versions/presenter/v2"
	v3 "github.com/go-chi/chi/v5/_examples/versions/presenter/v3"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	// version 3
	r.Route("v3", func(r chi.Router) {
		r.Use(apiVersionCtx("v3"))
		r.Mount("articles", articleRouter())
	})

	// API version 2.
	r.Route("/v2", func(r chi.Router) {
		r.Use(apiVersionCtx("v2"))
		r.Mount("/articles", articleRouter())
	})

	// API version 1.
	r.Route("/v1", func(r chi.Router) {
		r.Use(randomErrorMiddleware) // Simulate random error, ie. version 1 is buggy.
		r.Use(apiVersionCtx("v1"))
		r.Mount("/articles", articleRouter())
	})

	log.Fatalln(http.ListenAndServe(":3333", r))
}

func apiVersionCtx(version string) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), "api.version", version))
			next.ServeHTTP(w, r)
		})
	}
}

func articleRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", listArticles)
	r.Route("/{articleId}", func(r chi.Router) {
		r.Get("/", getArticle)
		// r.Delete("/", deleteArticle)
	})
	return r
}

func listArticles(w http.ResponseWriter, r *http.Request) {
	articles := make(chan render.Renderer, 5)

	go func() {
		// Load data asynchronously into the channel (simulate slow storage):
		for i := 0; i < 10; i++ {
			article := &data.Article{
				ID:                     i,
				Title:                  fmt.Sprintf("Article #%v", i),
				Data:                   []string{"one", "two", "three", "four"},
				CustomDataForAuthUsers: "secret data for auth'd users only",
			}
			apiVersion := r.Context().Value("api.version").(string)
			switch apiVersion {
			case "v1":
				articles <- v1.NewArticleResponse(article)
			case "v2":
				articles <- v2.NewArticleResponse(article)
			default:
				articles <- v3.NewArticleResponse(article)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	render.Respond(w, r, articles)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	// Load article.
	if chi.URLParam(r, "articleID") != "1" {
		render.Respond(w, r, data.ErrNotFound)
		return
	}
	article := &data.Article{
		ID:                     1,
		Title:                  "Article #1",
		Data:                   []string{"one", "two", "three", "four"},
		CustomDataForAuthUsers: "secret data for auth'd users only",
	}

	// Simulate some context values:
	// 1. ?auth=true simluates authenticated session/user.
	// 2. ?error=true simulates random error.
	if r.URL.Query().Get("auth") != "" {
		r = r.WithContext(context.WithValue(r.Context(), "auth", true))
	}
	if r.URL.Query().Get("error") != "" {
		render.Respond(w, r, errors.New("error"))
		return
	}

	var payload render.Renderer

	apiVersion := r.Context().Value("api.version").(string)
	switch apiVersion {
	case "v1":
		payload = v1.NewArticleResponse(article)
	case "v2":
		payload = v2.NewArticleResponse(article)
	default:
		payload = v3.NewArticleResponse(article)
	}

	_ = render.Render(w, r, payload)
}

func randomErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rand.Seed(time.Now().Unix())

		// One in three chance of random error.
		if rand.Int31n(3) == 0 {
			errs := []error{data.ErrUnauthorized, data.ErrForbidden, data.ErrNotFound}
			render.Respond(w, r, errs[rand.Intn(len(errs))])
			return
		}
		next.ServeHTTP(w, r)
	})
}

package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kireeti-28/soul-tidings/pkg/auth"
	"github.com/kireeti-28/soul-tidings/pkg/handlers"
)

func routes() http.Handler {
	router := chi.NewRouter()

	router.Use(auth.MiddlewareLog)
	router.Use(auth.MiddlewareCors)

	fsHandler := http.FileServer(http.Dir("./frontend"))
	router.Handle("/", fsHandler)
	router.Handle("/*", fsHandler)

	router.Post("/user/login", handlers.Login)
	router.Post("/user/register", handlers.Register)

	return router
}

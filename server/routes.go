package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/kireeti-28/soul-tidings/pkg/auth"
	"github.com/kireeti-28/soul-tidings/pkg/handlers"
)

func routes() http.Handler {
	router := chi.NewRouter()

	router.Use(auth.MiddlewareLog)
	router.Use(auth.MiddlewareCors)

	FileServer(router)

	apiRouter := chi.NewRouter()
	apiRouter.Post("/user/login", handlers.Login)
	apiRouter.Post("/user/register", handlers.Register)

	router.Mount("/api", apiRouter)

	return router
}

func FileServer(router *chi.Mux) {
	root := "./frontend"
	fs := http.FileServer(http.Dir(root))

	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(root + r.RequestURI); os.IsNotExist(err) {
			http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
		} else {
			fs.ServeHTTP(w, r)
		}
	})

}

package main

import (
	"net/http"

	"internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	chiRouter := chi.NewRouter()
	chiRouter.Use(middleware.Logger)

	chiRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	chiRouter.Mount("/login", handlers.LoginRouter())

	http.ListenAndServe(":3000", chiRouter)
}
